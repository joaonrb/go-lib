package convert_test

import (
	"errors"
	"fmt"
	"github.com/joaonrb/go-lib/convert"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestResultTypeOKShouldConvertCorrectlyWhenDoingCorrectOperation(t *testing.T) {
	var stringResult types.Result[string] = types.OK[string]{Value: "10"}
	result := convert.ResultType[string, int](stringResult).Then(intConverter)
	require.IsTypef(t, types.OK[int]{}, result, "result is expected to OK[int], got %T instead", result)
	value := result.(types.OK[int]).Value
	assert.Equal(t, 10, value, "result value is expected to be 10, got %d instead", value)
}

func TestResultTypeOKShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringResult types.Result[string] = types.OK[string]{Value: "1gg0"}
	result := convert.ResultType[string, int](stringResult).Then(intConverter)
	require.IsTypef(t, types.Error[int]{}, result, "result is expected to Error[int], got %T instead", result)
	err := result.(types.Error[int]).Err
	assert.Equal(t, "NotNumberError: string \"1gg0\" is cannot be converted to a number", err.Error())
}

func TestResultTypeErrorShouldNotConvertWhenDoingIncorrectOperation(t *testing.T) {
	var stringResult types.Result[string] = types.Error[string]{Err: errors.New("some error")}
	result := convert.ResultType[string, int](stringResult).Then(intConverter)
	require.IsTypef(t, types.Error[int]{}, result, "result is expected to Error[int], got %T instead", result)
	err := result.(types.Error[int]).Err
	assert.Equal(t, "some error", err.Error())
}

func TestResultTypeOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intResult types.Result[int] = types.OK[int]{Value: 10}
	assert.Equal(t, "OK[int, string]{Value: 10}", fmt.Sprint(convert.ResultType[int, string](intResult)))
	var stringResult types.Result[string] = types.OK[string]{Value: "10"}
	assert.Equal(t, "OK[string, int]{Value: \"10\"}", fmt.Sprint(convert.ResultType[string, int](stringResult)))
	var errorResult types.Result[int] = types.Error[int]{Err: errors.New("some error")}
	assert.Equal(t, "Error[int, string]{Err: \"some error\"}", fmt.Sprint(convert.ResultType[int, string](errorResult)))
}

func intConverter(number string) types.Result[int] {
	v, e := strconv.Atoi(number)
	if e != nil {
		return types.Error[int]{
			Err: fmt.Errorf("NotNumberError: string \"%s\" is cannot be converted to a number", number),
		}
	}
	return types.OK[int]{Value: v}
}
