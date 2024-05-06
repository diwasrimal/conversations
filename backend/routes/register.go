package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ParseJson(r.Body)
	log.Printf("Register request with body: %v\n", body)
	if err != nil {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Couldn't parse request body as json"})
		return
	}

	// Ensure data is provided and with reasonable lengths
	fname, fnameOk := body["fname"].(string)
	lname, lnameOk := body["lname"].(string)
	username, usernameOk := body["username"].(string)
	password, passwordOk := body["password"].(string)
	if !fnameOk || !lnameOk || !usernameOk || !passwordOk {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Missing some data"})
		return
	}
	fname = strings.Trim(fname, " \t\n\r")
	lname = strings.Trim(lname, " \t\n\r")
	username = strings.Trim(username, " \t\n\r")
	if len(fname) == 0 || len(lname) == 0 || len(username) == 0 || len(password) == 0 {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Data should not be empty"})
		return
	}

	// Hash password with bcrypt
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Password should be max 72 chars."})
		return
	}
	passwordHash := string(hashed)

	// Check if username is already taken
	taken, err := db.IsUsernameTaken(username)
	if err != nil {
		log.Printf("Error checking username's existence: %v\n", err)
		utils.SendJsonResp(w, http.StatusInternalServerError, types.Json{"message": "Error registering user"})
		return
	}
	if taken {
		utils.SendJsonResp(w, http.StatusConflict, types.Json{"message": "Username unavailable"})
		return
	}

	err = db.CreateUser(fname, lname, username, passwordHash)
	if err == nil {
		log.Println("Registered user!")
		utils.SendJsonResp(w, http.StatusCreated, types.Json{})
	} else {
		log.Printf("Error creating user in db: %v\n", err)
		utils.SendJsonResp(w, http.StatusInternalServerError, types.Json{"message": "Error registering user"})
	}
}
