package collections

import "github.com/joaonrb/go-lib/errors"

type inner = errors.Error

func CollectionClosed() *CollectionClosedError {
	return &CollectionClosedError{
		inner: errors.New("CollectionClosedError: collection is closed"),
	}
}

type CollectionClosedError struct {
	inner
}
