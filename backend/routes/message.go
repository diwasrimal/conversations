package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/api"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"
)

// Gets messages among two users from database.
// Should be used with authentication middleware
func MessagesGet(w http.ResponseWriter, r *http.Request) api.Response {
	userId := r.Context().Value("userId").(uint64)
	body, err := utils.ParseJson(r.Body)
	log.Printf("Message get request with body: %v\n", body)
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Couldn't parse request body as json"},
		}
	}
	pairId, ok := body["pairId"].(uint64)
	log.Printf("pairId: %v, %T", pairId, pairId)
	if !ok {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Invalid message pair information in request"},
		}
	}
	messages, err := db.GetMessagesAmong(userId, pairId)
	if err != nil {
		log.Printf("Error getting messsages among (%v, %v) from db: %v\n", userId, pairId, err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error retreiving messages"},
		}
	}
	for _, msg := range messages {
		log.Println("message:", msg)
	}
	return api.Response{
		Code:    http.StatusOK,
		Payload: types.Json{"message": "Messages ..."},
	}
}
