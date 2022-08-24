package order

import (
	"catalyst-case/module/product"
	"context"
	"database/sql"
	"time"
)

type order struct {
	orderCommand OrderCommand
	productQuery product.ProductQuery
	orderQuery   OrderQuery
}

type Order interface {
	AddNewOrder(ctx context.Context, order OrderEntry) (*OrderDto, error)
	GetOrders(ctx context.Context) ([]OrderDto, error)
	GetOrder(ctx context.Context, orderId int) (*OrderDto, error)
}

func NewOrder(orderCommand OrderCommand, productQuery product.ProductQuery, orderQuery OrderQuery) Order {
	return &order{
		orderCommand: orderCommand,
		productQuery: productQuery,
		orderQuery:   orderQuery,
	}
}

func (o *order) AddNewOrder(ctx context.Context, order OrderEntry) (*OrderDto, error) {
	// prep order
	var items = []OrderDetailEntity{}
	var grandTotal float64
	for _, det := range order.Items {
		product, e := o.productQuery.GetProductById(ctx, det.ProductId)
		if e != nil {
			return nil, e
		}
		grandTotal += product.Price * float64(det.Qty)
		items = append(items, OrderDetailEntity{
			ProductId:   product.ProductId,
			ProductName: product.ProductName,
			UnitPrice:   product.Price,
			Qty:         det.Qty,
		})
	}
	var master = OrderMasterEntity{
		GrandTotal: sql.NullFloat64{
			Float64: grandTotal,
			Valid:   true,
		},
		CreatedAt: time.Now().UTC(),
		CreatedBy: order.CreatedBy,
	}
	result, e := o.orderCommand.AddNewOrder(ctx, master, items)
	if e != nil {
		return nil, e
	}
	dto := wrapToDto(result)
	return dto, nil
}

func (o *order) GetOrders(ctx context.Context) ([]OrderDto, error) {
	res, e := o.orderQuery.GetOrders(ctx)
	if e != nil {
		return nil, e
	}
	var dtos = []OrderDto{}
	for _, item := range res {
		dto := wrapToDto(&item)
		dtos = append(dtos, *dto)
	}
	return dtos, nil
}

func wrapToDto(orderWrap *OrderWrap) *OrderDto {
	if orderWrap == nil {
		return nil
	}
	detailDtos := []OrderDetailDto{}
	for _, item := range orderWrap.OrderDetails {
		det := OrderDetailDto{
			OrderDetailId: item.OrderDetailId,
			OrderId:       item.OrderId,
			ProductName:   item.ProductName,
			Qty:           item.Qty,
			UnitPrice:     item.UnitPrice,
		}
		detailDtos = append(detailDtos, det)
	}
	dto := &OrderDto{
		OrderId:    orderWrap.Order.OrderId,
		GrandTotal: orderWrap.Order.GrandTotal.Float64,
		CreatedAt:  orderWrap.Order.CreatedAt,
		Items:      detailDtos,
		CreatedBy:  orderWrap.Order.CreatedBy,
	}
	return dto
}

func (o *order) GetOrder(ctx context.Context, orderId int) (*OrderDto, error) {
	res, e := o.orderQuery.GetOrder(ctx, orderId)
	if e != nil {
		return nil, e
	}
	dto := wrapToDto(res)
	return dto, nil
}
