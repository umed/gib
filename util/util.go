package util

import (
	"context"
	"log/slog"

	"github.com/umed/gib/lg"
	"github.com/umed/gib/xctx"
)

func DefaultRecover(ctx context.Context) {
	r := recover()
	if r == nil {
		return
	}
	if err, ok := r.(error); ok {
		xctx.Logger(ctx).Error("recovered panic", lg.Err(err))
	} else {
		xctx.Logger(ctx).Error("recovered panic", slog.Any("recover", err))
	}
}

func CustomRecover(ctx context.Context, fn func(ctx context.Context, r interface{})) {
	r := recover()
	if r == nil {
		return
	}
	xctx.Logger(ctx).Debug("custom recover triggered", slog.Any("recover", r))
	fn(ctx, r)
}
