package monad

import (
	"fmt"
)

var _ Maybe[any] = Nothing[any]{}

type Nothing[T any] struct{}

func (Nothing[T]) maybe() {}

func (nothing Nothing[T]) Then(call func(T) Maybe[T]) Maybe[T] {
	return nothing
}

func (nothing Nothing[T]) Else(call func() Maybe[T]) Maybe[T] {
	return call()
}

func (nothing Nothing[T]) WhenValue(call func(T)) Maybe[T] {
	return nothing
}

func (nothing Nothing[T]) WhenNothing(call func()) Maybe[T] {
	call()
	return nothing
}

func (nothing Nothing[T]) Or(value T) Maybe[T] {
	return Some[T]{Value: value}
}

func (nothing Nothing[T]) Is(value T) bool {
	return false
}

func (Nothing[T]) TryValue() T {
	panic(NewMaybeIsNothingError())
}

func (Nothing[T]) String() string {
	var value T
	return fmt.Sprintf("Nothing[%T]{}", value)
}
