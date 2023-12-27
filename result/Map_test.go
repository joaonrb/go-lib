package result_test

import (
	"errors"
	"github.com/joaonrb/go-lib/result"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMapWhenResultIsOkAndCallReturnsOkThenResultShouldBeOk(t *testing.T) {
	t.Parallel()
	var value result.Result[string, error] = result.OK[string, error]{Value: "10"}
	r := result.Map(value, stringToInt)
	require.IsType(t, result.OK[int64, error]{}, r)
	assert.IsType(t, int64(10), r.(result.OK[int64, error]).Value)
}

func TestMapWhenResultIsErrorThenResultShouldBeError(t *testing.T) {
	t.Parallel()
	err := errors.New("unsupported operation")
	var value result.Result[string, error] = result.Error[string, error]{Err: err}
	r := result.Map(value, stringToInt)
	require.IsType(t, result.Error[int64, error]{}, r)
	assert.IsType(t, err, r.(result.Error[int64, error]).Err)
}

func TestMapResultIsOkAndCallReturnsErrorThenResultShouldBeError(t *testing.T) {
	t.Parallel()
	var value result.Result[string, error] = result.OK[string, error]{Value: "José"}
	r := result.Map(value, stringToInt)
	require.IsType(t, result.Error[int64, error]{}, r)
	assert.IsType(t, &strconv.NumError{}, r.(result.Error[int64, error]).Err)
}

func stringToInt(value string) result.Result[int64, error] {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return result.Error[int64, error]{Err: err}
	}
	return result.OK[int64, error]{Value: intValue}
}
