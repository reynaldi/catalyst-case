package customer

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type customerQueryMock struct {
	mock.Mock
}

func (c *customerQueryMock) GetCustomerById(ctx context.Context, customerId int) (*CustomerEntity, error) {
	ret := c.Called(ctx, customerId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*CustomerEntity), ret.Error(1)
}

func (c *customerQueryMock) GetCustomerByEmail(ctx context.Context, email string) (*CustomerEntity, error) {
	ret := c.Called(ctx, email)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*CustomerEntity), ret.Error(1)
}
