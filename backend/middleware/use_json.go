package middleware

import (
	"mime"
	"net/http"

	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"
)

func UseJson(nextHandler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		mt, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"messsage": "Malformed Content-Type header"})
			return
		}
		if mt != "application/json" {
			utils.SendJsonResp(
				w,
				http.StatusUnsupportedMediaType,
				types.Json{"messsage": "Content-Type must be application/json"},
			)
			return
		}
		nextHandler.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}
