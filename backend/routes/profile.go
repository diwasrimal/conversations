package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/models"
	"github.com/diwasrimal/conversations/backend/types"
	"github.com/diwasrimal/conversations/backend/utils"
)

// Should be used with auth middleware
func ProfileGet(w http.ResponseWriter, r *http.Request) api.Response {
	userId := r.Context().Value("userId").(uint64)
	user, err := db.GetUserById(userId)
	if err != nil {
		log.Printf("Error getting user by id from db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error getting user data"},
		}
	}
	return api.Response{
		Code: http.StatusOK,
		Payload: types.Json{
			"fullname": user.Fullname,
			"username": user.Username,
			"bio":      user.Bio,
		},
	}
}

func ProfilePut(w http.ResponseWriter, r *http.Request) api.Response {
	body, err := utils.ParseJson(r.Body)
	log.Printf("Profile put request with body: %v\n", body)
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Couldn't parse request body as json"},
		}
	}

	// Get existing details and update with provided ones
	userId := r.Context().Value("userId").(uint64)
	user, err := db.GetUserById(userId)
	if err != nil {
		log.Printf("Error getting user by id from db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error getting user data"},
		}
	}

	var newUser models.User = *user
	update := false
	fullname, fullnameOk := body["fullname"].(string)
	bio, bioOk := body["bio"].(string)

	// TODO: Maybe just allow updating bio
	if fullnameOk && len(fullname) != 0 && fullname != user.Fullname {
		newUser.Fullname = fullname
		update = true
	}
	if bioOk && len(bio) != 0 && bio != user.Bio {
		newUser.Bio = bio
		update = true
	}

	if !update {
		return api.Response{
			Code:    http.StatusOK,
			Payload: types.Json{"message": "No new data to update"},
		}
	}

	err = db.UpdateUser(user.Id, newUser)
	if err == nil {
		log.Println("Updated user")
		return api.Response{
			Code:    http.StatusOK,
			Payload: types.Json{},
		}
	} else {
		log.Printf("Error updating user in db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error updating user data"},
		}
	}
}
