package monad

import (
	"fmt"
)

var _ Maybe[any] = Some[any]{}

type Some[T any] struct {
	Value T
}

func (some Some[T]) maybe() {}

func (some Some[T]) Then(call func(T) Maybe[T]) Maybe[T] {
	return call(some.Value)
}

func (some Some[T]) Else(func() Maybe[T]) Maybe[T] {
	return some
}

func (some Some[T]) WhenValue(call func(T)) Maybe[T] {
	call(some.Value)
	return some
}

func (some Some[T]) WhenNothing(func()) Maybe[T] {
	return some
}

func (some Some[T]) Or(T) Maybe[T] {
	return some
}

func (some Some[T]) If(comparator Comparator[T]) bool {
	return comparator(some.Value)
}

func (some Some[T]) DoIf(comparator Comparator[T], do func(T) Maybe[T]) Maybe[T] {
	if some.If(comparator) {
		return do(some.Value)
	}
	return some
}

func (some Some[T]) TryValue() T {
	return some.Value
}

func (some Some[T]) String() string {
	var this any = some.Value
	switch this := this.(type) {
	case string, fmt.Stringer:
		return fmt.Sprintf("Some[%T]{Value: \"%s\"}", this, this)
	default:
		return fmt.Sprintf("Some[%T]{Value: %v}", this, this)
	}

}
