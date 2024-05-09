package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/api"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
)

func LogoutGet(w http.ResponseWriter, r *http.Request) api.Response {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Couldn't find cookie with session credentials"},
		}
	}
	log.Println("got cookie", cookie)
	sessionId := cookie.Value
	if len(sessionId) == 0 {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Invalid session credentials for logging out"},
		}
	}
	err = db.DeleteUserSession(sessionId)
	if err != nil {
		log.Printf("Error deleting user session from db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error removing login credentials"},
		}
	}
	return api.Response{
		Code:    http.StatusOK,
		Payload: types.Json{},
	}
}
