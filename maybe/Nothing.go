package maybe

import (
	"fmt"
	"reflect"
)

type Nothing[T any] struct{}

func (Nothing[T]) maybe() {}

func (nothing Nothing[T]) Then(func(T) Maybe[T]) Maybe[T] {
	return nothing
}

func (nothing Nothing[T]) IfNothing(call func()) {
	call()
}

func (Nothing[T]) String() string {
	var value T
	return fmt.Sprintf("Just[%s]{Value: %v}", reflect.TypeOf(value).Name(), value)
}
