package types_test

import (
	"errors"
	"fmt"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResultErrorStringRepresentationShouldHaveTheType(t *testing.T) {
	var result types.Result[int] = types.Error[int]{Err: errors.New("foo")}
	assert.Equal(t, "Error[int]{Err: \"foo\"}", fmt.Sprint(result))
}
