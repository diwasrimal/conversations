package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/middleware"
	"github.com/diwasrimal/gochat/backend/routes"

	"github.com/rs/cors"
)

/*
- POST /api/login
- POST /api/register
- GET /api/messages/:id/:receiver_id
- GET /api/user/:id/profile
- PUT /api/user/:id/profile
*/

func main() {
	db.MustInit()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/login", routes.LoginPost)
	mux.HandleFunc("GET /api/logout", routes.LogoutGet)
	mux.HandleFunc("POST /api/register", routes.RegisterPost)

	profileGetHandler := middleware.UseAuth(http.HandlerFunc(routes.ProfileGet))
	profilePutHandler := middleware.UseAuth(http.HandlerFunc(routes.ProfilePut))
	mux.Handle("GET /api/profile", profileGetHandler)
	mux.Handle("PUT /api/profile", profilePutHandler)

	// messagesGetHandler := middleware.UseAuth(http.HandlerFunc(routes.MessagesGet))
	// mux.Handle("GET /api/messages/{sender_id}/{receiver_id}", messagesGetHandler)

	handler := cors.AllowAll().Handler(middleware.UseJson(mux))

	port := 3030
	addr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, handler))
}
