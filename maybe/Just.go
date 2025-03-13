package maybe

import (
	"fmt"
	"reflect"
)

type Just[T any] struct {
	Value T
}

func (just Just[T]) maybe() {}

func (just Just[T]) Then(call func(T) Maybe[T]) Maybe[T] {
	return call(just.Value)
}

func (just Just[T]) IfNothing(func()) {}

func (just Just[T]) String() string {
	return fmt.Sprintf("Just[%s]{Value: %v}", reflect.TypeOf(just.Value).Name(), just.Value)
}
