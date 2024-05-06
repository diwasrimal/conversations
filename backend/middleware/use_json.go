package middleware

import (
	"mime"
	"net/http"
)

func UseJson(nextHandler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		mt, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
			return
		}
		if mt != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		nextHandler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
