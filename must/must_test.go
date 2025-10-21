package must_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/umed/gib/must"
)

func TestMust(t *testing.T) {
	const expected = 42
	result := must.Must(42, nil)
	require.Equal(t, expected, result)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but did not get one")
		}
	}()
	err := errors.New("an error occurred")
	require.PanicsWithError(t, err.Error(), func() {
		must.Must(0, err)
	})
}
