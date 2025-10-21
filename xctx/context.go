package xctx

import (
	"context"

	"go.uber.org/zap"
)

type contextHolderKey int

const contextHolderKeyValue contextHolderKey = 1

type contextHolder struct {
	logger *zap.Logger
}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	holder := &contextHolder{}
	if val := ctx.Value(contextHolderKeyValue); val != nil {
		parentHolder, ok := val.(*contextHolder)
		if !ok {
			panic("unexpected value in context")
		}
		*holder = *parentHolder
	}
	holder.logger = logger
	return context.WithValue(ctx, contextHolderKeyValue, holder)
}

func Logger(ctx context.Context) *zap.Logger {
	if val, ok := ctx.Value(contextHolderKeyValue).(*contextHolder); ok && val != nil && val.logger != nil {
		return val.logger
	}
	return zap.NewNop()
}
