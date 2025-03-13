package monad

import (
	"fmt"
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

func (ok OK[T]) Is(value T) bool {
	var s, t any = ok.Value, value
	return s == t
}

func (ok OK[T]) IsIn(values ...T) bool {
	var (
		s any = ok.Value
		t any
	)
	result := false
	for i := 0; !result && i < len(values); i++ {
		t = values[i]
		result = result || s == t
	}
	return result
}

func (ok OK[T]) IsError(error) bool {
	return false
}

func (ok OK[T]) IsErrorIn(...error) bool {
	return false
}

func (ok OK[T]) AsError(any) bool {
	return false
}

func (ok OK[T]) AsErrorIn(...any) bool {
	return false
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
