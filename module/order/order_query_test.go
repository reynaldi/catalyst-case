package order

import (
	"catalyst-case/database"
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type orderQueryTestSuite struct {
	suite.Suite
	mock            sqlmock.Sqlmock
	db              *database.DB
	masterStructure []string
	detailStructure []string
	ctx             context.Context
}

func TestOrderQueryTestSuite(t *testing.T) {
	suite.Run(t, new(orderQueryTestSuite))
}

func (o *orderQueryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		o.Fail("fail to open stub db")
	}
	o.db = &database.DB{
		DB: db,
	}
	o.mock = mock
	o.masterStructure = []string{
		"order_id",
		"grand_total",
		"created_at",
		"created_by",
	}
	o.detailStructure = []string{
		"order_detail_id",
		"order_id",
		"product_id",
		"product_name",
		"unit_price",
		"qty",
	}
	o.ctx = context.Background()
}

func (o *orderQueryTestSuite) TestGetOrders_QueryMasterError() {
	o.mock.ExpectQuery(regexp.QuoteMeta(getOrders)).WillReturnError(errors.New("error db"))
	svc := NewOrderQuery(o.db)
	res, e := svc.GetOrders(o.ctx)
	o.Nil(res)
	o.EqualError(e, "error db")
}

func (o *orderQueryTestSuite) TestGetOrders_Success() {
	o.mock.ExpectQuery(regexp.QuoteMeta(getOrders)).
		WillReturnRows(sqlmock.NewRows(o.masterStructure).
			AddRow(1, float64(15000), time.Now().UTC(), 1))
	o.mock.ExpectQuery(regexp.QuoteMeta(getOrderDetail)).WithArgs(1).
		WillReturnRows(sqlmock.NewRows(o.detailStructure).
			AddRow(1, 1, 20, "jam tangan", float64(15000), 1))
	svc := NewOrderQuery(o.db)
	res, e := svc.GetOrders(o.ctx)
	o.Nil(e)
	o.Equal(1, res[0].Order.OrderId)
	o.Equal("jam tangan", res[0].OrderDetails[0].ProductName)
}

func (o *orderQueryTestSuite) TestGetOrder_QueryMasterError() {
	o.mock.ExpectQuery(regexp.QuoteMeta(getOrder)).WithArgs(1).
		WillReturnError(errors.New("db error"))
	svc := NewOrderQuery(o.db)
	res, e := svc.GetOrder(o.ctx, 1)
	o.Nil(res)
	o.EqualError(e, "db error")
}

func (o *orderQueryTestSuite) TestGetOrder_Success() {
	o.mock.ExpectQuery(regexp.QuoteMeta(getOrder)).
		WillReturnRows(sqlmock.NewRows(o.masterStructure).
			AddRow(1, float64(15000), time.Now().UTC(), 1))
	o.mock.ExpectQuery(regexp.QuoteMeta(getOrderDetail)).WithArgs(1).
		WillReturnRows(sqlmock.NewRows(o.detailStructure).
			AddRow(1, 1, 20, "jam tangan", float64(15000), 1))
	svc := NewOrderQuery(o.db)
	res, e := svc.GetOrder(o.ctx, 1)
	o.Nil(e)
	o.Equal(1, res.Order.OrderId)
	o.Equal("jam tangan", res.OrderDetails[0].ProductName)
}
