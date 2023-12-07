package maybe

type Maybe[T any] interface {
	maybe()
	Then(call func(T) Maybe[T]) Maybe[T]
	IfNothing(func())
}
