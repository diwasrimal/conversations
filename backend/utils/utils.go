package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/diwasrimal/gochat/backend/types"
)

func ParseJson(body io.ReadCloser) (types.Json, error) {
	data := make(types.Json)
	err := json.NewDecoder(body).Decode(&data)
	return data, err
}

func SendJsonResp(w http.ResponseWriter, status int, payload types.Json) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
