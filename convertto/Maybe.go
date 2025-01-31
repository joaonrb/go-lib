package convertto

import (
	"fmt"

	"github.com/joaonrb/go-lib/monad"
)

func Maybe[T1 any, T2 any](maybe monad.Maybe[T1]) MaybeConverter[T1, T2] {
	return maybeConverter[T1, T2]{Maybe: maybe}
}

type MaybeConverter[T1 any, T2 any] interface {
	Then(call func(T1) monad.Maybe[T2]) monad.Maybe[T2]
}

type maybeConverter[T1 any, T2 any] struct {
	Maybe monad.Maybe[T1]
}

func (oc maybeConverter[T1, T2]) Then(call func(T1) monad.Maybe[T2]) (maybe monad.Maybe[T2]) {
	oc.Maybe.
		WhenValue(func(t T1) {
			maybe = call(t)
		}).
		WhenNothing(func() {
			maybe = monad.Nothing[T2]{}
		})
	return
}
func (oc maybeConverter[T1, T2]) String() (str string) {
	var (
		value1 T1
		value2 T2
	)
	oc.Maybe.
		WhenValue(func(t T1) {
			switch value := any(t).(type) {
			case string, fmt.Stringer:
				str = fmt.Sprintf("Some[%T, %T]{Value: \"%s\"}", value, value2, value)
			default:
				str = fmt.Sprintf("Some[%T, %T]{Value: %v}", value, value2, value)
			}
		}).
		WhenNothing(func() {
			str = fmt.Sprintf("Nothing[%T, %T]{}", value1, value2)
		})
	return
}
