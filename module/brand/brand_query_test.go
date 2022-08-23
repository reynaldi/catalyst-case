package brand

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

type brandQueryTestSuite struct {
	suite.Suite
	ctx            context.Context
	db             *database.DB
	mock           sqlmock.Sqlmock
	brandStructure []string
}

func TestBrandQueryTestSuite(t *testing.T) {
	suite.Run(t, new(brandQueryTestSuite))
}

func (b *brandQueryTestSuite) SetupTest() {
	b.brandStructure = []string{
		"brand_id",
		"brand_name",
		"created_at",
		"deleted_at",
	}
	b.ctx = context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		b.Fail("fail to open stub db")
	}
	b.db = &database.DB{
		DB:      db,
		Dialect: "mysql",
	}
	b.mock = mock
}

func (b *brandQueryTestSuite) TestGetBrandById_ReturnErr() {
	b.mock.ExpectQuery(regexp.QuoteMeta(getBrandById)).
		WithArgs(1).
		WillReturnError(errors.New("new error"))
	var svc = NewBrandQuery(b.db)
	res, e := svc.GetBrandById(b.ctx, 1)
	b.Nil(res)
	b.EqualError(e, "new error")
}

func (b *brandQueryTestSuite) TestGetBrandById_ReturnNoRow() {
	b.mock.ExpectQuery(regexp.QuoteMeta(getBrandById)).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)
	var svc = NewBrandQuery(b.db)
	res, e := svc.GetBrandById(b.ctx, 1)
	b.Nil(res)
	b.Nil(e)
}

func (b *brandQueryTestSuite) TestGetBrandById_ReturnOk() {
	b.mock.ExpectQuery(regexp.QuoteMeta(getBrandById)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(b.brandStructure).AddRow(1, "test brand", time.Now().UTC(), nil))
	var svc = NewBrandQuery(b.db)
	res, e := svc.GetBrandById(b.ctx, 1)
	b.NotNil(res)
	b.Equal("test brand", res.BrandName)
	b.Nil(e)
}

func (b *brandQueryTestSuite) TestGetBrands_ReturnErrorQuery() {
	b.mock.ExpectQuery(regexp.QuoteMeta(getAllBrands)).WillReturnError(errors.New("new error"))
	var svc = NewBrandQuery(b.db)
	res, e := svc.GetBrands(b.ctx)
	b.EqualError(e, "new error")
	b.Nil(res)
}

func (b *brandQueryTestSuite) TestGetBrands_ReturnScanError() {
	b.mock.ExpectQuery(regexp.QuoteMeta(getAllBrands)).
		WillReturnRows(sqlmock.NewRows(b.brandStructure).AddRow(
			1, "test brand", time.Now().UTC(), nil,
		).RowError(0, errors.New("row error")))
	var svc = NewBrandQuery(b.db)
	res, e := svc.GetBrands(b.ctx)
	b.EqualError(e, "row error")
	b.Nil(res)
}

func (b *brandQueryTestSuite) TestGetBrands_ReturnScanFormatWrong() {
	b.mock.ExpectQuery(regexp.QuoteMeta(getAllBrands)).
		WillReturnRows(sqlmock.NewRows(b.brandStructure).AddRow(
			"one", "test brand", time.Now().UTC(), nil,
		))
	var svc = NewBrandQuery(b.db)
	res, e := svc.GetBrands(b.ctx)
	b.Error(e)
	b.Nil(res)
}

func (b *brandQueryTestSuite) TestGetBrands_ReturnOk() {
	b.mock.ExpectQuery(regexp.QuoteMeta(getAllBrands)).
		WillReturnRows(sqlmock.NewRows(b.brandStructure).AddRow(
			1, "test brand", time.Now().UTC(), nil,
		))
	var svc = NewBrandQuery(b.db)
	res, e := svc.GetBrands(b.ctx)
	b.Nil(e)
	b.Equal("test brand", res[0].BrandName)
}
