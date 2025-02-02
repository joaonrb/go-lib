package monad_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResultErrorStringRepresentationShouldHaveTheType(t *testing.T) {
	var result monad.Result[int] = monad.Error[int]{Err: errors.New("foo")}
	assert.Equal(t, "Error[int]{Err: \"foo\"}", fmt.Sprint(result))
}

func TestResultErrorShouldNotExecuteThenMethod(t *testing.T) {
	var test monad.Result[int] = monad.Error[int]{Err: errors.New("foo")}
	result := test.Then(func(i int) monad.Result[int] {
		return monad.OK[int]{Value: 1000}
	})
	require.IsType(t, monad.Error[int]{}, result)
	assert.Equal(t, "foo", result.(monad.Error[int]).Err.Error())
}

func TestResultErrorShouldNotExecuteWhenOKMethod(t *testing.T) {
	var test monad.Result[int] = monad.Error[int]{Err: errors.New("foo")}
	test.WhenOK(func(i int) {
		assert.Fail(t, "Error.WhenOK should not execute")
	})
}

func TestResultErrorShouldExecuteErrorMethod(t *testing.T) {
	var test monad.Result[int] = monad.Error[int]{Err: errors.New("foo")}
	result := test.Error(func(err error) monad.Result[int] {
		return monad.Error[int]{Err: errors.New("new foo")}
	})
	require.IsType(t, monad.Error[int]{}, result)
	assert.Equal(t, "new foo", result.(monad.Error[int]).Err.Error())
}

func TestResultErrorShouldExecuteWhenErrorMethod(t *testing.T) {
	var test monad.Result[int] = monad.Error[int]{Err: errors.New("foo")}
	var result monad.Result[int]
	test.WhenError(func(err error) {
		result = monad.Error[int]{Err: errors.New("new foo")}
	})
	require.IsType(t, monad.Error[int]{}, result)
	assert.Equal(t, "new foo", result.(monad.Error[int]).Err.Error())
}

func TestResultErrorTryValueShouldPanicTheError(t *testing.T) {
	var test monad.Result[int] = monad.Error[int]{Err: errors.New("foo")}
	assert.Panics(t, func() {
		_ = test.TryValue()
	})
}

func TestResultErrorTryErrorShouldReturnTheError(t *testing.T) {
	err := errors.New("foo")
	var test monad.Result[int] = monad.Error[int]{Err: err}
	assert.ErrorAs(t, test.TryError(), &err)
}

func TestResultErrorIsShouldReturnFalse(t *testing.T) {
	err := errors.New("foo")
	var test monad.Result[int] = monad.Error[int]{Err: err}
	assert.False(t, test.Is(10))
}

func TestResultErrorIsErrorShouldReturnTrueWhenUseTheSameError(t *testing.T) {
	err := errors.New("foo")
	var test monad.Result[int] = monad.Error[int]{Err: err}
	assert.True(t, test.IsError(err))
}

func TestResultErrorIsErrorShouldReturnFalseWhenUseDifferentError(t *testing.T) {
	err := errors.New("foo")
	var test monad.Result[int] = monad.Error[int]{Err: err}
	assert.False(t, test.IsError(errors.New("bar")))
}
