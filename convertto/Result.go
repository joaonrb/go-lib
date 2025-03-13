package convertto

import (
	"fmt"

	"github.com/joaonrb/go-lib/types"
)

func Result[T1 any, T2 any](result types.Result[T1]) ResultConverter[T1, T2] {
	return resultConverter[T1, T2]{Result: result}
}

type ResultConverter[T1 any, T2 any] interface {
	Then(call func(T1) types.Result[T2]) types.Result[T2]
}

type resultConverter[T1 any, T2 any] struct {
	Result types.Result[T1]
}

func (rc resultConverter[T1, T2]) Then(call func(T1) types.Result[T2]) (result types.Result[T2]) {
	rc.Result.
		WhenOK(func(t T1) {
			result = call(t)
		}).
		WhenError(func(err error) {
			result = types.Error[T2]{Err: err}
		})
	return
}
func (rc resultConverter[T1, T2]) String() (str string) {
	var (
		value1 T1
		value2 T2
	)
	rc.Result.
		WhenOK(func(t T1) {

			switch value := any(t).(type) {
			case string, fmt.Stringer:
				str = fmt.Sprintf("OK[%T, %T]{Value: \"%s\"}", value, value2, value)
			default:
				str = fmt.Sprintf("OK[%T, %T]{Value: %v}", value, value2, value)
			}
		}).
		WhenError(func(err error) {
			str = fmt.Sprintf("Error[%T, %T]{Err: \"%s\"}", value1, value2, err)
		})
	return
}
