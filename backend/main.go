package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/crypto"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/middleware"
	"github.com/diwasrimal/gochat/backend/types"
	"github.com/rs/cors"
)

var activeSessions = make(map[string]*db.Session)

func main() {
	db.MustInit()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/register", handleRegister)
	mux.HandleFunc("POST /api/login", handleLogin)
	mux.HandleFunc("GET /api/homedata", handleHomeData)

	handler := cors.AllowAll().Handler(middleware.EnforceJSON(mux))
	// handler := middleware.EnforceJSON(mux)

	port := 3030
	addr := fmt.Sprintf(":%v", port)
	log.Printf("Listening on port %v...\n", 3030)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func handleHomeData(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("sessionId")
	if err != nil {
		sendJsonResp(w, http.StatusUnauthorized, types.Json{})
		return
	}
	id := sessionCookie.Value
	session, err := retrieveSession(id)
	if err != nil {
		log.Printf("Error during session retrieval: %v\n", err)
		sendJsonResp(w, http.StatusInternalServerError, types.Json{})
		return
	}
	valid := session != nil
	if !valid {
		sendJsonResp(w, http.StatusUnauthorized, types.Json{})
		return
	}

	// Just send hello msg for now
	sendJsonResp(w, http.StatusOK, types.Json{
		"success": true,
		"message": fmt.Sprintf("Hello %v", session.Username),
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	body, err := parseReqBody(r)
	if err != nil {
		sendJsonResp(w, http.StatusBadRequest, types.Json{
			"success": false,
			"message": "Couldn't parse request body",
		})
		return
	}

	username, nameok := body["username"].(string)
	password, passok := body["password"].(string)
	if !nameok || !passok {
		sendJsonResp(w, http.StatusBadRequest, types.Json{
			"success": false,
			"message": "Missing username and/or password",
		})
		return
	}

	// Password is bcrypted, so can't be larger than 72 chars
	if len(password) > 72 {
		sendJsonResp(w, http.StatusBadRequest, types.Json{
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
		sendJsonResp(w, http.StatusInternalServerError, types.Json{})
		return
	}
	if taken {
		log.Println("Username is already taken!")
		sendJsonResp(w, http.StatusConflict, types.Json{
			"success": false,
			"message": "Username already taken!",
		})
		return
	}

	hash := crypto.MustHashPassword(password)
	err = db.RecordUser(&db.User{Username: username, PasswordHash: hash})
	if err != nil {
		log.Printf("Error inserting user to database: %v\n", err)
		sendJsonResp(w, http.StatusInternalServerError, types.Json{})
		return
	}

	log.Println("Registered")
	sendJsonResp(w, http.StatusCreated, types.Json{
		"success": true,
		"message": "Registered!",
	})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	body, err := parseReqBody(r)
	if err != nil {
		sendJsonResp(w, http.StatusBadRequest, types.Json{
			"success": false,
			"message": "Couldn't parse request body",
		})
		return
	}

	username, nameok := body["username"].(string)
	password, passok := body["password"].(string)
	if !nameok || !passok {
		sendJsonResp(w, http.StatusBadRequest, types.Json{
			"success": false,
			"message": "Missing username and/or password",
		})
		return
	}

	log.Printf("Got username: %q and pass: %q for login\n", username, password)

	user, err := db.LookupUser(username)
	if err != nil {
		log.Printf("Error looking user in db: %v\n", err)
		sendJsonResp(w, http.StatusInternalServerError, types.Json{})
		return
	}
	if user == nil {
		sendJsonResp(w, http.StatusBadRequest, types.Json{
			"success": false,
			"message": "Username is not registered!",
		})
		return
	}

	if !crypto.CheckPassword(password, user.PasswordHash) {
		sendJsonResp(w, http.StatusUnauthorized, types.Json{
			"success": false,
			"message": "Incorrect password!",
		})
		return
	}

	// Record session in db and in active sessions
	sessionId := crypto.RandSessionId()
	db.RecordSession(sessionId, username)
	activeSessions[sessionId] = &db.Session{
		Id:       sessionId,
		Username: username,
	}

	cookie := http.Cookie{
		Name:     "sessionId",
		Value:    sessionId,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	log.Printf("Logging in user: %v\n", username)
	sendJsonResp(w, http.StatusOK, types.Json{
		"success": true,
		"message": "Login successful",
	})
}

// Tries to get session associated with provided id from
// active sessions or database
func retrieveSession(id string) (*db.Session, error) {
	session, ok := activeSessions[id]
	if ok {
		return session, nil
	}
	session, err := db.LookupSession(id)
	if err != nil {
		return nil, fmt.Errorf("Error looking session in db: %v", err)
	}
	return session, nil
}

func sendJsonResp(w http.ResponseWriter, status int, body types.Json) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func parseReqBody(r *http.Request) (types.Json, error) {
	data := make(types.Json)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
