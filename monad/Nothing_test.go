package monad_test

import (
	"fmt"
	"testing"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaybeNothingStringRepresentationShouldHaveTheType(t *testing.T) {
	var intMaybe monad.Maybe[int] = monad.Nothing[int]{}
	assert.Equal(t, "Nothing[int]{}", fmt.Sprint(intMaybe))
	var stringMaybe monad.Maybe[string] = monad.Nothing[string]{}
	assert.Equal(t, "Nothing[string]{}", fmt.Sprint(stringMaybe))
}

func TestMaybeNothingShouldNotExecuteThenMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	result := test.Then(func(i int) monad.Maybe[int] {
		return monad.Some[int]{Value: 1000}
	})
	require.IsType(t, monad.Nothing[int]{}, result)
}

func TestMaybeNothingShouldNotExecuteWhenValueMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	test.WhenValue(func(i int) {
		assert.Fail(t, "Nothing.WhenValue should not execute")
	})
}

func TestMaybeNothingShouldExecuteElseMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	result := test.Else(func() monad.Maybe[int] {
		return monad.Some[int]{Value: 10}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 10, result.(monad.Some[int]).Value)
}

func TestMaybeNothingShouldExecuteWhenNothingMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	var result monad.Maybe[int]
	test.WhenNothing(func() {
		result = monad.Some[int]{Value: 10}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 10, result.(monad.Some[int]).Value)
}

func TestMaybeNothingTryValueShouldExecuteWhenNothingMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	assert.Panics(t, func() {
		_ = test.TryValue()
	})
}

func TestMaybeNothingIsShouldReturnFalse(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	assert.False(t, test.Is(10))
}

func TestMaybeNothingIsNotShouldReturnTrue(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	assert.True(t, test.IsNot(10))
}

func TestMaybeNothingIsInShouldReturnFalse(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	assert.False(t, test.IsIn(1, 2, 3, 4, 5, 10))
}

func TestMaybeValueIsInShouldReturnTrue(t *testing.T) {
	var test monad.Maybe[int] = monad.Nothing[int]{}
	assert.True(t, test.IsNotIn(1, 2, 3, 4, 5))
}
