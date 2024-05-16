package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/types"
)

// Get the status of friendship between requesting user
// and user mentioned in the request path. Friendship status
// is given from requesting user's point of view
// Should be used with auth middleware
func FriendshipStatusGet(w http.ResponseWriter, r *http.Request) api.Response {
	userId := r.Context().Value("userId").(uint64)
	targetId, err := strconv.Atoi(r.PathValue("targetId"))
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Invalid target user id in request"},
		}
	}
	status, err := db.GetFriendshipStatus(userId, uint64(targetId)) // status from userId's point of view
	if err != nil {
		log.Printf("Error getting friendship status from db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{},
		}
	}
	return api.Response{
		Code:    http.StatusOK,
		Payload: types.Json{"status": status},
	}
}
