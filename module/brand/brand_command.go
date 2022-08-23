package brand

import (
	"catalyst-case/database"
	"context"
	"errors"
)

const (
	addBrand = `INSERT INTO brands (brand_name, created_at, deleted_at) VALUES (?, ?, NULL)`
)

type brandCommand struct {
	*database.DB
}

type BrandCommand interface {
	AddBrand(ctx context.Context, brand BrandEntity) error
}

func NewBrandCommand(db *database.DB) BrandCommand {
	return &brandCommand{
		DB: db,
	}
}

func (b *brandCommand) AddBrand(ctx context.Context, brand BrandEntity) error {
	res, e := b.ExecContext(ctx, addBrand, brand.BrandName, brand.CreatedAt)
	if e != nil {
		return e
	}

	affected, e := res.RowsAffected()
	if e != nil {
		return e
	}

	if affected < 1 {
		return errors.New("couldn't add new brand")
	}

	return nil
}
