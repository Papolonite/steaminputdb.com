package logging

import (
	"context"
	"log/slog"
	"time"

	"github.com/uptrace/bun"
)

type queryHook struct{}

// NewQueryHook creates a new query hook that logs executed SQL queries.
func NewQueryHook() *queryHook {
	return &queryHook{}
}

func (qh *queryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (qh *queryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	duration := time.Since(event.StartTime)

	if event.Err != nil {
		slog.Error("SQL Query Error",
			"query", event.Query,
			"args", event.QueryArgs,
			"duration", duration,
			"error", event.Err,
		)
		return
	}

	slog.Debug("SQL Query",
		"query", event.Query,
		"args", event.QueryArgs,
		"result", event.Result,
		"duration", duration,
	)
}
