package brand

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type brandTestSuite struct {
	suite.Suite
	queryMock   *BrandQueryMock
	commandMock *BrandCommandMock
	ctx         context.Context
}

func TestBrandTestSuite(t *testing.T) {
	suite.Run(t, new(brandTestSuite))
}

func (b *brandTestSuite) SetupTest() {
	b.commandMock = new(BrandCommandMock)
	b.queryMock = new(BrandQueryMock)
	b.ctx = context.Background()
}

func (b *brandTestSuite) TestAddNewBrand_Error() {
	b.commandMock.On("AddBrand", b.ctx, mock.Anything).Return(errors.New("new error"))
	var srv = NewBrand(b.queryMock, b.commandMock)
	var e = srv.AddNewBrand(b.ctx, NewBrandDto{})
	b.EqualError(e, "new error")
}

func (b *brandTestSuite) TestAddNewBrand_Ok() {
	b.commandMock.On("AddBrand", b.ctx, mock.Anything).Return(nil)
	var srv = NewBrand(b.queryMock, b.commandMock)
	var e = srv.AddNewBrand(b.ctx, NewBrandDto{})
	b.Nil(e)
}

func (b *brandTestSuite) TestGetBrandById_ReturnError() {
	b.queryMock.On("GetBrandById", b.ctx, 1).Return(nil, errors.New("test error"))
	var svc = NewBrand(b.queryMock, b.commandMock)
	res, e := svc.GetBrandById(b.ctx, 1)
	b.Nil(res)
	b.EqualError(e, "test error")
}

func (b *brandTestSuite) TestGetBrandById_ResultNil() {
	b.queryMock.On("GetBrandById", b.ctx, 1).Return(nil, nil)
	var svc = NewBrand(b.queryMock, b.commandMock)
	res, e := svc.GetBrandById(b.ctx, 1)
	b.Nil(res)
	b.Nil(e)
}

func (b *brandTestSuite) TestGetBrandById_ResultOk() {
	b.queryMock.On("GetBrandById", b.ctx, 1).Return(&BrandEntity{
		BrandId:   1,
		BrandName: "test brand",
		CreatedAt: time.Now().UTC(),
	}, nil)
	var svc = NewBrand(b.queryMock, b.commandMock)
	res, e := svc.GetBrandById(b.ctx, 1)
	b.NotNil(res)
	b.Nil(e)
	b.Equal("test brand", res.BrandName)
	b.Equal(1, res.BrandId)
}

func (b *brandTestSuite) TestGetAllBrands_ReturnError() {
	b.queryMock.On("GetBrands", b.ctx).Return(nil, errors.New("all brands error"))
	svc := NewBrand(b.queryMock, b.commandMock)
	res, e := svc.GetAllBrands(b.ctx)
	b.Nil(res)
	b.EqualError(e, "all brands error")
}

func (b *brandTestSuite) TestGetAllBrands_ReturnNil() {
	b.queryMock.On("GetBrands", b.ctx).Return(nil, nil)
	svc := NewBrand(b.queryMock, b.commandMock)
	res, e := svc.GetAllBrands(b.ctx)
	b.Nil(res)
	b.Nil(e)
}

func (b *brandTestSuite) TestGetAllBrands_ReturnOk() {
	b.queryMock.On("GetBrands", b.ctx).Return([]BrandEntity{
		{
			BrandId:   1,
			BrandName: "test1",
			CreatedAt: time.Now().UTC(),
		},
		{
			BrandId:   2,
			BrandName: "test2",
			CreatedAt: time.Now().UTC(),
		},
	}, nil)
	svc := NewBrand(b.queryMock, b.commandMock)
	res, e := svc.GetAllBrands(b.ctx)
	b.Nil(e)
	b.Equal("test1", res[0].BrandName)
	b.Equal("test2", res[1].BrandName)
}
