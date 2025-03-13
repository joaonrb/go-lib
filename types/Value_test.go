package types_test

import (
	"fmt"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOptionValueStringRepresentationShouldHaveTheType(t *testing.T) {
	var intOption types.Option[int] = types.Value[int]{This: 10}
	assert.Equal(t, "Value[int]{This: 10}", fmt.Sprint(intOption))
	var stringOption types.Option[string] = types.Value[string]{This: "João Nuno"}
	assert.Equal(t, "Value[string]{This: \"João Nuno\"}", fmt.Sprint(stringOption))
}

func TestOptionValueShouldExecuteThenMethod(t *testing.T) {
	var test types.Option[int] = types.Value[int]{This: 10}
	result := test.Then(func(i int) types.Option[int] {
		return types.Value[int]{This: 1000}
	})
	require.IsType(t, types.Value[int]{}, result)
	assert.Equal(t, 1000, result.(types.Value[int]).This)
}

func TestOptionValueShouldExecuteWhenValueMethod(t *testing.T) {
	var test types.Option[int] = types.Value[int]{This: 10}
	var result types.Option[int]
	test.WhenValue(func(i int) {
		result = types.Value[int]{This: 1000}
	})
	require.IsType(t, types.Value[int]{}, result)
	assert.Equal(t, 1000, result.(types.Value[int]).This)
}

func TestOptionValueShouldNotExecuteElseMethod(t *testing.T) {
	var test types.Option[int] = types.Value[int]{This: 10}
	result := test.Else(func() types.Option[int] {
		return types.Nothing[int]{}
	})
	require.IsType(t, types.Value[int]{}, result)
	assert.Equal(t, 10, result.(types.Value[int]).This)
}

func TestOptionValueShouldNotExecuteWhenNothingMethod(t *testing.T) {
	var test types.Option[int] = types.Value[int]{This: 10}
	test.WhenNothing(func() {
		assert.Fail(t, "Value.WhenNothing should not execute")
	})
}
