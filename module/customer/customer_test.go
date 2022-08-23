package customer

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type customerServiceTestSuite struct {
	suite.Suite
	queryMock   *customerQueryMock
	commandMock *customerCommandMock
	ctx         context.Context
}

func TestCustomerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(customerServiceTestSuite))
}

func (c *customerServiceTestSuite) SetupTest() {
	c.commandMock = new(customerCommandMock)
	c.queryMock = new(customerQueryMock)
	c.ctx = context.Background()
}

func (c *customerServiceTestSuite) TestAddNewCustomer_Error() {
	c.commandMock.On("AddCustomer", c.ctx, mock.Anything).Return(errors.New("error add customer"))
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	var e = svc.AddNewCustomer(c.ctx, NewCustomerDto{})
	c.EqualError(e, "error add customer")
}

func (c *customerServiceTestSuite) TestAddNewCustomer_Ok() {
	c.commandMock.On("AddCustomer", c.ctx, mock.Anything).Return(nil)
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	var e = svc.AddNewCustomer(c.ctx, NewCustomerDto{})
	c.Nil(e)
}

func (c *customerServiceTestSuite) TestGetById_ReturnError() {
	c.queryMock.On("GetCustomerById", c.ctx, 1).Return(nil, errors.New("error get by id"))
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	res, e := svc.GetCustomerById(c.ctx, 1)
	c.Nil(res)
	c.EqualError(e, "error get by id")
}

func (c *customerServiceTestSuite) TestGetById_ReturnCustomerNil() {
	c.queryMock.On("GetCustomerById", c.ctx, 1).Return(nil, nil)
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	res, e := svc.GetCustomerById(c.ctx, 1)
	c.Nil(res)
	c.Nil(e)
}

func (c *customerServiceTestSuite) TestGetById_ReturnOk() {
	var expected = &CustomerEntity{
		CustomerId: 1,
		Email:      "test@mail.com",
		Name:       "test",
	}
	c.queryMock.On("GetCustomerById", c.ctx, 1).Return(expected, nil)
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	res, e := svc.GetCustomerById(c.ctx, 1)
	c.NotNil(res)
	c.Equal(expected.Email, res.Email)
	c.Equal(expected.Name, res.Name)
	c.Nil(e)
}

func (c *customerServiceTestSuite) TestGetByEmail_ReturnError() {
	var email = "test@mail.com"
	c.queryMock.On("GetCustomerByEmail", c.ctx, email).Return(nil, errors.New("error get by email"))
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	res, e := svc.GetCustomerByEmail(c.ctx, email)
	c.Nil(res)
	c.EqualError(e, "error get by email")
}

func (c *customerServiceTestSuite) TestGetByEmail_ReturnCustomerNil() {
	var email = "test@mail.com"
	c.queryMock.On("GetCustomerByEmail", c.ctx, email).Return(nil, nil)
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	res, e := svc.GetCustomerByEmail(c.ctx, email)
	c.Nil(res)
	c.Nil(e)
}

func (c *customerServiceTestSuite) TestGetByEmail_ReturnOk() {
	var expected = &CustomerEntity{
		CustomerId: 1,
		Email:      "test@mail.com",
		Name:       "test",
	}
	c.queryMock.On("GetCustomerByEmail", c.ctx, expected.Email).Return(expected, nil)
	var svc = NewCustomerService(c.queryMock, c.commandMock)
	res, e := svc.GetCustomerByEmail(c.ctx, expected.Email)
	c.NotNil(res)
	c.Equal(expected.Email, res.Email)
	c.Equal(expected.Name, res.Name)
	c.Nil(e)
}
