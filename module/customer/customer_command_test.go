package customer

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

type customerCommandTestSuite struct {
	suite.Suite
	db                *database.DB
	mock              sqlmock.Sqlmock
	ctx               context.Context
	customerStructure []string
}

func TestCustomerCommandTestSuite(t *testing.T) {
	suite.Run(t, new(customerCommandTestSuite))
}

func (c *customerCommandTestSuite) SetupTest() {
	c.ctx = context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		c.Fail("error while opening stub db")
	}
	c.db = &database.DB{
		DB:      db,
		Dialect: "mysql",
	}
	c.mock = mock
	c.customerStructure = []string{
		"customer_id",
		"email",
		"created_at",
	}
}

func (c *customerCommandTestSuite) TestAddCustomer_ExecError() {
	var expectedEntity = CustomerEntity{
		Email:     "test@mail.com",
		CreatedAt: time.Now().UTC(),
	}
	c.mock.ExpectExec(regexp.QuoteMeta(addNewCustomer)).
		WithArgs(expectedEntity.Email, expectedEntity.CreatedAt).
		WillReturnError(errors.New("new error"))
	var command = NewCustomerCommand(c.db)
	var res = command.AddCustomer(c.ctx, expectedEntity)
	c.NotNil(res)
	c.EqualError(res, "new error")
}

func (c *customerCommandTestSuite) TestAddCustomer_AffectedZero() {
	var expectedEntity = CustomerEntity{
		Email:     "test@mail.com",
		CreatedAt: time.Now().UTC(),
	}
	c.mock.ExpectExec(regexp.QuoteMeta(addNewCustomer)).
		WithArgs(expectedEntity.Email, expectedEntity.CreatedAt).
		WillReturnResult(sqlmock.NewResult(0, 0))
	var command = NewCustomerCommand(c.db)
	var res = command.AddCustomer(c.ctx, expectedEntity)
	c.NotNil(res)
	c.EqualError(res, "failed to add new customer")
}

func (c *customerCommandTestSuite) TestAddCustomer_Ok() {
	var expectedEntity = CustomerEntity{
		Email:     "test@mail.com",
		CreatedAt: time.Now().UTC(),
	}
	c.mock.ExpectExec(regexp.QuoteMeta(addNewCustomer)).
		WithArgs(expectedEntity.Email, expectedEntity.CreatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))
	var command = NewCustomerCommand(c.db)
	var res = command.AddCustomer(c.ctx, expectedEntity)
	c.Nil(res)
}
