package convert

import (
	"fmt"

	"github.com/joaonrb/go-lib/monad"
)

func ToResult[T1 any, T2 any](value monad.Result[T1]) Result[T1, T2] {
	return result[T1, T2]{Result: value}
}

type Result[T1 any, T2 any] interface {
	Then(call func(T1) monad.Result[T2]) monad.Result[T2]
}

type result[T1 any, T2 any] struct {
	Result monad.Result[T1]
}

func (rc result[T1, T2]) Then(call func(T1) monad.Result[T2]) (value monad.Result[T2]) {
	rc.Result.
		WhenOK(func(t T1) {
			value = call(t)
		}).
		WhenError(func(err error) {
			value = monad.Error[T2]{Err: err}
		})
	return
}
func (rc result[T1, T2]) String() (str string) {
	var (
		value1 T1
		value2 T2
	)
	rc.Result.
		WhenOK(func(t T1) {

			switch value := any(t).(type) {
			case string, fmt.Stringer:
				str = fmt.Sprintf("OK[%T, %T]{Some: \"%s\"}", value, value2, value)
			default:
				str = fmt.Sprintf("OK[%T, %T]{Some: %v}", value, value2, value)
			}
		}).
		WhenError(func(err error) {
			str = fmt.Sprintf("Error[%T, %T]{Err: \"%s\"}", value1, value2, err)
		})
	return
}
