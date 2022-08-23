package product

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type productCommandMock struct {
	mock.Mock
}

func (c *productCommandMock) AddProduct(ctx context.Context, product ProductEntity) error {
	ret := c.Called(ctx, product)
	return ret.Error(0)
}
