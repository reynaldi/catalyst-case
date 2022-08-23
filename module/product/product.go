package product

import (
	"context"
	"time"
)

type product struct {
	query   ProductQuery
	command ProductCommand
}

type Product interface {
	AddNewProduct(ctx context.Context, product NewProductDto) error
	GetProductById(ctx context.Context, productId int) (*ProductDto, error)
	GetProductsByBrand(ctx context.Context, brandId int) ([]ProductDto, error)
}

func NewProduct(query ProductQuery, command ProductCommand) Product {
	return &product{
		query:   query,
		command: command,
	}
}

func (p *product) AddNewProduct(ctx context.Context, product NewProductDto) error {
	entity := ProductEntity{
		BrandId:     product.BrandId,
		ProductName: product.ProductName,
		Price:       product.Price,
		CreatedAt:   time.Now().UTC(),
	}
	e := p.command.AddProduct(ctx, entity)
	return e
}

func (p *product) GetProductById(ctx context.Context, productId int) (*ProductDto, error) {
	res, e := p.query.GetProductById(ctx, productId)
	return res, e
}

func (p *product) GetProductsByBrand(ctx context.Context, brandId int) ([]ProductDto, error) {
	res, e := p.query.GetProductsByBrand(ctx, brandId)
	return res, e
}
