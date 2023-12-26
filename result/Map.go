package result

func Map[R1 any, R2 any, E error](
	r1 Result[R1, E],
	call func(result R1) Result[R2, E],
) (r2 Result[R2, E]) {
	r1.result()
	r1.Then(func(r R1) Result[R1, E] {
		r2 = call(r)
		return r1
	}).Error(func(e E) {
		r2 = Error[R2, E]{Err: e}
	})
	return
}
