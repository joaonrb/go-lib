package types

import "github.com/joaonrb/go-lib/errors"

type inner = errors.Error

type NothingError struct {
	inner
}

func NewNothingError() NothingError {
	return NothingError{inner: errors.New("result is nothing")}
}
