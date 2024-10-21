package types_test

import (
	"fmt"
	"testing"

	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionNothingStringRepresentationShouldHaveTheType(t *testing.T) {
	var intOption types.Option[int] = types.Nothing[int]{}
	assert.Equal(t, "Nothing[int]{}", fmt.Sprint(intOption))
	var stringOption types.Option[string] = types.Nothing[string]{}
	assert.Equal(t, "Nothing[string]{}", fmt.Sprint(stringOption))
}

func TestOptionNothingShouldNotExecuteThenMethod(t *testing.T) {
	var test types.Option[int] = types.Nothing[int]{}
	result := test.Then(func(i int) types.Option[int] {
		return types.Value[int]{This: 1000}
	})
	require.IsType(t, types.Nothing[int]{}, result)
}

func TestOptionNothingShouldNotExecuteWhenValueMethod(t *testing.T) {
	var test types.Option[int] = types.Nothing[int]{}
	test.WhenValue(func(i int) {
		assert.Fail(t, "Nothing.WhenValue should not execute")
	})
}

func TestOptionNothingShouldExecuteElseMethod(t *testing.T) {
	var test types.Option[int] = types.Nothing[int]{}
	result := test.Else(func() types.Option[int] {
		return types.Value[int]{This: 10}
	})
	require.IsType(t, types.Value[int]{}, result)
	assert.Equal(t, 10, result.(types.Value[int]).This)
}

func TestOptionNothingShouldExecuteWhenNothingMethod(t *testing.T) {
	var test types.Option[int] = types.Nothing[int]{}
	var result types.Option[int]
	test.WhenNothing(func() {
		result = types.Value[int]{This: 10}
	})
	require.IsType(t, types.Value[int]{}, result)
	assert.Equal(t, 10, result.(types.Value[int]).This)
}

func TestOptionNothingMustValueShouldExecuteWhenNothingMethod(t *testing.T) {
	var test types.Option[int] = types.Nothing[int]{}
	assert.Panics(t, func() {
		_ = test.MustValue()
	})
}
