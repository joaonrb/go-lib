package convert

import (
	"github.com/joaonrb/go-lib/log"
	"github.com/joaonrb/go-lib/types"
)

func ResultToOption[T any](result types.Result[T]) (option types.Option[T]) {
	result.
		WhenOK(func(t T) {
			option = types.Value[T]{This: t}
		}).
		WhenError(func(err error) {
			var t T
			log.Debug("ResultToOption converted error %s to Nothing[%T]", err, t)
			option = types.Nothing[T]{}
		})
	return
}
