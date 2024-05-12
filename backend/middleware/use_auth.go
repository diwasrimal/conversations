package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/types"
	"github.com/diwasrimal/conversations/backend/utils"
)

// Authorizes request by validating session id contained in the cookie
// and adds user id assosiated with the session to the request context.
func UseAuth(nextHandler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionId")
		if err != nil {
			utils.SendJsonResp(w, http.StatusUnauthorized, types.Json{"message": "Missing session cookie"})
			return
		}
		sessionId := cookie.Value
		if len(sessionId) == 0 {
			utils.SendJsonResp(w, http.StatusUnauthorized, types.Json{"message": "Invalid session credentials"})
			return
		}
		session, err := db.GetSession(sessionId)
		if err != nil {
			log.Printf("Error getting session from db: %v\n", err)
			utils.SendJsonResp(w, http.StatusInternalServerError, types.Json{"message": "Error validating session credentials"})
			return
		}
		if session == nil {
			utils.SendJsonResp(w, http.StatusUnauthorized, types.Json{"message": "Invalid session credentials"})
			return
		}

		ctx := context.WithValue(r.Context(), "userId", session.UserId)
		nextHandler.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
