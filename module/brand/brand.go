package brand

import (
	"context"
	"time"
)

type brand struct {
	query   BrandQuery
	command BrandCommand
}

type Brand interface {
	AddNewBrand(ctx context.Context, brand NewBrandDto) error
	GetBrandById(ctx context.Context, brandId int) (*BrandDto, error)
	GetAllBrands(ctx context.Context) ([]BrandDto, error)
}

func NewBrand(query BrandQuery, command BrandCommand) Brand {
	return &brand{
		query:   query,
		command: command,
	}
}

func (b *brand) AddNewBrand(ctx context.Context, brand NewBrandDto) error {
	var entity = mapNewBrand(brand)
	var e = b.command.AddBrand(ctx, entity)
	return e
}

func (b *brand) GetBrandById(ctx context.Context, brandId int) (*BrandDto, error) {
	res, e := b.query.GetBrandById(ctx, brandId)
	if e != nil {
		return nil, e
	}
	if res == nil {
		return nil, nil
	}
	var result = mapBrandToDto(res)
	return &result, nil
}

func (b *brand) GetAllBrands(ctx context.Context) ([]BrandDto, error) {
	res, e := b.query.GetBrands(ctx)
	if e != nil {
		return nil, e
	}
	var result []BrandDto
	for _, item := range res {
		b := mapBrandToDto(&item)
		result = append(result, b)
	}

	return result, nil
}

func mapNewBrand(brand NewBrandDto) BrandEntity {
	return BrandEntity{
		BrandName: brand.BrandName,
		CreatedAt: time.Now().UTC(),
	}
}

func mapBrandToDto(brand *BrandEntity) BrandDto {
	return BrandDto{
		BrandId:   brand.BrandId,
		BrandName: brand.BrandName,
	}
}
