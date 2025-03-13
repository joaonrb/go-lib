package errors_test

import (
	goerrs "errors"
	"github.com/joaonrb/go-lib/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicShouldPanicWhenError(t *testing.T) {
	err := goerrs.New("foo")
	assert.PanicsWithError(t, err.Error(), func() {
		errors.Panic(err)
	})
}

func TestPanicShouldNotPanicWhenNil(t *testing.T) {
	assert.NotPanics(t, func() {
		errors.Panic(nil)
	})
}
