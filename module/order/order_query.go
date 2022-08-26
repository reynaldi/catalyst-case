package order

import (
	"catalyst-case/database"
	"context"
	"database/sql"
)

type orderQuery struct {
	*database.DB
}

type OrderQuery interface {
	GetOrders(ctx context.Context) ([]OrderWrap, error)
	GetOrder(ctx context.Context, orderId int) (*OrderWrap, error)
}

func NewOrderQuery(db *database.DB) OrderQuery {
	return &orderQuery{
		DB: db,
	}
}

const (
	getOrders      = `SELECT order_id, grand_total, created_at, created_by FROM orders ORDER BY created_at`
	getOrderDetail = `SELECT order_detail_id, order_id, product_id, product_name, unit_price, qty FROM order_details WHERE order_id = ?`
	getOrder       = `SELECT order_id, grand_total, created_at, created_by FROM orders WHERE order_id = ?`
)

func (o *orderQuery) GetOrders(ctx context.Context) ([]OrderWrap, error) {
	res, e := o.QueryContext(ctx, getOrders)
	if e != nil {
		return nil, e
	}
	defer res.Close()

	var orders = []OrderMasterEntity{}
	for res.Next() {
		var master OrderMasterEntity
		e = res.Scan(&master.OrderId, &master.GrandTotal, &master.CreatedAt, &master.CreatedBy)
		if e != nil {
			return nil, e
		}
		orders = append(orders, master)
	}

	var orderAll = []OrderWrap{}
	for _, order := range orders {
		// get detail
		wrap, e := o.wrapOrder(ctx, order)
		if e != nil {
			return nil, e
		}
		orderAll = append(orderAll, *wrap)
	}

	return orderAll, nil
}

func (o *orderQuery) GetOrder(ctx context.Context, orderId int) (*OrderWrap, error) {
	res := o.QueryRowContext(ctx, getOrder, orderId)
	var master OrderMasterEntity
	e := res.Scan(&master.OrderId, &master.GrandTotal, &master.CreatedAt, &master.CreatedBy)
	if e != nil {
		if e == sql.ErrNoRows {
			return nil, nil
		}
		return nil, e
	}
	wrap, e := o.wrapOrder(ctx, master)
	if e != nil {
		return nil, nil
	}
	return wrap, nil
}

func (o *orderQuery) wrapOrder(ctx context.Context, order OrderMasterEntity) (*OrderWrap, error) {
	// get detail
	var details = []OrderDetailEntity{}
	det, e := o.QueryContext(ctx, getOrderDetail, order.OrderId)
	if e != nil {
		return nil, e
	}
	defer det.Close()
	for det.Next() {
		var detail OrderDetailEntity
		e = det.Scan(&detail.OrderDetailId, &detail.OrderId, &detail.ProductId, &detail.ProductName, &detail.UnitPrice, &detail.Qty)
		if e != nil {
			return nil, e
		}
		details = append(details, detail)
	}
	return &OrderWrap{
		Order:        order,
		OrderDetails: details,
	}, nil
}
