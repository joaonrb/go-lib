package convert

import (
	"fmt"
	"github.com/joaonrb/go-lib/types"
)

func OptionType[T1 any, T2 any](option types.Option[T1]) OptionTypeConverter[T1, T2] {
	return optionType[T1, T2]{Option: option}
}

type OptionTypeConverter[T1 any, T2 any] interface {
	Then(call func(T1) types.Option[T2]) types.Option[T2]
}

type optionType[T1 any, T2 any] struct {
	Option types.Option[T1]
}

func (ot optionType[T1, T2]) Then(call func(T1) types.Option[T2]) (option types.Option[T2]) {
	ot.Option.
		WhenValue(func(t T1) {
			option = call(t)
		}).
		WhenNothing(func() {
			option = types.Nothing[T2]{}
		})
	return
}
func (ot optionType[T1, T2]) String() (str string) {
	var (
		value1 T1
		value2 T2
	)
	ot.Option.
		WhenValue(func(t T1) {
			switch value := any(t).(type) {
			case string, fmt.Stringer:
				str = fmt.Sprintf("Value[%T, %T]{This: \"%s\"}", value, value2, value)
			default:
				str = fmt.Sprintf("Value[%T, %T]{This: %v}", value, value2, value)
			}
		}).
		WhenNothing(func() {
			str = fmt.Sprintf("Nothing[%T, %T]{}", value1, value2)
		})
	return
}
