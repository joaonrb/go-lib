package convertto_test

import (
	"testing"

	"github.com/joaonrb/go-lib/convertto"
	"github.com/stretchr/testify/require"
)

func TestValueShouldReturnTheValueOfThePointer(t *testing.T) {
	t.Parallel()
	pointer := convertto.Pointer("João")
	value := convertto.Value(pointer)

	require.Equal(t, value, *pointer)
	value += " Nuno"
	require.Equal(t, "João", *pointer)
	require.Equal(t, "João Nuno", value)
}
