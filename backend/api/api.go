package api

import (
	"net/http"

	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"
)

type Response struct {
	Code    int
	Payload types.Json
}

type APIFunc func(http.ResponseWriter, *http.Request) Response

func MakeHandler(f APIFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resp := f(w, r)
		utils.SendJsonResp(w, resp.Code, resp.Payload)
	}
	return http.HandlerFunc(fn)
}

func CreateResp(status int, payload types.Json) Response {
	return Response{
		Code:    status,
		Payload: payload,
	}
}
