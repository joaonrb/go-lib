package monad_test

import (
	"errors"
	"fmt"
	"github.com/joaonrb/go-lib/op"
	"testing"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResultOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intResult monad.Result[int] = monad.OK[int]{Value: 10}
	assert.Equal(t, "OK[int]{Some: 10}", fmt.Sprint(intResult))
	var stringResult monad.Result[string] = monad.OK[string]{Value: "João Nuno"}
	assert.Equal(t, "OK[string]{Some: \"João Nuno\"}", fmt.Sprint(stringResult))
}

func TestResultOKShouldExecuteThenMethod(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	result := test.Then(func(i int) monad.Result[int] {
		return monad.OK[int]{Value: 1000}
	})
	require.IsType(t, monad.OK[int]{}, result)
	assert.Equal(t, 1000, result.(monad.OK[int]).Value)
}

func TestResultOKShouldExecuteWhenOKMethod(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	var result monad.Result[int]
	test.WhenOK(func(i int) {
		result = monad.OK[int]{Value: 1000}
	})
	require.IsType(t, monad.OK[int]{}, result)
	assert.Equal(t, 1000, result.(monad.OK[int]).Value)
}

func TestResultOKShouldNotExecuteErrorMethod(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	result := test.Error(func(err error) monad.Result[int] {
		return monad.Error[int]{Err: errors.New("new foo")}
	})
	require.IsType(t, monad.OK[int]{}, result)
	assert.Equal(t, 10, result.(monad.OK[int]).Value)
}

func TestResultOKShouldNotExecuteWhenErrorMethod(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	test.WhenError(func(err error) {
		assert.Fail(t, "OK.WhenError should not execute")
	})
}

func TestResultOKTryValueShouldReturnTheRightValue(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.NotPanics(t, func() {
		assert.Equal(t, 10, test.TryValue(), "OK.TryValue should return the correct value")
	})
}

func TestResultOKTryErrorShouldRaiseAnError(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.PanicsWithError(t, "result is ok", func() {
		assert.Equal(t, 10, test.TryError(), "OK.TryError should panic")
	})
}

func TestResultOKIfEqualShouldReturnTrueWhenUseTheSameValue(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.True(t, test.If(op.Equal(10)))
}

func TestResultOKIfEqualShouldReturnFalseWhenUseTheDifferentValue(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.False(t, test.If(op.Equal(11)))
}

func TestResultOKIfInShouldReturnTrueWhenHaveAnyEqualValue(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.True(t, test.If(op.In(1, 2, 3, 4, 5, 10)))
}

func TestResultOKIfInShouldReturnFalseWhenDoesNotHaveAnyEqualValue(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.False(t, test.If(op.In(1, 2, 3, 4, 5)))
}

func TestResultOKIfErrorEqualShouldReturnFalse(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.False(t, test.IfError(op.Equal(errors.New("foo"))))
}

func TestResultOKIsErrorInShouldReturnFalse(t *testing.T) {
	var test monad.Result[int] = monad.OK[int]{Value: 10}
	assert.False(t, test.IfError(op.In(errors.New("foo"), errors.New("bar"))))
}
