package convert

import (
	"fmt"

	"github.com/joaonrb/go-lib/monad"
)

func ToMaybe[T1 any, T2 any](value monad.Maybe[T1]) Maybe[T1, T2] {
	return maybe[T1, T2]{Maybe: value}
}

type Maybe[T1 any, T2 any] interface {
	Then(call func(T1) monad.Maybe[T2]) monad.Maybe[T2]
}

type maybe[T1 any, T2 any] struct {
	Maybe monad.Maybe[T1]
}

func (oc maybe[T1, T2]) Then(call func(T1) monad.Maybe[T2]) (value monad.Maybe[T2]) {
	oc.Maybe.
		WhenValue(func(t T1) {
			value = call(t)
		}).
		WhenNothing(func() {
			value = monad.Nothing[T2]{}
		})
	return
}
func (oc maybe[T1, T2]) String() (str string) {
	var (
		value1 T1
		value2 T2
	)
	oc.Maybe.
		WhenValue(func(t T1) {
			switch value := any(t).(type) {
			case string, fmt.Stringer:
				str = fmt.Sprintf("Some[%T, %T]{ToValue: \"%s\"}", value, value2, value)
			default:
				str = fmt.Sprintf("Some[%T, %T]{ToValue: %v}", value, value2, value)
			}
		}).
		WhenNothing(func() {
			str = fmt.Sprintf("Nothing[%T, %T]{}", value1, value2)
		})
	return
}
