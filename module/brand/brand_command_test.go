package brand

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

type brandCommandTestSuite struct {
	suite.Suite
	ctx            context.Context
	db             *database.DB
	mock           sqlmock.Sqlmock
	brandStructure []string
}

func TestBrandCommandTestSuite(t *testing.T) {
	suite.Run(t, new(brandCommandTestSuite))
}

func (b *brandCommandTestSuite) SetupTest() {
	b.ctx = context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		b.Fail("failed to open stub db")
	}
	b.db = &database.DB{
		DB:      db,
		Dialect: "mysql",
	}
	b.mock = mock
	b.brandStructure = []string{
		"brand_id",
		"brand_name",
		"created_at",
		"deleted_at",
	}
}

func (b *brandCommandTestSuite) TestAddBrand_ExecReturnError() {
	var brandName = "test brand"
	var createdAt = time.Now().UTC()
	b.mock.ExpectExec(regexp.QuoteMeta(addBrand)).
		WithArgs(brandName, createdAt).
		WillReturnError(errors.New("error db"))
	var command = NewBrandCommand(b.db)
	var e = command.AddBrand(b.ctx, BrandEntity{
		BrandName: brandName,
		CreatedAt: createdAt,
	})
	b.NotNil(e)
	b.EqualError(e, "error db")
}

func (b *brandCommandTestSuite) TestAddBrand_RowsAffectedError() {
	var brandName = "test brand"
	var createdAt = time.Now().UTC()
	b.mock.ExpectExec(regexp.QuoteMeta(addBrand)).
		WithArgs(brandName, createdAt).
		WillReturnResult(sqlmock.NewResult(0, 0))
	var command = NewBrandCommand(b.db)
	var e = command.AddBrand(b.ctx, BrandEntity{
		BrandName: brandName,
		CreatedAt: createdAt,
	})
	b.NotNil(e)
	b.EqualError(e, "couldn't add new brand")
}

func (b *brandCommandTestSuite) TestAddBrand_Ok() {
	var brandName = "test brand"
	var createdAt = time.Now().UTC()
	b.mock.ExpectExec(regexp.QuoteMeta(addBrand)).
		WithArgs(brandName, createdAt).
		WillReturnResult(sqlmock.NewResult(0, 1))
	var command = NewBrandCommand(b.db)
	var e = command.AddBrand(b.ctx, BrandEntity{
		BrandName: brandName,
		CreatedAt: createdAt,
	})
	b.Nil(e)
}
