package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/api"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
)

// Gets list of conversations for a user.
// Should be used with authentication middleware.
func ConversationsGet(w http.ResponseWriter, r *http.Request) api.Response {
	userId := r.Context().Value("userId").(uint64)
	conversations, err := db.GetConversationsOf(userId)
	if err != nil {
		log.Printf("Error getting conversations of %v from db: %v\n", userId, err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error retreiving conversations"},
		}
	}
	return api.Response{
		Code:    http.StatusOK,
		Payload: types.Json{"conversations": conversations},
	}
}
