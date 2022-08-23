package product

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type productTestSuite struct {
	suite.Suite
	queryMock   *productQueryMock
	commandMock *productCommandMock
	ctx         context.Context
}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(productTestSuite))
}

func (p *productTestSuite) SetupTest() {
	p.ctx = context.Background()
	p.queryMock = new(productQueryMock)
	p.commandMock = new(productCommandMock)
}

func (p *productTestSuite) TestAddNewProduct_ReturnError() {
	p.commandMock.On("AddProduct", p.ctx, mock.Anything).Return(errors.New("new error"))
	svc := NewProduct(p.queryMock, p.commandMock)
	e := svc.AddNewProduct(p.ctx, NewProductDto{})
	p.EqualError(e, "new error")
}

func (p *productTestSuite) TestAddNewProduct_ReturnOk() {
	p.commandMock.On("AddProduct", p.ctx, mock.Anything).Return(nil)
	svc := NewProduct(p.queryMock, p.commandMock)
	e := svc.AddNewProduct(p.ctx, NewProductDto{})
	p.Nil(e)
}

func (p *productTestSuite) TestGetProductById_ReturnError() {
	p.queryMock.On("GetProductById", p.ctx, 1).Return(nil, errors.New("new error"))
	svc := NewProduct(p.queryMock, p.commandMock)
	res, e := svc.GetProductById(p.ctx, 1)
	p.Nil(res)
	p.EqualError(e, "new error")
}

func (p *productTestSuite) TestGetProductById_ReturnOk() {
	p.queryMock.On("GetProductById", p.ctx, 1).Return(&ProductDto{
		ProductId:   1,
		BrandId:     2,
		BrandName:   "test brand",
		ProductName: "test product",
		Price:       float64(35000),
		CreatedAt:   time.Now().UTC(),
	}, nil)
	svc := NewProduct(p.queryMock, p.commandMock)
	res, e := svc.GetProductById(p.ctx, 1)
	p.Nil(e)
	p.Equal("test brand", res.BrandName)
	p.Equal("test product", res.ProductName)
	p.Equal(float64(35000), res.Price)
}

func (p *productTestSuite) TestGetProductsByBrand_ReturnError() {
	p.queryMock.On("GetProductsByBrand", p.ctx, 1).Return(nil, errors.New("new error"))
	svc := NewProduct(p.queryMock, p.commandMock)
	res, e := svc.GetProductsByBrand(p.ctx, 1)
	p.Nil(res)
	p.EqualError(e, "new error")
}

func (p *productTestSuite) TestGetProductsByBrand_ReturnOk() {
	p.queryMock.On("GetProductsByBrand", p.ctx, 1).Return([]ProductDto{
		{
			ProductId:   1,
			BrandId:     2,
			BrandName:   "test1 brand",
			ProductName: "test1 product",
			Price:       float64(60000),
			CreatedAt:   time.Now().UTC(),
		},
		{
			ProductId:   2,
			BrandId:     2,
			BrandName:   "test1 brand",
			ProductName: "test2 product",
			Price:       float64(100000),
			CreatedAt:   time.Now().UTC(),
		},
	}, nil)
	svc := NewProduct(p.queryMock, p.commandMock)
	res, e := svc.GetProductsByBrand(p.ctx, 1)
	p.Nil(e)
	p.Equal(2, len(res))
	p.Equal("test2 product", res[1].ProductName)
}
