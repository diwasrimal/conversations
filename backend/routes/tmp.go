package routes

import (
	"net/http"

	"github.com/diwasrimal/conversations/backend/api"
)

func TmpGet(w http.ResponseWriter, r *http.Request) api.Response {
	return api.Response{
		Code:    200,
		Payload: map[string]any{"message": "Hello from tmp route"},
	}
}
