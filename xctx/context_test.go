package xctx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/umed/gib/lg"
	"github.com/umed/gib/xctx"
)

func TestWithLogger(t *testing.T) {
	logger := lg.New()
	ctx := xctx.WithLogger(context.Background(), logger)
	require.Equal(t, logger, xctx.Logger(ctx))
}
