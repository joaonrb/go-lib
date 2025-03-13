package result

import (
	"fmt"
	"reflect"
)

type OK[T any, E error] struct {
	Value T
}

func (ok OK[T, E]) result() {}

func (ok OK[T, E]) Then(call func(T) Result[T, E]) Result[T, E] {
	return call(ok.Value)
}

func (ok OK[T, E]) Error(func(E)) {}

func (ok OK[T, E]) String() string {
	return fmt.Sprintf("OK[%s]{Value: %v}", reflect.TypeOf(ok.Value).Name(), ok.Value)
}
