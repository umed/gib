package xctx

import (
	"context"

	"github.com/umed/gib/lg"
)

type contextHolderKey int

const contextHolderKeyValue contextHolderKey = 1

type contextHolder struct {
	logger *lg.Logger
}

func WithLogger(ctx context.Context, logger *lg.Logger) context.Context {
	holder := &contextHolder{}
	if val := ctx.Value(contextHolderKeyValue); val != nil {
		parentHolder, ok := val.(*contextHolder)
		if !ok {
			panic("unexpected value in context")
		}
		*holder = *parentHolder
	}
	holder.logger = logger
	return context.WithValue(ctx, contextHolderKeyValue, &holder)
}

func Logger(ctx context.Context) *lg.Logger {
	if val, ok := ctx.Value(contextHolderKeyValue).(*contextHolder); ok && val != nil && val.logger != nil {
		return val.logger
	}
	return &lg.NopLogger
}
