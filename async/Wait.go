package async

func Wait[T any](promises ...Future[T]) Future[<-chan Future[T]] {
	return Run(func() <-chan Future[T] {
		result := make(chan Future[T])
		go func() {
			defer close(result)
			race(promises, func(i int, value Future[T]) {
				result <- value
			})
		}()
		return result
	})
}
