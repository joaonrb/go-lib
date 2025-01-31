package convert

import (
	"github.com/joaonrb/go-lib/log"
	"github.com/joaonrb/go-lib/monad"
)

func ResultToMaybe[T any](result monad.Result[T]) (maybe monad.Maybe[T]) {
	result.
		WhenOK(func(t T) {
			maybe = monad.Some[T]{Value: t}
		}).
		WhenError(func(err error) {
			var t T
			log.Debug("ResultToMaybe converted error %s to Nothing[%T]", err, t)
			maybe = monad.Nothing[T]{}
		})
	return
}
