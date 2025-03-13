package types

type Option[T any] interface {
	option()
	Then(call func(T) Option[T]) Option[T]
	Else(call func() Option[T]) Option[T]
	WhenValue(call func(T)) Option[T]
	WhenNothing(call func()) Option[T]
	Or(value T) Option[T]
	TryValue() T
}
