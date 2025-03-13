package monad

import (
	"fmt"
	"github.com/joaonrb/go-lib/op"
)

var _ Result[any] = OK[any]{}

type OK[T any] struct {
	Value T
}

func (ok OK[T]) result() {}

func (ok OK[T]) Then(call func(T) Result[T]) Result[T] {
	return call(ok.Value)
}

func (ok OK[T]) Error(func(error) Result[T]) Result[T] {
	return ok
}

func (ok OK[T]) WhenOK(call func(T)) Result[T] {
	call(ok.Value)
	return ok
}

func (ok OK[T]) WhenError(func(error)) Result[T] {
	return ok
}

func (ok OK[T]) Or(T) Result[T] {
	return ok
}

func (ok OK[T]) If(operator op.Operator[T]) bool {
	return operator(ok.Value)
}

func (ok OK[T]) DoIf(comparator op.Operator[T], do func(T) Result[T]) Result[T] {
	if ok.If(comparator) {
		return do(ok.Value)
	} else {
		return ok
	}
}

func (ok OK[T]) IfError(op.Operator[error]) bool {
	return false
}

func (ok OK[T]) DoIfError(op.Operator[error], func(error) Result[T]) Result[T] {
	return ok
}

func (ok OK[T]) TryValue() T {
	return ok.Value
}

func (ok OK[T]) TryError() error {
	panic(NewResultIsOkError())
}

func (ok OK[T]) String() string {
	var value any = ok.Value
	switch value := value.(type) {
	case string, fmt.Stringer:
		return fmt.Sprintf("OK[%T]{Some: \"%s\"}", value, value)
	default:
		return fmt.Sprintf("OK[%T]{Some: %v}", value, value)
	}
}
