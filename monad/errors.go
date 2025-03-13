package monad

import "github.com/joaonrb/go-lib/errors"

type inner = errors.Error

type MaybeIsNothingError struct {
	inner
}

func NewMaybeIsNothingError() MaybeIsNothingError {
	return MaybeIsNothingError{inner: errors.New("maybe is nothing")}
}

type ResultIsOkError struct {
	inner
}

func NewResultIsOkError() ResultIsOkError {
	return ResultIsOkError{inner: errors.New("result is ok")}
}

type ResultIsErrorError struct {
	inner
	Err error
}

func NewResultIsErrorError(err error) ResultIsErrorError {
	return ResultIsErrorError{
		inner: errors.Newf("result is error: %s", err),
		Err:   err,
	}
}
