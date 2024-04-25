package middleware

import (
	"log"
	"mime"
	"net/http"
)

// Enforces requests to come in JSON format by checking the
// Content-Type header and responding with error if not set
func EnforceJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Checking content type for enforcing JSON")
		contentType := r.Header.Get("Content-Type")
		mt, _, err := mime.ParseMediaType(contentType)

		if err != nil {
			http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
			return
		}
		if mt != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// func Jsonifier(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Middleware logic...
// 		next.ServeHTTP(w, r)
// 	})
// }
