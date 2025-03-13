package errors_test

import (
	"testing"

	"github.com/joaonrb/go-lib/errors"
	"github.com/stretchr/testify/assert"
)

func TestPanicShouldPanicWhenError(t *testing.T) {
	err := errors.New("foo")
	assert.PanicsWithError(t, err.Error(), func() {
		errors.Panic(err)
	})
}

func TestPanicShouldNotPanicWhenNil(t *testing.T) {
	assert.NotPanics(t, func() {
		errors.Panic(nil)
	})
}
