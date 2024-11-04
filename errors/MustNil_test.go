package errors_test

import (
	"testing"

	"github.com/joaonrb/go-lib/errors"
	"github.com/stretchr/testify/assert"
)

func TestMustNilShouldPanicWhenError(t *testing.T) {
	err := errors.New("foo")
	assert.PanicsWithError(t, err.Error(), func() {
		errors.MustNil(err)
	})
}

func TestMustNilShouldNotPanicWhenNil(t *testing.T) {
	assert.NotPanics(t, func() {
		errors.MustNil(nil)
	})
}
