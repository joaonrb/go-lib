package types

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

func (ok OK[T]) String() string {
	var t any = ok.Value
	switch t.(type) {
	case string:
		return fmt.Sprintf("OK[%T]{Value: \"%v\"}", ok.Value, ok.Value)
	default:
		return fmt.Sprintf("OK[%T]{Value: %v}", ok.Value, ok.Value)
	}
}
