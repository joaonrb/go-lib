package errors

import "github.com/go-errors/errors"

type Error interface {
	Error() string
	Stack() []byte
	ErrorStack() string
}

func New(text string) Error {
	return errors.Wrap(text, 1)
}
