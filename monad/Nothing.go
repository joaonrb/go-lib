package monad

import (
	"fmt"
	"github.com/joaonrb/go-lib/op"
)

var _ Maybe[any] = Nothing[any]{}

type Nothing[T any] struct{}

func (Nothing[T]) maybe() {}

func (nothing Nothing[T]) Then(func(T) Maybe[T]) Maybe[T] {
	return nothing
}

func (nothing Nothing[T]) Else(call func() Maybe[T]) Maybe[T] {
	return call()
}

func (nothing Nothing[T]) WhenValue(func(T)) Maybe[T] {
	return nothing
}

func (nothing Nothing[T]) WhenNothing(call func()) Maybe[T] {
	call()
	return nothing
}

func (nothing Nothing[T]) Or(value T) Maybe[T] {
	return Some[T]{Value: value}
}

func (nothing Nothing[T]) If(op.Operator[T]) bool {
	return false
}

func (nothing Nothing[T]) DoIf(op.Operator[T], func(T) Maybe[T]) Maybe[T] {
	return nothing
}

func (Nothing[T]) TryValue() T {
	panic(NewMaybeIsNothingError())
}

func (Nothing[T]) String() string {
	var value T
	return fmt.Sprintf("Nothing[%T]{}", value)
}
