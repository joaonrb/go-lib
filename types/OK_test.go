package types_test

import (
	"errors"
	"fmt"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResultOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intResult types.Result[int] = types.OK[int]{Value: 10}
	assert.Equal(t, "OK[int]{Value: 10}", fmt.Sprint(intResult))
	var stringResult types.Result[string] = types.OK[string]{Value: "João Nuno"}
	assert.Equal(t, "OK[string]{Value: \"João Nuno\"}", fmt.Sprint(stringResult))
}

func TestResultOKShouldExecuteThenMethod(t *testing.T) {
	var test types.Result[int] = types.OK[int]{Value: 10}
	result := test.Then(func(i int) types.Result[int] {
		return types.OK[int]{Value: 1000}
	})
	require.IsType(t, types.OK[int]{}, result)
	assert.Equal(t, 1000, result.(types.OK[int]).Value)
}

func TestResultOKShouldExecuteWhenOKMethod(t *testing.T) {
	var test types.Result[int] = types.OK[int]{Value: 10}
	var result types.Result[int]
	test.WhenOK(func(i int) {
		result = types.OK[int]{Value: 1000}
	})
	require.IsType(t, types.OK[int]{}, result)
	assert.Equal(t, 1000, result.(types.OK[int]).Value)
}

func TestResultOKShouldNotExecuteErrorMethod(t *testing.T) {
	var test types.Result[int] = types.OK[int]{Value: 10}
	result := test.Error(func(err error) types.Result[int] {
		return types.Error[int]{Err: errors.New("new foo")}
	})
	require.IsType(t, types.OK[int]{}, result)
	assert.Equal(t, 10, result.(types.OK[int]).Value)
}

func TestResultOKShouldNotExecuteWhenErrorMethod(t *testing.T) {
	var test types.Result[int] = types.OK[int]{Value: 10}
	test.WhenError(func(err error) {
		assert.Fail(t, "OK.WhenError should not execute")
	})
}
