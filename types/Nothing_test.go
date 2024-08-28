package types_test

import (
	"fmt"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionNothingStringRepresentationShouldHaveTheType(t *testing.T) {
	var intOption types.Option[int] = types.Nothing[int]{}
	assert.Equal(t, "Nothing[int]{}", fmt.Sprint(intOption))
	var stringOption types.Option[string] = types.Nothing[string]{}
	assert.Equal(t, "Nothing[string]{}", fmt.Sprint(stringOption))
}
