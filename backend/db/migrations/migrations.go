package migrations

import (
	"context"
	"embed"
	"log/slog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

//go:embed [0-9]*.*
var migrationFiles embed.FS

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(migrationFiles); err != nil {
		panic(err)
	}
}

func Migrate(ctx context.Context, db *bun.DB) error {
	migrator := migrate.NewMigrator(db, Migrations)
	err := migrator.Init(ctx)
	if err != nil {
		return err
	}
	if err := migrator.Lock(ctx); err != nil {
		return err
	}
	defer migrator.Unlock(ctx) //nolint:errcheck

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		slog.Debug("there are no new migrations to run (database is up to date)")
		return nil
	}
	slog.Info("migrated", "to", group)
	return nil
}
