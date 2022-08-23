package database

import (
	"catalyst-case/config"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DatabaseTestSuite struct {
	suite.Suite
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}

func (d *DatabaseTestSuite) TestOpen_ConfigNil() {
	_, err := Open(nil)
	d.NotNil(err)
	d.EqualError(err, "config is required")
}

func (d *DatabaseTestSuite) TestOpen_ConnStringEmpty() {
	var setting = &config.Config{
		ConnectionString: "",
		Dialect:          "",
	}
	_, err := Open(setting)
	d.NotNil(err)
	d.EqualError(err, "database connection string is required")
}
