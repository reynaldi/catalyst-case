package product

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type productQueryMock struct {
	mock.Mock
}

func (q *productQueryMock) GetProductById(ctx context.Context, productId int) (*ProductDto, error) {
	ret := q.Called(ctx, productId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*ProductDto), ret.Error(1)
}

func (q *productQueryMock) GetProductsByBrand(ctx context.Context, brandId int) ([]ProductDto, error) {
	ret := q.Called(ctx, brandId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).([]ProductDto), ret.Error(1)
}
