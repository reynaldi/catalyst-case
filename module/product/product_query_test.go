package product

import (
	"catalyst-case/database"
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type productQueryTestSuite struct {
	suite.Suite
	db               *database.DB
	mock             sqlmock.Sqlmock
	ctx              context.Context
	productStructure []string
}

func TestProductQueryTestSuite(t *testing.T) {
	suite.Run(t, new(productQueryTestSuite))
}

func (p *productQueryTestSuite) SetupTest() {
	p.ctx = context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		p.Fail("Fail to open stub db")
	}
	p.db = &database.DB{
		DB:      db,
		Dialect: "mysql",
	}
	p.mock = mock
	p.productStructure = []string{
		"product_id",
		"product_name",
		"brand_id",
		"brand_name",
		"price",
		"created_at",
	}
}

func (p *productQueryTestSuite) TestGetProductById_ReturnError() {
	p.mock.ExpectQuery(regexp.QuoteMeta(getProductById)).WithArgs(1).WillReturnError(errors.New("query error"))
	q := NewProductQuery(p.db)
	r, e := q.GetProductById(p.ctx, 1)
	p.Nil(r)
	p.EqualError(e, "query error")
}

func (p *productQueryTestSuite) TestGetProductById_ReturnNil() {
	p.mock.ExpectQuery(regexp.QuoteMeta(getProductById)).WithArgs(1).WillReturnError(sql.ErrNoRows)
	q := NewProductQuery(p.db)
	r, e := q.GetProductById(p.ctx, 1)
	p.Nil(r)
	p.Nil(e)
}

func (p *productQueryTestSuite) TestGetProductById_ReturnOk() {
	p.mock.ExpectQuery(regexp.QuoteMeta(getProductById)).WithArgs(1).
		WillReturnRows(sqlmock.NewRows(p.productStructure).AddRow(1, "test1 product", 2, "test2 brand", 15000, time.Now().UTC()))
	q := NewProductQuery(p.db)
	r, e := q.GetProductById(p.ctx, 1)
	p.Nil(e)
	p.Equal("test1 product", r.ProductName)
	p.Equal(float64(15000), r.Price)
	p.Equal(2, r.BrandId)
}

func (p *productQueryTestSuite) TestGetProductsByBrands_ReturnError() {
	p.mock.ExpectQuery(regexp.QuoteMeta(getProductsByBrand)).WithArgs(2).WillReturnError(errors.New("error db"))
	q := NewProductQuery(p.db)
	r, e := q.GetProductsByBrand(p.ctx, 2)
	p.Nil(r)
	p.EqualError(e, "error db")
}

func (p *productQueryTestSuite) TestGetProductsByBrands_ReturnOk() {
	p.mock.ExpectQuery(regexp.QuoteMeta(getProductsByBrand)).WithArgs(2).WillReturnRows(sqlmock.NewRows(p.productStructure).AddRow(
		1, "test1 product", 2, "test1 brand", 15000, time.Now().UTC(),
	).AddRow(
		2, "test2 product", 2, "test1 brand", 25000, time.Now().UTC(),
	))
	q := NewProductQuery(p.db)
	r, e := q.GetProductsByBrand(p.ctx, 2)
	p.Nil(e)
	p.Equal(2, len(r))
	p.Equal("test1 product", r[0].ProductName)
	p.Equal("test2 product", r[1].ProductName)
}
