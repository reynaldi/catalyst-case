package brand

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type BrandQueryMock struct {
	mock.Mock
}

func (b *BrandQueryMock) GetBrandById(ctx context.Context, brandId int) (*BrandEntity, error) {
	ret := b.Called(ctx, brandId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*BrandEntity), ret.Error(1)
}

func (b *BrandQueryMock) GetBrands(ctx context.Context) ([]BrandEntity, error) {
	ret := b.Called(ctx)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).([]BrandEntity), ret.Error(1)
}
