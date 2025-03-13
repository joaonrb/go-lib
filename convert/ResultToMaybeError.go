package convert

import "github.com/joaonrb/go-lib/monad"

func ResultToMaybeError[T any](result monad.Result[T]) (maybe monad.Maybe[error]) {
	result.
		WhenOK(func(t T) {
			maybe = monad.Nothing[error]{}
		}).
		WhenError(func(err error) {
			maybe = monad.Some[error]{Value: err}
		})
	return
}
