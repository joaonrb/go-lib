package convertto_test

import (
	"testing"

	"github.com/joaonrb/go-lib/convertto"
	"github.com/stretchr/testify/require"
)

func TestPointerShouldReturnAPointerOfTheSameValue(t *testing.T) {
	t.Parallel()
	value := "João"
	pointer := convertto.Pointer(value)
	require.Equal(t, value, *pointer)
	*pointer += " Nuno"
	require.Equal(t, "João", value)
	require.Equal(t, "João Nuno", *pointer)
}
