package convert_test

import (
	"testing"

	"github.com/joaonrb/go-lib/convert"

	"github.com/stretchr/testify/require"
)

func TestToPointerShouldReturnAPointerOfTheSameValue(t *testing.T) {
	t.Parallel()
	value := "João"
	pointer := convert.ToPointer(value)
	require.Equal(t, value, *pointer)
	*pointer += " Nuno"
	require.Equal(t, "João", value)
	require.Equal(t, "João Nuno", *pointer)
}
