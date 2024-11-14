package request

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&payload); err != nil {
		return payload, err
	}

	return payload, nil
}
