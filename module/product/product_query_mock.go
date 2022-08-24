package product

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ProductQueryMock struct {
	mock.Mock
}

func (q *ProductQueryMock) GetProductById(ctx context.Context, productId int) (*ProductDto, error) {
	ret := q.Called(ctx, productId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*ProductDto), ret.Error(1)
}

func (q *ProductQueryMock) GetProductsByBrand(ctx context.Context, brandId int) ([]ProductDto, error) {
	ret := q.Called(ctx, brandId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).([]ProductDto), ret.Error(1)
}
