package types

import (
	"fmt"
)

var _ Option[any] = Value[any]{}

type Value[T any] struct {
	This T
}

func (value Value[T]) option() {}

func (value Value[T]) Then(call func(T) Option[T]) Option[T] {
	return call(value.This)
}

func (value Value[T]) Else(func() Option[T]) Option[T] {
	return value
}

func (value Value[T]) WhenValue(call func(T)) Option[T] {
	call(value.This)
	return value
}

func (value Value[T]) WhenNothing(func()) Option[T] {
	return value
}

func (value Value[T]) String() string {
	var this any = value.This
	switch this := this.(type) {
	case string, fmt.Stringer:
		return fmt.Sprintf("Value[%T]{This: \"%s\"}", this, this)
	default:
		return fmt.Sprintf("Value[%T]{This: %v}", this, this)
	}

}
