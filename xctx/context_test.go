package xctx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/umed/gib/xctx"
	"go.uber.org/zap"
)

func TestWithLogger(t *testing.T) {
	logger := zap.NewNop()
	ctx := xctx.WithLogger(context.Background(), logger)
	require.Equal(t, logger, xctx.Logger(ctx))
}
