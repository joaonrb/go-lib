package convertto

import (
	"encoding/json"

	"github.com/joaonrb/go-lib/types"
)

func JSON(value any) types.Result[string] {
	raw, err := json.Marshal(value)
	if err != nil {
		return types.Error[string]{Err: err}
	}
	return types.OK[string]{Value: string(raw)}
}

func PrettyJSON(value any) types.Result[string] {
	raw, err := json.MarshalIndent(value, "", "    ")
	if err != nil {
		return types.Error[string]{Err: err}
	}
	return types.OK[string]{Value: string(raw)}
}
