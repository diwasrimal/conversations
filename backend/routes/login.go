package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/types"
	"github.com/diwasrimal/conversations/backend/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginPost(w http.ResponseWriter, r *http.Request) api.Response {
	body, err := utils.ParseJson(r.Body)
	log.Printf("Hit LoginPost() with body: %v\n", body)
	if err != nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Couldn't parse request body as json"},
		}
	}

	// Ensure both username and password are given
	username, usernameOk := body["username"].(string)
	password, passwordOk := body["password"].(string)
	if !usernameOk || !passwordOk {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Missing data"},
		}
	}
	username = strings.Trim(username, " \t\n\r")
	if len(username) == 0 {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Username is empty"},
		}
	}

	// Retreive user details from username and check password
	user, err := db.GetUserByUsername(username)
	if err != nil {
		log.Printf("Error getting user details from db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{"message": "Error logging in"},
		}
	}
	if user == nil {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "No such username exists"},
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return api.Response{
			Code:    http.StatusUnauthorized,
			Payload: types.Json{"message": "Incorrect password"},
		}
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
			Path:     "/",
		})
	}
	log.Println("Logged in!")
	return api.Response{
		Code:    http.StatusAccepted,
		Payload: types.Json{"userId": user.Id},
	}
}

// Should be used with auth middleware to work as expected.
// This function assumes that authentication was handled by
// middleware and hence just returns a ok status with logged in userid
func LoginStatusGet(w http.ResponseWriter, r *http.Request) api.Response {
	userId := r.Context().Value("userId").(uint64)
	log.Printf("Login status valid for userId: %v\n", userId)
	return api.Response{
		Code:    http.StatusOK,
		Payload: types.Json{"userId": userId},
	}
}
