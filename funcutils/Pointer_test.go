package funcutils_test

import (
	"github.com/joaonrb/go-lib/funcutils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPointerShouldReturnAPointerOfTheSameValue(t *testing.T) {
	t.Parallel()
	value := "João"
	pointer := funcutils.Pointer(value)
	require.Equal(t, value, *pointer)
	*pointer += " Nuno"
	require.Equal(t, "João", value)
	require.Equal(t, "João Nuno", *pointer)
}
