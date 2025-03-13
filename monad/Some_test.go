package monad_test

import (
	"fmt"
	"github.com/joaonrb/go-lib/op"
	"testing"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaybeSomeStringRepresentationShouldHaveTheType(t *testing.T) {
	var intMaybe monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.Equal(t, "Some[int]{Value: 10}", fmt.Sprint(intMaybe))
	var stringMaybe monad.Maybe[string] = monad.Some[string]{Value: "João Nuno"}
	assert.Equal(t, "Some[string]{Value: \"João Nuno\"}", fmt.Sprint(stringMaybe))
}

func TestMaybeSomeThenShouldExecuteThenMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	result := test.Then(func(i int) monad.Maybe[int] {
		return monad.Some[int]{Value: 1000}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 1000, result.(monad.Some[int]).Value)
}

func TestMaybeSomeShouldExecuteWhenValueMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	var result monad.Maybe[int]
	test.WhenValue(func(i int) {
		result = monad.Some[int]{Value: 1000}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 1000, result.(monad.Some[int]).Value)
}

func TestMaybeSomeShouldNotExecuteElseMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	result := test.Else(func() monad.Maybe[int] {
		return monad.Nothing[int]{}
	})
	require.IsType(t, monad.Some[int]{}, result)
	assert.Equal(t, 10, result.(monad.Some[int]).Value)
}

func TestMaybeSomeShouldNotExecuteWhenNothingMethod(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	test.WhenNothing(func() {
		assert.Fail(t, "Some.WhenNothing should not execute")
	})
}

func TestMaybeSomeTryValueShouldReturnTheCorrectValue(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.NotPanics(t, func() {
		assert.Equal(t, 10, test.TryValue(), "Some.TryValue should return 10")
	})
}

func TestMaybeSomeIfEqualShouldReturnTrueWhenUseTheSameValue(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.True(t, test.If(op.Equal(10)))
}

func TestMaybeSomeIfEqualShouldReturnFalseWhenUseTheDifferentValue(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.False(t, test.If(op.Equal(11)))
}

func TestMaybeSomeIfInShouldReturnTrueWhenHaveAtLeastOneEqualValue(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.True(t, test.If(op.In(1, 2, 3, 4, 5, 10)))
}

func TestMaybeSomeIfInShouldReturnFalseWhenDontHaveAnyEqualValue(t *testing.T) {
	var test monad.Maybe[int] = monad.Some[int]{Value: 10}
	assert.False(t, test.If(op.In(1, 2, 3, 4, 5)))
}
