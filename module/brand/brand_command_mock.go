package brand

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type BrandCommandMock struct {
	mock.Mock
}

func (b BrandCommandMock) AddBrand(ctx context.Context, brand BrandEntity) error {
	ret := b.Called(ctx, brand)
	return ret.Error(0)
}
