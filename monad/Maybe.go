package monad

type Maybe[T any] interface {
	maybe()
	Then(call func(T) Maybe[T]) Maybe[T]
	Else(call func() Maybe[T]) Maybe[T]
	WhenValue(call func(T)) Maybe[T]
	WhenNothing(call func()) Maybe[T]
	Or(value T) Maybe[T]
	If(comparator Comparator[T]) bool
	DoIf(comparator Comparator[T], do func(T) Maybe[T]) Maybe[T]
	TryValue() T
}

func NewMaybe[T any](value *T) Maybe[T] {
	if value == nil {
		return Nothing[T]{}
	}
	return Some[T]{Value: *value}
}
