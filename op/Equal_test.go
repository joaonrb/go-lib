package op_test

import (
	"github.com/joaonrb/go-lib/op"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertTrue[T any](t *testing.T, operator op.Operator[T], value T) {
	assert.True(t, operator.Evaluate(value))
}

func TestEqualShouldReturnTrueWhenUseTheSameValue(t *testing.T) {
	assertTrue(t, op.Equal(10), 10)
}
