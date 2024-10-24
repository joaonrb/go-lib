package types

import (
	"fmt"
)

var _ Option[any] = Nothing[any]{}

type Nothing[T any] struct{}

func (Nothing[T]) option() {}

func (nothing Nothing[T]) Then(call func(T) Option[T]) Option[T] {
	return nothing
}

func (nothing Nothing[T]) Else(call func() Option[T]) Option[T] {
	return call()
}

func (nothing Nothing[T]) WhenValue(call func(T)) Option[T] {
	return nothing
}

func (nothing Nothing[T]) WhenNothing(call func()) Option[T] {
	call()
	return nothing
}

func (Nothing[T]) MustValue() T {
	panic(NewNothingError())
}

func (Nothing[T]) String() string {
	var value T
	return fmt.Sprintf("Nothing[%T]{}", value)
}
