package types_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResultErrorStringRepresentationShouldHaveTheType(t *testing.T) {
	var result types.Result[int] = types.Error[int]{Err: errors.New("foo")}
	assert.Equal(t, "Error[int]{Err: \"foo\"}", fmt.Sprint(result))
}

func TestResultErrorShouldNotExecuteThenMethod(t *testing.T) {
	var test types.Result[int] = types.Error[int]{Err: errors.New("foo")}
	result := test.Then(func(i int) types.Result[int] {
		return types.OK[int]{Value: 1000}
	})
	require.IsType(t, types.Error[int]{}, result)
	assert.Equal(t, "foo", result.(types.Error[int]).Err.Error())
}

func TestResultErrorShouldNotExecuteWhenOKMethod(t *testing.T) {
	var test types.Result[int] = types.Error[int]{Err: errors.New("foo")}
	test.WhenOK(func(i int) {
		assert.Fail(t, "Error.WhenOK should not execute")
	})
}

func TestResultErrorShouldExecuteErrorMethod(t *testing.T) {
	var test types.Result[int] = types.Error[int]{Err: errors.New("foo")}
	result := test.Error(func(err error) types.Result[int] {
		return types.Error[int]{Err: errors.New("new foo")}
	})
	require.IsType(t, types.Error[int]{}, result)
	assert.Equal(t, "new foo", result.(types.Error[int]).Err.Error())
}

func TestResultErrorShouldExecuteWhenErrorMethod(t *testing.T) {
	var test types.Result[int] = types.Error[int]{Err: errors.New("foo")}
	var result types.Result[int]
	test.WhenError(func(err error) {
		result = types.Error[int]{Err: errors.New("new foo")}
	})
	require.IsType(t, types.Error[int]{}, result)
	assert.Equal(t, "new foo", result.(types.Error[int]).Err.Error())
}

func TestResultErrorMustValueShouldPanicTheError(t *testing.T) {
	var test types.Result[int] = types.Error[int]{Err: errors.New("foo")}
	assert.Panics(t, func() {
		_ = test.MustValue()
	})
}
