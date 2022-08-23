package brand

import "time"

type BrandEntity struct {
	BrandId   int        `db:"brand_id"`
	BrandName string     `db:"brand_name"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type NewBrandDto struct {
	BrandName string `json:"brand_name" validate:"required"`
}

type BrandDto struct {
	BrandId   int    `json:"brand_id"`
	BrandName string `json:"brand_name"`
}
