package convertto_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/joaonrb/go-lib/convertto"
	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResultOKShouldConvertCorrectlyWhenDoingCorrectOperation(t *testing.T) {
	var stringResult monad.Result[string] = monad.OK[string]{Value: "10"}
	result := convertto.Result[string, int](stringResult).Then(intConverter)
	require.IsTypef(
		t,
		monad.OK[int]{},
		result,
		"result is expected to OK[int], got %T instead",
		result,
	)
	value := result.(monad.OK[int]).Value
	assert.Equal(t, 10, value, "result value is expected to be 10, got %d instead", value)
}

func TestResultOKShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringResult monad.Result[string] = monad.OK[string]{Value: "1gg0"}
	result := convertto.Result[string, int](stringResult).Then(intConverter)
	require.IsTypef(
		t,
		monad.Error[int]{},
		result,
		"result is expected to Error[int], got %T instead",
		result,
	)
	err := result.(monad.Error[int]).Err
	assert.Equal(
		t,
		"NotNumberError: string \"1gg0\" is cannot be converted to a number",
		err.Error(),
	)
}

func TestResultErrorShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringResult monad.Result[string] = monad.Error[string]{Err: errors.New("some error")}
	result := convertto.Result[string, int](stringResult).Then(intConverter)
	require.IsTypef(
		t,
		monad.Error[int]{},
		result,
		"result is expected to Error[int], got %T instead",
		result,
	)
	err := result.(monad.Error[int]).Err
	assert.Equal(t, "some error", err.Error())
}

func TestResultOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intResult monad.Result[int] = monad.OK[int]{Value: 10}
	assert.Equal(
		t,
		"OK[int, string]{Some: 10}",
		fmt.Sprint(convertto.Result[int, string](intResult)),
	)
	var stringResult monad.Result[string] = monad.OK[string]{Value: "10"}
	assert.Equal(
		t,
		"OK[string, int]{Some: \"10\"}",
		fmt.Sprint(convertto.Result[string, int](stringResult)),
	)
	var errorResult monad.Result[int] = monad.Error[int]{Err: errors.New("some error")}
	assert.Equal(
		t,
		"Error[int, string]{Err: \"some error\"}",
		fmt.Sprint(convertto.Result[int, string](errorResult)),
	)
}

func intConverter(number string) monad.Result[int] {
	v, e := strconv.Atoi(number)
	if e != nil {
		return monad.Error[int]{
			Err: fmt.Errorf(
				"NotNumberError: string \"%s\" is cannot be converted to a number",
				number,
			),
		}
	}
	return monad.OK[int]{Value: v}
}
