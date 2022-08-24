package order

import (
	"catalyst-case/module/product"
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type orderTestSuite struct {
	suite.Suite
	productQuery *product.ProductQueryMock
	orderCommand *orderCommandMock
	orderQuery   *orderQueryMock
	ctx          context.Context
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(orderTestSuite))
}

func (o *orderTestSuite) SetupTest() {
	o.productQuery = new(product.ProductQueryMock)
	o.orderCommand = new(orderCommandMock)
	o.orderQuery = new(orderQueryMock)
	o.ctx = context.Background()
}

func (o *orderTestSuite) TestAddNewOrder_ReturnError() {
	o.productQuery.On("GetProductById", o.ctx, 1).Return(nil, errors.New("query error"))
	svc := NewOrder(o.orderCommand, o.productQuery, o.orderQuery)
	res, e := svc.AddNewOrder(o.ctx, OrderEntry{
		Items: []OrderItem{
			{
				ProductId: 1,
				Qty:       2,
			},
		},
	})
	o.Nil(res)
	o.EqualError(e, "query error")
}

func (o *orderTestSuite) TestAddNewOrder_ReturnOk() {
	o.productQuery.On("GetProductById", o.ctx, 1).Return(&product.ProductDto{
		ProductId:   1,
		BrandId:     2,
		BrandName:   "new brand",
		ProductName: "jam tangan",
		Price:       10000,
		CreatedAt:   time.Now().UTC(),
	}, nil)
	o.orderCommand.On("AddNewOrder", o.ctx, mock.Anything, mock.Anything).Return(&OrderWrap{
		Order: OrderMasterEntity{
			OrderId: 1,
			GrandTotal: sql.NullFloat64{
				Float64: 20000,
				Valid:   true,
			},
			CreatedAt: time.Now().UTC(),
			CreatedBy: 1,
		},
		OrderDetails: []OrderDetailEntity{
			{
				OrderDetailId: 1,
				OrderId:       1,
				ProductId:     12,
				ProductName:   "jam tangan",
				UnitPrice:     10000,
				Qty:           2,
			},
		},
	}, nil)
	svc := NewOrder(o.orderCommand, o.productQuery, o.orderQuery)
	res, e := svc.AddNewOrder(o.ctx, OrderEntry{
		Items: []OrderItem{
			{
				ProductId: 1,
				Qty:       2,
			},
		},
		CreatedBy: 1,
	})
	o.Equal("jam tangan", res.Items[0].ProductName)
	o.Equal(1, res.CreatedBy)
	o.Nil(e)
}

func (o *orderTestSuite) TestGetOrders_ReturnError() {
	o.orderQuery.On("GetOrders", o.ctx).Return(nil, errors.New("new error"))
	svc := NewOrder(o.orderCommand, o.productQuery, o.orderQuery)
	res, e := svc.GetOrders(o.ctx)
	o.Nil(res)
	o.EqualError(e, "new error")
}

func (o *orderTestSuite) TestGetOrders_ReturnOk() {
	o.orderQuery.On("GetOrders", o.ctx).Return([]OrderWrap{
		{
			Order: OrderMasterEntity{
				OrderId: 1,
				GrandTotal: sql.NullFloat64{
					Float64: 20000,
					Valid:   true,
				},
				CreatedAt: time.Now().UTC(),
				CreatedBy: 12,
			},
			OrderDetails: []OrderDetailEntity{
				{
					OrderDetailId: 21,
					OrderId:       1,
					ProductId:     13,
					ProductName:   "jam tangan",
					UnitPrice:     20000,
					Qty:           1,
				},
			},
		},
	}, nil)
	svc := NewOrder(o.orderCommand, o.productQuery, o.orderQuery)
	res, e := svc.GetOrders(o.ctx)
	o.Nil(e)
	o.Equal(1, res[0].OrderId)
}

func (o *orderTestSuite) TestGetOrder_ReturnError() {
	o.orderQuery.On("GetOrder", o.ctx, 1).Return(nil, errors.New("query error"))
	svc := NewOrder(o.orderCommand, o.productQuery, o.orderQuery)
	res, e := svc.GetOrder(o.ctx, 1)
	o.Nil(res)
	o.EqualError(e, "query error")
}

func (o *orderTestSuite) TestGetOrder_ReturnOk() {
	o.orderQuery.On("GetOrder", o.ctx, 1).Return(&OrderWrap{
		Order: OrderMasterEntity{
			OrderId: 1,
			GrandTotal: sql.NullFloat64{
				Float64: 20000,
				Valid:   true,
			},
			CreatedAt: time.Now().UTC(),
			CreatedBy: 1,
		},
		OrderDetails: []OrderDetailEntity{
			{
				OrderDetailId: 12,
				OrderId:       1,
				ProductId:     22,
				ProductName:   "jam tangan",
				UnitPrice:     20000,
				Qty:           1,
			},
		},
	}, nil)
	svc := NewOrder(o.orderCommand, o.productQuery, o.orderQuery)
	res, e := svc.GetOrder(o.ctx, 1)
	o.Nil(e)
	o.Equal("jam tangan", res.Items[0].ProductName)
}
