package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/diwasrimal/gochat/backend/api"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterPost(w http.ResponseWriter, r *http.Request) api.Response {
	body, err := utils.ParseJson(r.Body)
	log.Printf("Register request with body: %v\n", body)
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Couldn't parse request body as json"},
		}
	}

	// Ensure data is provided and with reasonable lengths
	fullname, fullnameOk := body["fullname"].(string)
	username, usernameOk := body["username"].(string)
	password, passwordOk := body["password"].(string)
	if !fullnameOk || !usernameOk || !passwordOk {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Missing some data"},
		}
	}
	fullname = strings.Trim(fullname, " \t\n\r")
	username = strings.Trim(username, " \t\n\r")
	if len(fullname) == 0 || len(username) == 0 || len(password) == 0 {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Data should not be empty"},
		}
	}
	if strings.Contains(username, " \t\n\r") {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Username cannot contain spaces."},
		}
	}

	// Hash password with bcrypt
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Password should be max 72 chars."},
		}
	}
	passwordHash := string(hashed)

	// Check if username is already taken
	taken, err := db.IsUsernameTaken(username)
	if err != nil {
		log.Printf("Error checking username's existence: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error registering user"},
		}
	}
	if taken {
		return api.Response{
			Code:    http.StatusConflict,
			Payload: types.Json{"message": "Username unavailable"},
		}
	}

	err = db.CreateUser(fullname, username, passwordHash)
	if err == nil {
		log.Println("Registered user!")
		return api.Response{
			Code:    http.StatusCreated,
			Payload: types.Json{},
		}
	} else {
		log.Printf("Error creating user in db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error registering user"},
		}
	}
}
