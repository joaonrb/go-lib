package result

func Map[R1 any, R2 any, E error](
	r Result[R1, E],
	call func(result R1) Result[R2, E],
) (r2 Result[R2, E]) {
	r.Then(func(r1 R1) Result[R1, E] {
		r2 = call(r1)
		return r
	}).Error(func(e E) {
		r2 = Error[R2, E]{Err: e}
	})
	return
}
