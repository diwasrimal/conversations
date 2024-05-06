package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
	"github.com/diwasrimal/gochat/backend/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ParseJson(r.Body)
	log.Printf("Login request with body: %v\n", body)
	if err != nil {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Couldn't parse request body as json"})
		return
	}

	// Ensure both username and password are given
	username, usernameOk := body["username"].(string)
	password, passwordOk := body["password"].(string)
	if !usernameOk || !passwordOk {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Missing data"})
		return
	}
	username = strings.Trim(username, " \t\n\r")
	if len(username) == 0 {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "Username is empty"})
		return
	}

	// Retreive user details from username and check password
	user, err := db.GetUserByUsername(username)
	if err != nil {
		log.Printf("Error getting user details from db: %v\n", err)
		utils.SendJsonResp(w, http.StatusInternalServerError, types.Json{"message": "Error logging in"})
		return
	}
	if user == nil {
		utils.SendJsonResp(w, http.StatusBadRequest, types.Json{"message": "No such username exists"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		utils.SendJsonResp(w, http.StatusUnauthorized, types.Json{"message": "Incorrect password"})
		return
	}

	// Create a session and send a cookie with session id
	sessionId := uuid.New().String()
	err = db.CreateUserSession(user.Id, sessionId)
	if err != nil {
		log.Printf("Error creating session in db: %v\n", err)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionId",
			Value:    sessionId,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
	}
	log.Println("Logged in!")
	utils.SendJsonResp(w, http.StatusAccepted, types.Json{})
}
