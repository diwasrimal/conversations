package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/types"
	"github.com/diwasrimal/conversations/backend/utils"
)

// Records a new entry into the friend requests table.
// Accepts json payload with field "targetId", which is the user
// that will receive this friend request. Requestor is the one
// who made this request, i.e. the logged in user.
// Should be used with auth middleware
func FriendRequestPost(w http.ResponseWriter, r *http.Request) api.Response {
	body, err := utils.ParseJson(r.Body)
	log.Printf("Friend request with body: %v\n", body)
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Couldn't parse request body as json"},
		}
	}
	userId := r.Context().Value("userId").(uint64)
	targetId, ok := body["targetId"].(float64)
	if !ok {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Missing/Invalid targetId in body"},
		}
	}

	err = db.RecordFriendRequest(userId, uint64(targetId)) // from userId -> targetId
	if err != nil {
		log.Printf("Error recording friend request in db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{},
		}
	}
	return api.Response{
		Code:    http.StatusCreated,
		Payload: types.Json{},
	}
}

func FriendRequestDelete(w http.ResponseWriter, r *http.Request) api.Response {
	log.Printf("Friend req delte request with body: %v\n", r.Body)
	return api.Response{
		Code:    http.StatusNotImplemented,
		Payload: map[string]any{},
	}
}
