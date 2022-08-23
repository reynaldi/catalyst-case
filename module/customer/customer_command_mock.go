package customer

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type customerCommandMock struct {
	mock.Mock
}

func (c *customerCommandMock) AddCustomer(ctx context.Context, customer CustomerEntity) error {
	ret := c.Called(ctx, customer)
	return ret.Error(0)
}
