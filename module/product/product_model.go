package product

import "time"

type ProductEntity struct {
	ProductId   int        `db:"product_id"`
	BrandId     int        `db:"brand_id"`
	ProductName string     `db:"product_name"`
	Price       float64    `db:"price"`
	CreatedAt   time.Time  `db:"created_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type NewProductDto struct {
	BrandId     int     `json:"brand_id" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type ProductDto struct {
	ProductId   int       `json:"product_id"`
	BrandId     int       `json:"brand_id"`
	BrandName   string    `json:"brand_name"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}
