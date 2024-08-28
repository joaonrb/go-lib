package types_test

import (
	"fmt"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionValueStringRepresentationShouldHaveTheType(t *testing.T) {
	var intOption types.Option[int] = types.Value[int]{This: 10}
	assert.Equal(t, "Value[int]{This: 10}", fmt.Sprint(intOption))
	var stringOption types.Option[string] = types.Value[string]{This: "João Nuno"}
	assert.Equal(t, "Value[string]{This: \"João Nuno\"}", fmt.Sprint(stringOption))
}
