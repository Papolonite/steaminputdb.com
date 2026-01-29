package migrations

import (
	"context"
	"log/slog"

	"github.com/Alia5/steaminputdb.com/db/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		slog.Info("running init migration")

		_, err := db.NewCreateTable().Model((*models.ControllerType)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateIndex().Model((*models.ControllerType)(nil)).Column("name").Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}, nil)
}
