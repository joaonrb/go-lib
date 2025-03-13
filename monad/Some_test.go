package monad_test

import (
	"fmt"
	"testing"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaybeValueStringRepresentationShouldHaveTheType(t *testing.T) {
	var intMaybe monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.Equal(t, "Some[int]{Value: 10}", fmt.Sprint(intMaybe))
	var stringMaybe monad.Maybe[string] = monad.Some[string]{Value: "João Nuno"}
	assert.Equal(t, "Some[string]{Value: \"João Nuno\"}", fmt.Sprint(stringMaybe))
}

func TestMaybeValueShouldExecuteThenMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	result := test.Then(func(i int) monad.Maybe[int] {
		return monad.Some[int]{Value: 1000}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 1000, result.(monad.Some[int]).Value)
}

func TestMaybeValueShouldExecuteWhenValueMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	var result monad.Maybe[int]
	test.WhenValue(func(i int) {
		result = monad.Some[int]{Value: 1000}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 1000, result.(monad.Some[int]).Value)
}

func TestMaybeValueShouldNotExecuteElseMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	result := test.Else(func() monad.Maybe[int] {
		return monad.Nothing[int]{}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 10, result.(monad.Some[int]).Value)
}

func TestMaybeValueShouldNotExecuteWhenNothingMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	test.WhenNothing(func() {
		assert.Fail(t, "Some.WhenNothing should not execute")
	})
}

func TestMaybeValueTryValueShouldReturnTheCorrectValue(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.NotPanics(t, func() {
		assert.Equal(t, 10, test.TryValue(), "Some.TryValue should return 10")
	})
}
