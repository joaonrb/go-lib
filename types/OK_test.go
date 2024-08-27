package types_test

import (
	"fmt"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResultOKStringRepresentationShouldHaveTheType(t *testing.T) {
	var intResult types.Result[int] = types.OK[int]{Value: 10}
	assert.Equal(t, "OK[int]{Value: 10}", fmt.Sprint(intResult))
	var stringResult types.Result[string] = types.OK[string]{Value: "João Nuno"}
	assert.Equal(t, "OK[string]{Value: \"João Nuno\"}", fmt.Sprint(stringResult))
}
