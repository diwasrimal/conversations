package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/crypto"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/middleware"

	"github.com/rs/cors"
)

type Json map[string]any

func main() {
	db.MustInit()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/register", handleRegister)
	mux.HandleFunc("POST /api/login", handleLogin)
	mux.HandleFunc("POST /api/auth", handleAuth)

	handler := cors.AllowAll().Handler(middleware.EnforceJSON(mux))

	port := 3030
	addr := fmt.Sprintf(":%v", port)
	log.Printf("Listening on port %v...\n", 3030)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	body, err := parseReqBody(r)
	if err != nil {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Couldn't parse request body",
		})
		return
	}

	username, nameok := body["username"]
	password, passok := body["password"]
	if !nameok || !passok {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Missing username and/or password",
		})
		return
	}

	// Password is bcrypted, so can't be larger than 72 chars
	if len(password) > 72 {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Password must be within 72 characters",
		})
		return
	}

	log.Printf("Got username: %q and pass: %q\n", username, password)

	// Check if username is already taken
	taken, err := db.IsUsernameTaken(username)
	if err != nil {
		log.Printf("Error looking user in db: %v\n", err)
		sendJsonResp(w, http.StatusInternalServerError, Json{})
		return
	}
	if taken {
		log.Println("Username is already taken!")
		sendJsonResp(w, http.StatusConflict, Json{
			"success": false,
			"message": "Username already taken!",
		})
		return
	}

	hash := crypto.MustHashPassword(password)
	err = db.RecordUser(&db.User{Username: username, PasswordHash: hash})
	if err != nil {
		log.Printf("Error inserting user to database: %v\n", err)
		sendJsonResp(w, http.StatusInternalServerError, Json{})
		return
	}

	log.Println("Registered")
	sendJsonResp(w, http.StatusCreated, Json{
		"success": true,
		"message": "Registered!",
	})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	body, err := parseReqBody(r)
	if err != nil {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Couldn't parse request body",
		})
		return
	}

	log.Println("Message body:", body)
	username, nameok := body["username"]
	password, passok := body["password"]
	if !nameok || !passok {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Missing username and/or password",
		})
		return
	}

	log.Printf("Got username: %q and pass: %q for login\n", username, password)

	user, err := db.LookupUser(username)
	if err != nil {
		log.Printf("Error looking user in db: %v\n", err)
		sendJsonResp(w, http.StatusInternalServerError, Json{})
		return
	}
	if user == nil {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Username is not registered!",
		})
		return
	}

	if !crypto.CheckPassword(password, user.PasswordHash) {
		sendJsonResp(w, http.StatusUnauthorized, Json{
			"success": false,
			"message": "Incorrect password!",
		})
		return
	}

	log.Printf("Logging in user: %v\n", username)
	sessionId := crypto.RandSessionId()
	db.RecordSession(sessionId, username)
	sendJsonResp(w, http.StatusOK, Json{
		"success": true,
		"message": "Login successful", "sessionId": sessionId,
	})
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	body, err := parseReqBody(r)
	if err != nil {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Couldn't parse request body",
		})
		return
	}

	sessionId, ok := body["sessionId"]
	if !ok {
		sendJsonResp(w, http.StatusBadRequest, Json{
			"success": false,
			"message": "Must provide session id for authentication",
		})
		return
	}

	session, err := db.LookupSession(sessionId)
	if err != nil {
		log.Printf("Error looking session in db: %v\n", err)
		sendJsonResp(w, http.StatusInternalServerError, Json{})
		return
	}

	// No such session exists
	if session == nil {
		sendJsonResp(w, http.StatusUnauthorized, Json{})
		return
	}

	log.Printf("user: %v logged via session id: %v\n", session.Username, sessionId)
	sendJsonResp(w, http.StatusOK, Json{
		"success": true,
		"message": "Authorized",
	})
}

func sendJsonResp(w http.ResponseWriter, status int, body Json) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func parseReqBody(r *http.Request) (map[string]string, error) {
	data := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
