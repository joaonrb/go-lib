package types

import "github.com/joaonrb/go-lib/errors"

type inner = errors.Error

type OptionIsNothingError struct {
	inner
}

func NewOptionsIsNothingError() OptionIsNothingError {
	return OptionIsNothingError{inner: errors.New("option is nothing")}
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
