package db

import (
	"context"
	"database/sql"

	"github.com/Alia5/steaminputdb.com/config"
	"github.com/Alia5/steaminputdb.com/db/migrations"
	"github.com/Alia5/steaminputdb.com/logging"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Init(cfg config.DB) error {

	ctx := context.Background()

	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(cfg.DatabaseURL),
	))

	sqldb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqldb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqldb.SetConnMaxLifetime(cfg.MaxConnLifetime)
	sqldb.SetConnMaxIdleTime(cfg.MaxConnIdleTime)

	db := bun.NewDB(sqldb, pgdialect.New()).
		WithQueryHook(logging.NewQueryHook())

	err := migrations.Migrate(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
