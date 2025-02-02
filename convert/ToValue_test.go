package convert_test

import (
	"testing"

	"github.com/joaonrb/go-lib/convert"

	"github.com/stretchr/testify/require"
)

func TestToValueShouldReturnTheValueOfThePointer(t *testing.T) {
	t.Parallel()
	pointer := convert.ToPointer("João")
	value := convert.ToValue(pointer)

	require.Equal(t, value, *pointer)
	value += " Nuno"
	require.Equal(t, "João", *pointer)
	require.Equal(t, "João Nuno", value)
}
