package convert_test

import (
	"errors"
	"testing"

	"github.com/joaonrb/go-lib/convert"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResultToMaybeErrorShouldReturnNothingWhenResultIsOK(t *testing.T) {
	result := convert.ResultToMaybeError(monad.OK[int]{Value: 10})
	require.IsType(t, monad.Nothing[error]{}, result)
}

func TestResultToMaybeErrorShouldReturnValueWhenResultIsError(t *testing.T) {
	result := convert.ResultToMaybeError(monad.Error[int]{Err: errors.New("an error")})
	require.IsType(t, monad.Some[error]{}, result)
	assert.Equal(t, "an error", result.(monad.Some[error]).Value.Error())
}
