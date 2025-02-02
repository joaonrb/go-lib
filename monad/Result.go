package monad

type Result[T any] interface {
	result()
	Then(call func(T) Result[T]) Result[T]
	Error(call func(error) Result[T]) Result[T]
	WhenOK(call func(T)) Result[T]
	WhenError(call func(error)) Result[T]
	Or(value T) Result[T]
	Is(value T) bool
	IsError(value error) bool
	TryValue() T
	TryError() error
}
