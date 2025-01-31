package queue

import "github.com/joaonrb/go-lib/errors"

type inner = errors.Error

func NewQueueClosedError() *QueueClosedError {
	return &QueueClosedError{
		inner: errors.New("QueueClosedError: collection is closed"),
	}
}

type QueueClosedError struct {
	inner
}
