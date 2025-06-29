package util_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/umed/gib/must"
)

func TestMust(t *testing.T) {
	const expectedErrorMessage = "test"
	require.PanicsWithError(t, expectedErrorMessage, func() {
		fn := func() (int, error) {
			return 0, errors.New(expectedErrorMessage)
		}
		_ = must.Must(fn())
	})

	require.NotPanics(t, func() {
		fn := func() (int, error) {
			return 0, nil
		}
		_ = must.Must(fn())
	})
}
