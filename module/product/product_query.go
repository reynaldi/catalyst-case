package product

import (
	"catalyst-case/database"
	"context"
	"database/sql"
	"log"
)

const (
	getProductById = `SELECT p.product_id, p.product_name, p.brand_id, b.brand_name, p.price, p.created_at
						FROM products p JOIN brands b ON b.brand_id = p.brand_id 
						WHERE p.product_id = ? AND p.deleted_at IS NULL`
	getProductsByBrand = `SELECT p.product_id, p.product_name, p.brand_id, b.brand_name, p.price, p.created_at
							FROM products p JOIN brands b ON b.brand_id = p.brand_id 
							WHERE p.brand_id = ? AND p.deleted_at IS NULL`
)

type productQuery struct {
	*database.DB
}

type ProductQuery interface {
	GetProductById(ctx context.Context, productId int) (*ProductDto, error)
	GetProductsByBrand(ctx context.Context, brandId int) ([]ProductDto, error)
}

func NewProductQuery(db *database.DB) ProductQuery {
	return &productQuery{
		DB: db,
	}
}

func (p *productQuery) GetProductById(ctx context.Context, productId int) (*ProductDto, error) {
	row := p.QueryRowContext(ctx, getProductById, productId)
	result, e := ScanProduct(row)
	return result, e
}

func (p *productQuery) GetProductsByBrand(ctx context.Context, brandId int) ([]ProductDto, error) {
	rows, e := p.QueryContext(ctx, getProductsByBrand, brandId)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	var res []ProductDto
	for rows.Next() {
		var result ProductDto
		if err := rows.Scan(&result.ProductId, &result.ProductName, &result.BrandId, &result.BrandName, &result.Price, &result.CreatedAt); err != nil {
			log.Println(err)
			return nil, err
		}
		res = append(res, result)
	}
	e = rows.Err()
	if e != nil {
		return nil, e
	}
	return res, nil
}

func ScanProduct(row *sql.Row) (*ProductDto, error) {
	var result ProductDto
	e := row.Scan(&result.ProductId, &result.ProductName, &result.BrandId, &result.BrandName, &result.Price, &result.CreatedAt)
	if e != nil {
		if e == sql.ErrNoRows {
			return nil, nil
		}
		return nil, e
	}
	return &result, e
}
