package order

import (
	"catalyst-case/database"
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type orderCommandTestSuite struct {
	suite.Suite
	mock            sqlmock.Sqlmock
	db              *database.DB
	ctx             context.Context
	masterStructure []string
	detailStructure []string
}

func TestOrderCommandTestSuite(t *testing.T) {
	suite.Run(t, new(orderCommandTestSuite))
}

func (o *orderCommandTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		o.Fail("cannot open stub db")
	}
	o.db = &database.DB{
		DB: db,
	}
	o.mock = mock
	o.ctx = context.Background()
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
}

func (o *orderCommandTestSuite) TestAddNewOrder_ReturnErrorBeginTx() {
	o.mock.ExpectBegin().WillReturnError(errors.New("begin tx error"))
	var svc = NewOrderCommand(o.db)
	wrap, c := svc.AddNewOrder(o.ctx, OrderMasterEntity{}, []OrderDetailEntity{})
	o.Nil(wrap)
	o.EqualError(c, "begin tx error")
}

func (o *orderCommandTestSuite) TestAddNewOrder_ExecMasterError() {
	grandTotal := sql.NullFloat64{
		Float64: float64(150000),
		Valid:   true,
	}
	createdAt := time.Now().UTC()
	createdBy := 1
	o.mock.ExpectBegin()
	o.mock.ExpectExec(regexp.QuoteMeta(addNewMaster)).
		WithArgs(grandTotal, createdAt, createdBy).
		WillReturnError(errors.New("tx error"))
	var svc = NewOrderCommand(o.db)
	wrap, c := svc.AddNewOrder(o.ctx, OrderMasterEntity{
		GrandTotal: grandTotal,
		CreatedAt:  createdAt,
		CreatedBy:  createdBy,
	}, []OrderDetailEntity{})
	o.Nil(wrap)
	o.EqualError(c, "tx error")
}

func (o *orderCommandTestSuite) TestAddNewOrder_ExecDetailError() {
	grandTotal := sql.NullFloat64{
		Float64: float64(150000),
		Valid:   true,
	}
	createdAt := time.Now().UTC()
	createdBy := 1
	item := OrderDetailEntity{
		OrderId:     1,
		ProductId:   12,
		ProductName: "jam tangan",
		UnitPrice:   float64(150000),
		Qty:         1,
	}
	o.mock.ExpectBegin()
	o.mock.ExpectExec(regexp.QuoteMeta(addNewMaster)).
		WithArgs(grandTotal, createdAt, createdBy).
		WillReturnResult(sqlmock.NewResult(1, 1))
	o.mock.ExpectExec(regexp.QuoteMeta(addNewDetail)).
		WithArgs(item.OrderId, item.ProductId, item.ProductName, item.UnitPrice, item.Qty).
		WillReturnError(errors.New("tx error"))
	var svc = NewOrderCommand(o.db)
	wrap, c := svc.AddNewOrder(o.ctx, OrderMasterEntity{
		GrandTotal: grandTotal,
		CreatedAt:  createdAt,
		CreatedBy:  createdBy,
	}, []OrderDetailEntity{
		item,
	})
	o.Nil(wrap)
	o.EqualError(c, "tx error")
}

func (o *orderCommandTestSuite) TestAddNewOrder_CommitSuccess() {
	grandTotal := sql.NullFloat64{
		Float64: float64(150000),
		Valid:   true,
	}
	createdAt := time.Now().UTC()
	createdBy := 1
	item := OrderDetailEntity{
		OrderId:     1,
		ProductId:   12,
		ProductName: "jam tangan",
		UnitPrice:   float64(150000),
		Qty:         1,
	}
	o.mock.ExpectBegin()
	o.mock.ExpectExec(regexp.QuoteMeta(addNewMaster)).
		WithArgs(grandTotal, createdAt, createdBy).
		WillReturnResult(sqlmock.NewResult(1, 1))
	o.mock.ExpectExec(regexp.QuoteMeta(addNewDetail)).
		WithArgs(item.OrderId, item.ProductId, item.ProductName, item.UnitPrice, item.Qty).
		WillReturnResult(sqlmock.NewResult(1, 1))
	o.mock.ExpectCommit()
	var svc = NewOrderCommand(o.db)
	wrap, c := svc.AddNewOrder(o.ctx, OrderMasterEntity{
		GrandTotal: grandTotal,
		CreatedAt:  createdAt,
		CreatedBy:  createdBy,
	}, []OrderDetailEntity{
		item,
	})
	if err := o.mock.ExpectationsWereMet(); err != nil {
		o.Fail("there were unfulfilled expectations: %s", err)
	}
	o.Nil(c)
	o.Equal(1, wrap.Order.OrderId)
	o.Equal(1, wrap.OrderDetails[0].OrderDetailId)
}
