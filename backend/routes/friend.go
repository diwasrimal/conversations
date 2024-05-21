package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/types"
	"github.com/diwasrimal/conversations/backend/utils"
)

// Records mutual friendship among requesting user and given
// user. Accepts json payload with field "targetId",
// which is the user that will be befriended.
// Should be used with auth middleware.
func FriendPost(w http.ResponseWriter, r *http.Request) api.Response {
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

	status, err := db.GetFriendshipStatus(userId, uint64(targetId))
	if err != nil {
		log.Printf("Error checking friendship status while creating new friend: %v\n", err)
		return api.Response {
			Code: http.StatusInternalServerError,
			Payload: types.Json{},
		}
	}
	if status != "req-received" {
		return api.Response{
			Code: http.StatusBadRequest,
			Payload: types.Json{"message": "No request was received from other user"},
		}
	}
	
	err = db.RecordFriendship(userId, uint64(targetId))
	if err != nil {
		log.Printf("Error recording friendship among (%v, %v) in db: %v\n", userId, targetId, err)
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
