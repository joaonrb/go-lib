package convertto

import (
	"encoding/json"

	"github.com/joaonrb/go-lib/monad"
)

func JSON(value any) monad.Result[string] {
	raw, err := json.Marshal(value)
	if err != nil {
		return monad.Error[string]{Err: err}
	}
	return monad.OK[string]{Value: string(raw)}
}

func PrettyJSON(value any) monad.Result[string] {
	raw, err := json.MarshalIndent(value, "", "    ")
	if err != nil {
		return monad.Error[string]{Err: err}
	}
	return monad.OK[string]{Value: string(raw)}
}
