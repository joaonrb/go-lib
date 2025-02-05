package async

import "github.com/joaonrb/go-lib/errors"

type inner = errors.Error

type FutureNotStartedError struct {
	inner
}

func NewFutureNotStartedError() FutureNotStartedError {
	return FutureNotStartedError{
		inner: errors.New("Future is not started yet"),
	}
}

type FutureNotFinishedError struct {
	inner
}

func NewFutureNotFinishedError() FutureNotFinishedError {
	return FutureNotFinishedError{
		inner: errors.New("Future is not finished yet"),
	}
}
