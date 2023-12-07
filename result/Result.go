package result

type Result[R any, E error] interface {
	result()
	Then(func(R) Result[R, E]) Result[R, E]
	Error(func(E))
}
