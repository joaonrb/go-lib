package maybe

func Map[T1, T2 any](m1 Maybe[T1], call func(T1) Maybe[T2]) (m2 Maybe[T2]) {
	m1.maybe()
	m1.Then(func(v T1) Maybe[T1] {
		m2 = call(v)
		return m1
	}).IfNothing(func() {
		m2 = Nothing[T2]{}
	})
	return
}
