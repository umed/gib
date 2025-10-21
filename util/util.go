package util

import (
	"context"

	"github.com/umed/gib/xctx"
	"go.uber.org/zap"
)

func DefaultRecover(ctx context.Context) {
	r := recover()
	if r == nil {
		return
	}
	if err, ok := r.(error); ok {
		xctx.Logger(ctx).Error("recovered panic", zap.Error(err))
	} else {
		xctx.Logger(ctx).Error("recovered panic", zap.Any("recover", err))
	}
}

func CustomRecover(ctx context.Context, fn func(ctx context.Context, r any)) {
	r := recover()
	if r == nil {
		return
	}
	xctx.Logger(ctx).Debug("custom recover triggered", zap.Any("recover", r))
	fn(ctx, r)
}
