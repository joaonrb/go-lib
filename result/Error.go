package result

import (
	"fmt"
	"reflect"
)

type Error[T any, E error] struct {
	Err E
}

func (err Error[T, E]) result() {}

func (err Error[T, E]) Then(func(T) Result[T, E]) Result[T, E] {
	return err
}

func (err Error[T, E]) Error(call func(E)) {
	call(err.Err)
}

func (err Error[T, E]) String() string {
	return fmt.Sprintf("Error[%s]{Err: %v}", reflect.TypeOf(err.Error).Name(), err.Err)
}
