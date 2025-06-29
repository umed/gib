package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/umed/gib/ptr"
)

func TestValueOr(t *testing.T) {
	var nullInt *int = nil
	require.Equal(t, 1, ptr.ValueOr(nullInt, 1))

	nonNullInt := ptr.Of(10)
	require.Equal(t, *nonNullInt, ptr.ValueOr(nonNullInt, -1))
}

func TestValue(t *testing.T) {
	var nullInt *int = nil
	require.Equal(t, 0, ptr.Value(nullInt))

	nonNullInt := ptr.Of(10)
	require.Equal(t, *nonNullInt, ptr.Value(nonNullInt))
}
