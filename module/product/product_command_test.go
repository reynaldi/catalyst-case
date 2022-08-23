package product

import (
	"catalyst-case/database"
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type productCommandTestSuite struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	db               *database.DB
	ctx              context.Context
	productStructure []string
}

func TestProductCommandSuite(t *testing.T) {
	suite.Run(t, new(productCommandTestSuite))
}

func (p *productCommandTestSuite) SetupTest() {
	p.ctx = context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		p.Fail("fail to open stub db")
	}
	p.db = &database.DB{
		DB:      db,
		Dialect: "mysql",
	}
	p.mock = mock
	p.productStructure = []string{
		"product_id",
		"brand_id",
		"product_name",
		"price",
		"created_at",
		"deleted_at",
	}
}

func (p *productCommandTestSuite) TestAddProduct_ReturnError() {
	now := time.Now().UTC()
	p.mock.ExpectExec(regexp.QuoteMeta(addProduct)).WithArgs(2, "test1", float64(15000), now).WillReturnError(errors.New("error db"))
	c := NewProductCommand(p.db)
	e := c.AddProduct(p.ctx, ProductEntity{
		ProductId:   1,
		BrandId:     2,
		ProductName: "test1",
		CreatedAt:   now,
		Price:       float64(15000),
	})
	p.NotNil(e)
	p.EqualError(e, "error db")
}

func (p *productCommandTestSuite) TestAddProduct_ReturnOk() {
	now := time.Now().UTC()
	p.mock.ExpectExec(regexp.QuoteMeta(addProduct)).WithArgs(2, "test1", float64(15000), now).WillReturnResult(sqlmock.NewResult(0, 1))
	c := NewProductCommand(p.db)
	e := c.AddProduct(p.ctx, ProductEntity{
		ProductId:   1,
		BrandId:     2,
		ProductName: "test1",
		CreatedAt:   now,
		Price:       float64(15000),
	})
	p.Nil(e)
}
