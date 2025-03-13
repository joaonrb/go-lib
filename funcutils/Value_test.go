package funcutils_test

import (
	"github.com/joaonrb/go-lib/funcutils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValueShouldReturnTheValueOfThePointer(t *testing.T) {
	t.Parallel()
	pointer := funcutils.Pointer("João")
	value := funcutils.Value(pointer)

	require.Equal(t, value, *pointer)
	value += " Nuno"
	require.Equal(t, "João", *pointer)
	require.Equal(t, "João Nuno", value)
}
