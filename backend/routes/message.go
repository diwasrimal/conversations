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
	id, ok := body["pairId"].(float64)
	if !ok {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Invalid data about chat pair"},
		}
	}
	pairId := uint64(id)
	messages, err := db.GetMessagesAmong(userId, pairId)
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
