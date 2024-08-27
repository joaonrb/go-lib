package convert

import "github.com/joaonrb/go-lib/types"

func ResultToErrorOption[T any](result types.Result[T]) (option types.Option[error]) {
	result.
		WhenOK(func(t T) {
			option = types.Nothing[error]{}
		}).
		WhenError(func(err error) {
			option = types.Value[error]{This: err}
		})
	return
}
