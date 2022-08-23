package customer

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

type customerQueryTestSuite struct {
	suite.Suite
	db                *database.DB
	mock              sqlmock.Sqlmock
	ctx               context.Context
	customerStructure []string
}

func TestCustomerQueryTestSuite(t *testing.T) {
	suite.Run(t, new(customerQueryTestSuite))
}

func (c *customerQueryTestSuite) SetupTest() {
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

func (c *customerQueryTestSuite) TestGetCustomerById_ReturnNoRow() {
	c.mock.ExpectQuery(regexp.QuoteMeta(getCustomerById)).WithArgs(1).WillReturnError(sql.ErrNoRows)
	var query = NewCustomerQuery(c.db)
	res, e := query.GetCustomerById(c.ctx, 1)
	c.Nil(res)
	c.Nil(e)
}

func (c *customerQueryTestSuite) TestGetCustomerById_ReturnError() {
	c.mock.ExpectQuery(regexp.QuoteMeta(getCustomerById)).WithArgs(1).WillReturnError(errors.New("new error"))
	var query = NewCustomerQuery(c.db)
	res, e := query.GetCustomerById(c.ctx, 1)
	c.Nil(res)
	c.EqualError(e, "new error")
}

func (c *customerQueryTestSuite) TestGetCustomerById_ReturnOk() {
	c.mock.ExpectQuery(regexp.QuoteMeta(getCustomerById)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(c.customerStructure).AddRow(
			1,
			"test@mail.com",
			time.Now().UTC(),
		))
	var query = NewCustomerQuery(c.db)
	res, e := query.GetCustomerById(c.ctx, 1)
	c.Nil(e)
	c.Equal(res.CustomerId, 1)
	c.Equal(res.Email, "test@mail.com")
}

func (c *customerQueryTestSuite) TestGetCustomerByEmail_ReturnNoRow() {
	c.mock.ExpectQuery(regexp.QuoteMeta(getCustomerByEmail)).WithArgs("test@mail.com").WillReturnError(sql.ErrNoRows)
	var query = NewCustomerQuery(c.db)
	res, e := query.GetCustomerByEmail(c.ctx, "test@mail.com")
	c.Nil(res)
	c.Nil(e)
}

func (c *customerQueryTestSuite) TestGetCustomerByEmail_ReturnError() {
	c.mock.ExpectQuery(regexp.QuoteMeta(getCustomerByEmail)).WithArgs("test@mail.com").WillReturnError(errors.New("new error"))
	var query = NewCustomerQuery(c.db)
	res, e := query.GetCustomerByEmail(c.ctx, "test@mail.com")
	c.Nil(res)
	c.EqualError(e, "new error")
}

func (c *customerQueryTestSuite) TestGetCustomerByEmail_ReturnOk() {
	c.mock.ExpectQuery(regexp.QuoteMeta(getCustomerByEmail)).
		WithArgs("test@mail.com").
		WillReturnRows(sqlmock.NewRows(c.customerStructure).AddRow(
			1,
			"test@mail.com",
			time.Now().UTC(),
		))
	var query = NewCustomerQuery(c.db)
	res, e := query.GetCustomerByEmail(c.ctx, "test@mail.com")
	c.Nil(e)
	c.Equal(res.CustomerId, 1)
	c.Equal(res.Email, "test@mail.com")
}
