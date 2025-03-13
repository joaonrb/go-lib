package op

type Equatable[T any] interface {
	Equal(other T) bool
}
