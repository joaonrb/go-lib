package monad

import "github.com/joaonrb/go-lib/op"

type Result[T any] interface {
	result()
	Then(call func(T) Result[T]) Result[T]
	Error(call func(error) Result[T]) Result[T]
	WhenOK(call func(T)) Result[T]
	WhenError(call func(error)) Result[T]
	Or(value T) Result[T]
	If(operator op.Operator[T]) bool
	DoIf(comparator op.Operator[T], do func(T) Result[T]) Result[T]
	IfError(operator op.Operator[error]) bool
	DoIfError(operator op.Operator[error], do func(error) Result[T]) Result[T]
	TryValue() T
	TryError() error
}
