package convertto_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/joaonrb/go-lib/convertto"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionValueShouldConvertCorrectlyWhenDoingCorrectOperation(t *testing.T) {
	var stringOption types.Option[string] = types.Value[string]{This: "10"}
	option := convertto.Option[string, int](stringOption).Then(mustIntConverter)
	require.IsTypef(
		t,
		types.Value[int]{},
		option,
		"option is expected to Value[int], got %T instead",
		option,
	)
	value := option.(types.Value[int]).This
	assert.Equal(t, 10, value, "option is expected to be 10, got %d instead", value)
}

func TestOptionValueShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringOption types.Option[string] = types.Value[string]{This: "1gg0"}
	option := convertto.Option[string, int](stringOption).Then(mustIntConverter)
	require.IsTypef(
		t,
		types.Nothing[int]{},
		option,
		"option is expected to Nothing[int], got %T instead",
		option,
	)
}

func TestOptionNothingShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringOption types.Option[string] = types.Nothing[string]{}
	option := convertto.Option[string, int](stringOption).Then(mustIntConverter)
	require.IsTypef(
		t,
		types.Nothing[int]{},
		option,
		"option is expected to Nothing[int], got %T instead",
		option,
	)
}

func TestOptionOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intOption types.Option[int] = types.Value[int]{This: 10}
	assert.Equal(
		t,
		"Value[int, string]{This: 10}",
		fmt.Sprint(convertto.Option[int, string](intOption)),
	)
	var stringOption types.Option[string] = types.Value[string]{This: "10"}
	assert.Equal(
		t,
		"Value[string, int]{This: \"10\"}",
		fmt.Sprint(convertto.Option[string, int](stringOption)),
	)
	var errorResult types.Option[int] = types.Nothing[int]{}
	assert.Equal(
		t,
		"Nothing[int, string]{}",
		fmt.Sprint(convertto.Option[int, string](errorResult)),
	)
}

func mustIntConverter(number string) types.Option[int] {
	v, e := strconv.Atoi(number)
	if e != nil {
		return types.Nothing[int]{}
	}
	return types.Value[int]{This: v}
}
