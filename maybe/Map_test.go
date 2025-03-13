package maybe_test

import (
	"github.com/joaonrb/go-lib/maybe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMapWhenMaybeIsJustAndCallReturnsJustThenResultShouldBeJust(t *testing.T) {
	t.Parallel()
	var value maybe.Maybe[string] = maybe.Just[string]{Value: "10"}
	r := maybe.Map(value, stringToInt)
	require.IsType(t, maybe.Just[int64]{}, r)
	assert.IsType(t, int64(10), r.(maybe.Just[int64]).Value)
}

func TestMapWhenMaybeIsNothingThenMaybeShouldBeNothing(t *testing.T) {
	t.Parallel()
	var value maybe.Maybe[string] = maybe.Nothing[string]{}
	r := maybe.Map(value, stringToInt)
	require.IsType(t, maybe.Nothing[int64]{}, r)
}

func TestMapResultIsOkAndCallReturnsErrorThenResultShouldBeError(t *testing.T) {
	t.Parallel()
	var value maybe.Maybe[string] = maybe.Just[string]{Value: "José"}
	r := maybe.Map(value, stringToInt)
	require.IsType(t, maybe.Nothing[int64]{}, r)
}

func stringToInt(value string) maybe.Maybe[int64] {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return maybe.Nothing[int64]{}
	}
	return maybe.Just[int64]{Value: intValue}
}
