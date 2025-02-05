package atomic

import "github.com/joaonrb/go-lib/errors"

type inner = errors.Error

type ValueNotInitializedError struct {
	inner
}

func NewValueNotInitializedError() ValueNotInitializedError {
	return ValueNotInitializedError{
		inner: errors.New("Value is not initialized yet"),
	}
}
