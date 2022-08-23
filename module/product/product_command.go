package product

import (
	"catalyst-case/database"
	"context"
	"errors"
)

type productCommand struct {
	*database.DB
}

const (
	addProduct = `INSERT INTO products (brand_id, product_name, price, created_at) VALUES (?, ?, ?, ?)`
)

type ProductCommand interface {
	AddProduct(ctx context.Context, product ProductEntity) error
}

func NewProductCommand(db *database.DB) ProductCommand {
	return &productCommand{
		DB: db,
	}
}

func (c *productCommand) AddProduct(ctx context.Context, product ProductEntity) error {
	r, e := c.ExecContext(ctx, addProduct, product.BrandId, product.ProductName, product.Price, product.CreatedAt)
	if e != nil {
		return e
	}

	row, e := r.RowsAffected()
	if e != nil {
		return e
	}

	if row < 1 {
		return errors.New("couldn't add new product")
	}
	return nil
}
