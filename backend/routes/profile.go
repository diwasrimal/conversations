package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/api"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/models"
	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"
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
			"fname":    user.Fname,
			"lname":    user.Lname,
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
	fname, fnameOk := body["fname"].(string)
	lname, lnameOk := body["lname"].(string)
	bio, bioOk := body["bio"].(string)

	// TODO: Maybe just allow updating bio
	if fnameOk && len(fname) != 0 && fname != user.Fname {
		newUser.Fname = fname
		update = true
	}
	if lnameOk && len(lname) != 0 && lname != user.Lname {
		newUser.Lname = lname
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
