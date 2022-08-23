package database

import (
	"catalyst-case/config"
	"database/sql"
	"errors"
)

type DB struct {
	*sql.DB
	Dialect string
}

func Open(cfg *config.Config) (*DB, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}

	if cfg.ConnectionString == "" {
		return nil, errors.New("database connection string is required")
	}

	dbConn, err := sql.Open(cfg.Dialect, cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	var db = &DB{
		DB:      dbConn,
		Dialect: cfg.Dialect,
	}
	return db, nil
}
