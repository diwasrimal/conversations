package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/rs/cors"
)

func main() {
	dburl := os.Getenv("DATABASE_URL")
	if len(dburl) == 0 {
		panic("DATABASE_URL not set")
	}
	fmt.Println("dburl:", dburl)
	conn, err := pgx.Connect(context.Background(), dburl)
	checkErr(err)
	fmt.Println("conn:", conn)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/login", loginHandler)
	mux.HandleFunc("POST /api/register", registerHandler)

	corsOptions := cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5173"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	}
	handler := cors.New(corsOptions).Handler(mux)

	port := 3030
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("Listening on port %v...\n", port)
	http.ListenAndServe(addr, handler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("checking and logging...")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !r.Form.Has("username") || !r.Form.Has("password") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username or password missing"))
		return
	}
	username := strings.Trim(r.Form.Get("username"), " \t\n")
	password := r.Form.Get("password")
	fmt.Printf("Registering user: %v, pass: %v\n", username, password)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
