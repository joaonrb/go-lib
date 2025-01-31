package convertto_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/joaonrb/go-lib/convertto"
	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionValueShouldConvertCorrectlyWhenDoingCorrectOperation(t *testing.T) {
	var stringOption monad.Maybe[string] = monad.Some[string]{Value: "10"}
	maybe := convertto.Maybe[string, int](stringOption).Then(mustIntConverter)
	require.IsTypef(
		t,
		monad.Some[int]{},
		maybe,
		"maybe is expected to Some[int], got %T instead",
		maybe,
	)
	value := maybe.(monad.Some[int]).Value
	assert.Equal(t, 10, value, "maybe is expected to be 10, got %d instead", value)
}

func TestOptionValueShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringOption monad.Maybe[string] = monad.Some[string]{Value: "1gg0"}
	maybe := convertto.Maybe[string, int](stringOption).Then(mustIntConverter)
	require.IsTypef(
		t,
		monad.Nothing[int]{},
		maybe,
		"maybe is expected to Nothing[int], got %T instead",
		maybe,
	)
}

func TestOptionNothingShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringOption monad.Maybe[string] = monad.Nothing[string]{}
	maybe := convertto.Maybe[string, int](stringOption).Then(mustIntConverter)
	require.IsTypef(
		t,
		monad.Nothing[int]{},
		maybe,
		"maybe is expected to Nothing[int], got %T instead",
		maybe,
	)
}

func TestOptionOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intOption monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.Equal(
		t,
		"Some[int, string]{Value: 10}",
		fmt.Sprint(convertto.Maybe[int, string](intOption)),
	)
	var stringOption monad.Maybe[string] = monad.Some[string]{Value: "10"}
	assert.Equal(
		t,
		"Some[string, int]{Value: \"10\"}",
		fmt.Sprint(convertto.Maybe[string, int](stringOption)),
	)
	var errorResult monad.Maybe[int] = monad.Nothing[int]{}
	assert.Equal(
		t,
		"Nothing[int, string]{}",
		fmt.Sprint(convertto.Maybe[int, string](errorResult)),
	)
}

func mustIntConverter(number string) monad.Maybe[int] {
	v, e := strconv.Atoi(number)
	if e != nil {
		return monad.Nothing[int]{}
	}
	return monad.Some[int]{Value: v}
}
