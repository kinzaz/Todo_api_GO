package request

import (
	"net/http"
	"todo_GO/pkg/response"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		response.Json(w, body, 402)
		return nil, err
	}
	return &body, nil
}
