package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/rs/cors"
)

type Response struct {
	success bool
	message string
}

var conn *pgx.Conn

func main() {
	conn = initDatabase()
	defer conn.Close(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/login", handleLogin)
	mux.HandleFunc("POST /api/register", handleRegister)

	// corsOptions := cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{http.MethodGet, http.MethodPost},
	// }
	// handler := cors.New(corsOptions).Handler(mux)
	handler := cors.AllowAll().Handler(mux)

	port := 3030
	addr := fmt.Sprintf(":%v", port)
	log.Printf("Listening on port %v...\n", port)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc, dec := json.NewEncoder(w), json.NewDecoder(r.Body)

	// Decode username and password
	data := make(map[string]string)
	err := dec.Decode(&data)
	if err != nil {
		log.Printf("Error decoding: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		err = enc.Encode(map[string]any{
			"success": false,
			"message": "Couldn't parse json data",
		})
		mustNotErr(err)
		return
	}

	// Validate
	username := strings.Trim(data["username"], " \t\n")
	password := data["password"]
	if len(username) == 0 || len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err := enc.Encode(map[string]any{
			"success": false,
			"message": "Username or password is empty",
		})
		mustNotErr(err)
		return
	}

	// Check from database
	rows, err := conn.Query(
		context.Background(),
		"SELECT * FROM users WHERE username = $1 AND password = $2",
		username,
		password,
	)
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		err := enc.Encode(map[string]any{})
		mustNotErr(err)
	}
	userExists := rows.Next()
	rows.Close()

	if userExists {
		log.Printf("Successful login for %v\n", username)
		w.WriteHeader(http.StatusOK)
		err := enc.Encode(map[string]any{
			"success": true,
			"message": "Successful login",
		})
		mustNotErr(err)
	} else {
		log.Printf("Unsuccessful login for %v\n", username)
		w.WriteHeader(http.StatusUnauthorized)
		err := enc.Encode(map[string]any{
			"success": false,
			"message": "Invalid username and/or password",
		})
		mustNotErr(err)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc, dec := json.NewEncoder(w), json.NewDecoder(r.Body)

	// Decode username and password
	data := make(map[string]string)
	err := dec.Decode(&data)
	if err != nil {
		log.Printf("Error decoding: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		err = enc.Encode(map[string]any{
			"success": false,
			"message": "Couldn't parse json data",
		})
		mustNotErr(err)
		return
	}

	// Validate
	username := strings.Trim(data["username"], " \t\n")
	password := data["password"]
	if len(username) == 0 || len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err := enc.Encode(map[string]any{
			"success": false,
			"message": "Username or password is empty",
		})
		mustNotErr(err)
		return
	}

	// Record into database
	// Check if username is already taken
	rows, err := conn.Query(context.Background(), "SELECT username FROM users WHERE username = $1", username)
	if err != nil {
		log.Printf("Error looking for existing usernames: %v\n", err)
	}
	usernameTaken := rows.Next()
	rows.Close()

	if usernameTaken {
		log.Printf("Username: %q not already taken!\n", username)
		w.WriteHeader(http.StatusConflict)
		err := enc.Encode(map[string]any{
			"success": false,
			"message": "Username not available",
		})
		mustNotErr(err)
		return
	}

	// Register new user with given data
	_, err = conn.Exec(
		context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		username,
		password,
	)
	if err != nil {
		log.Printf("Error inserting to database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		err := enc.Encode(map[string]any{})
		mustNotErr(err)
		return
	}

	// Send successful response
	log.Printf("Registered user: %v, password: %v\n", username, password)
	w.WriteHeader(http.StatusCreated)
	enc.Encode(map[string]any{
		"success": true,
		"message": "Registered!",
	})
}

func initDatabase() *pgx.Conn {
	dburl := os.Getenv("DATABASE_URL")
	if len(dburl) == 0 {
		panic("DATABASE_URL not set")
	}
	fmt.Println("dburl:", dburl)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dburl)
	mustNotErr(err)
	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			username TEXT NOT NULL PRIMARY KEY,
			password TEXT NOT NULL
		)
	`)
	mustNotErr(err)
	return conn
}

// Checks the error and panics if it exists
func mustNotErr(e error) {
	if e != nil {
		panic(e)
	}
}
