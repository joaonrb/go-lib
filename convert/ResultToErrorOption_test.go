package convert_test

import (
	"errors"
	"github.com/joaonrb/go-lib/convert"
	"testing"

	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResultToErrorOptionShouldReturnNothingWhenResultIsOK(t *testing.T) {
	result := convert.ResultToErrorOption(types.OK[int]{Value: 10})
	require.IsType(t, types.Nothing[error]{}, result)
}

func TestResultToErrorOptionShouldReturnValueWhenResultIsError(t *testing.T) {
	result := convert.ResultToErrorOption(types.Error[int]{Err: errors.New("an error")})
	require.IsType(t, types.Value[error]{}, result)
	assert.Equal(t, "an error", result.(types.Value[error]).This.Error())
}
