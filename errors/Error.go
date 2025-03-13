package errors

import (
	"fmt"

	"github.com/go-errors/errors"
)

type Error interface {
	Error() string
	Stack() []byte
	ErrorStack() string
}

func New(text string) Error {
	return errors.Wrap(text, 1)
}

func Newf(tpl string, args ...any) Error {
	return errors.Wrap(fmt.Sprintf(tpl, args...), 1)
}
