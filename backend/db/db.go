package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/Alia5/steaminputdb.com/config"
	"github.com/Alia5/steaminputdb.com/db/migrations"
	"github.com/Alia5/steaminputdb.com/logging"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Init(cfg config.DB) error {

	ctx := context.Background()

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
	))

	sqldb.SetMaxOpenConns(25)
	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxLifetime(5 * time.Minute)
	sqldb.SetConnMaxIdleTime(5 * time.Minute)

	db := bun.NewDB(sqldb, pgdialect.New()).
		WithQueryHook(logging.NewQueryHook())

	err := migrations.Migrate(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
