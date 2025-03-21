package op

type Operator[T any] struct {
	Evaluator[T]
}

func (operator Operator[T]) And(other Operator[T]) Operator[T] {
	return Operator[T]{Evaluator: &And[T]{one: operator, other: other}}
}

func (operator Operator[T]) Or(other Operator[T]) Operator[T] {
	return Operator[T]{Evaluator: &Or[T]{one: operator, other: other}}
}
