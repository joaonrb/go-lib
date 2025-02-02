package convert_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/joaonrb/go-lib/convert"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToMaybeValueShouldConvertCorrectlyWhenDoingCorrectOperation(t *testing.T) {
	var stringMaybe monad.Maybe[string] = monad.Some[string]{Value: "10"}
	maybe := convert.ToMaybe[string, int](stringMaybe).Then(mustIntConverter)
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

func TestToMaybeValueShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringMaybe monad.Maybe[string] = monad.Some[string]{Value: "1gg0"}
	maybe := convert.ToMaybe[string, int](stringMaybe).Then(mustIntConverter)
	require.IsTypef(
		t,
		monad.Nothing[int]{},
		maybe,
		"maybe is expected to Nothing[int], got %T instead",
		maybe,
	)
}

func TestToMaybeNothingShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringMaybe monad.Maybe[string] = monad.Nothing[string]{}
	maybe := convert.ToMaybe[string, int](stringMaybe).Then(mustIntConverter)
	require.IsTypef(
		t,
		monad.Nothing[int]{},
		maybe,
		"maybe is expected to Nothing[int], got %T instead",
		maybe,
	)
}

func TestToMaybeOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intMaybe monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.Equal(
		t,
		"Some[int, string]{ToValue: 10}",
		fmt.Sprint(convert.ToMaybe[int, string](intMaybe)),
	)
	var stringMaybe monad.Maybe[string] = monad.Some[string]{Value: "10"}
	assert.Equal(
		t,
		"Some[string, int]{ToValue: \"10\"}",
		fmt.Sprint(convert.ToMaybe[string, int](stringMaybe)),
	)
	var errorResult monad.Maybe[int] = monad.Nothing[int]{}
	assert.Equal(
		t,
		"Nothing[int, string]{}",
		fmt.Sprint(convert.ToMaybe[int, string](errorResult)),
	)
}

func mustIntConverter(number string) monad.Maybe[int] {
	v, e := strconv.Atoi(number)
	if e != nil {
		return monad.Nothing[int]{}
	}
	return monad.Some[int]{Value: v}
}
