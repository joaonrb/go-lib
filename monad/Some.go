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

func (some Some[T]) Is(value T) bool {
	var s, t any = some.Value, value
	return s == t
}

func (some Some[T]) IsIn(values ...T) bool {
	var (
		s any = some.Value
		t any
	)
	result := false
	for i := 0; !result && i < len(values); i++ {
		t = values[i]
		result = result || s == t
	}
	return result
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
