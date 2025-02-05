package async

func Gather[T any](promises ...Future[T]) Future[[]Future[T]] {
	return Run(func() []Future[T] {
		result := make([]Future[T], len(promises))
		race(promises, func(i int, value Future[T]) {
			result[i] = value
		})
		return result
	})
}
