package convert

import "github.com/joaonrb/go-lib/types"

func ResultToOption[T any](result types.Result[T]) (option types.Option[T]) {
	result.
		WhenOK(func(t T) {
			option = types.Nothing[T]{}
			option = types.Value[T]{This: err}
		}).
		WhenError(func(err error) {
			log
			option = types.Nothing[T]{}
		})
	return
}
