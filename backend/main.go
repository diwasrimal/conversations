package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/api"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/middleware"
	"github.com/diwasrimal/gochat/backend/routes"

	"github.com/rs/cors"
)

func main() {
	db.MustInit()
	defer db.Close()

	handlers := map[string]http.Handler{
		"POST /api/login":            api.MakeHandler(routes.LoginPost),
		"POST /api/logout":           api.MakeHandler(routes.LogoutGet),
		"POST /api/register":         api.MakeHandler(routes.RegisterPost),
		"GET /api/tmp":               api.MakeHandler(routes.TmpGet),
		"GET /api/profile":           middleware.UseAuth(api.MakeHandler(routes.ProfileGet)),
		"PUT /api/profile":           middleware.UseAuth(api.MakeHandler(routes.ProfilePut)),
		"GET /api/messages/{pairId}": middleware.UseAuth(api.MakeHandler(routes.MessagesGet)),
		"GET /api/conversations":     middleware.UseAuth(api.MakeHandler(routes.ConversationsGet)),
	}
	mux := http.NewServeMux()
	for route, handler := range handlers {
		mux.Handle(route, handler)
	}

	finalHandler := cors.AllowAll().Handler(middleware.UseJson(mux))
	port := 3030
	addr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, finalHandler))
}
