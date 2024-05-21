package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/types"
)

// Gets messages among two users from database.
// Should be used with authentication middleware
func MessagesGet(w http.ResponseWriter, r *http.Request) api.Response {
	userId := r.Context().Value("userId").(uint64)
	log.Printf("Hit MessagesGet() with userId: %v\n", userId)
	pairId, err := strconv.Atoi(r.PathValue("pairId"))
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Invalid data about chat pair"},
		}
	}
	messages, err := db.GetMessagesAmong(userId, uint64(pairId))
	if err != nil {
		log.Printf("Error getting messsages among (%v, %v) from db: %v\n", userId, pairId, err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error retreiving messages"},
		}
	}

	return api.Response{
		Code:    http.StatusOK,
		Payload: types.Json{"messages": messages},
	}
}
