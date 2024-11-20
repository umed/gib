package xctx

import (
	"context"

	"github.com/umed/gib/logging"
)

type ctxKey struct{}

func WithLogger(ctx context.Context, logger *logging.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, &logger)
}

func Logger(ctx context.Context) *logging.Logger {
	if logger, ok := ctx.Value(ctxKey{}).(*logging.Logger); ok {
		return logger
	}
	return logging.DefaultLogger
}
