package errors_test

import (
	"testing"

	"github.com/joaonrb/go-lib/errors"
	"github.com/stretchr/testify/assert"
)

func TestRaiseShouldPanicWhenError(t *testing.T) {
	err := errors.New("foo")
	assert.PanicsWithError(t, err.Error(), func() {
		errors.Raise(err)
	})
}

func TestRaiseShouldNotPanicWhenNil(t *testing.T) {
	assert.NotPanics(t, func() {
		errors.Raise(nil)
	})
}
