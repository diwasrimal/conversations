package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"
)

func LogoutGet(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Couldn't find cookie with session credentials"})
		return
	}
	log.Println("got cookie", cookie)
	sessionId := cookie.Value
	if len(sessionId) == 0 {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Invalid session credentials for logging out"})
		return
	}
	err = db.DeleteUserSession(sessionId)
	if err != nil {
		log.Printf("Error deleting user session from db: %v\n", err)
		utils.SendJsonResp(w, http.StatusInternalServerError, types.Json{"message": "Error removing login credentials"})
		return
	}
	utils.SendJsonResp(w, http.StatusOK, types.Json{})
}
