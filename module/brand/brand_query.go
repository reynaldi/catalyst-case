package brand

import (
	"catalyst-case/database"
	"context"
	"database/sql"
	"log"
)

const (
	getBrandById = `SELECT * FROM brands WHERE brand_id = ? AND deleted_at IS NULL`
	getAllBrands = `SELECT * FROM brands WHERE deleted_at IS NULL`
)

type brandQuery struct {
	*database.DB
}

type BrandQuery interface {
	GetBrandById(ctx context.Context, brandId int) (*BrandEntity, error)
	GetBrands(ctx context.Context) ([]BrandEntity, error)
}

func NewBrandQuery(db *database.DB) BrandQuery {
	return &brandQuery{
		DB: db,
	}
}

func (b *brandQuery) GetBrandById(ctx context.Context, brandId int) (*BrandEntity, error) {
	row := b.QueryRowContext(ctx, getBrandById, brandId)
	result, e := ScanBrand(row)
	return result, e
}

func (b *brandQuery) GetBrands(ctx context.Context) ([]BrandEntity, error) {
	rows, e := b.QueryContext(ctx, getAllBrands)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	var result []BrandEntity
	for rows.Next() {
		var item BrandEntity
		if err := rows.Scan(&item.BrandId, &item.BrandName, &item.CreatedAt, &item.DeletedAt); err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, item)
	}
	e = rows.Err()
	if e != nil {
		return nil, e
	}
	return result, nil
}

func ScanBrand(row *sql.Row) (*BrandEntity, error) {
	var result BrandEntity
	e := row.Scan(&result.BrandId, &result.BrandName, &result.CreatedAt, &result.DeletedAt)
	if e != nil {
		if e == sql.ErrNoRows {
			return nil, nil
		}
		return nil, e
	}
	return &result, nil
}
