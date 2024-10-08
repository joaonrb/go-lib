package convert_test

import (
	"testing"

	"github.com/joaonrb/go-lib/convert"
	"github.com/stretchr/testify/require"
)

func TestValueShouldReturnTheValueOfThePointer(t *testing.T) {
	t.Parallel()
	pointer := convert.Pointer("João")
	value := convert.Value(pointer)

	require.Equal(t, value, *pointer)
	value += " Nuno"
	require.Equal(t, "João", *pointer)
	require.Equal(t, "João Nuno", value)
}
