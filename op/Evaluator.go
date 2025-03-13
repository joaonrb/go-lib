package op

type Evaluator[T any] interface {
	Evaluate(T) bool
}
