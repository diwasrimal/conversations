package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/middleware"
	"github.com/diwasrimal/conversations/backend/routes"

	"github.com/rs/cors"
)

func main() {
	db.MustInit()
	defer db.Close()

	handlers := map[string]http.Handler{
		"GET /api/logout":            api.MakeHandler(routes.LogoutGet),
		"GET /api/tmp":               api.MakeHandler(routes.TmpGet),
		"POST /api/login":            api.MakeHandler(routes.LoginPost),
		"POST /api/register":         api.MakeHandler(routes.RegisterPost),
		"GET /api/profile":           middleware.UseAuth(api.MakeHandler(routes.ProfileGet)),
		"PUT /api/profile":           middleware.UseAuth(api.MakeHandler(routes.ProfilePut)),
		"GET /api/chat-partners":     middleware.UseAuth(api.MakeHandler(routes.ChatPartnersGet)),
		"GET /api/search":            middleware.UseAuth(api.MakeHandler(routes.SearchGet)),
		"GET /api/messages/{pairId}": middleware.UseAuth(api.MakeHandler(routes.MessagesGet)),

		"GET /api/friendship-status/{targetId}": middleware.UseAuth(api.MakeHandler(routes.FriendshipStatusGet)),
		"POST /api/friend-requests":             middleware.UseAuth(api.MakeHandler(routes.FriendRequestPost)),
		"DELETE /api/friend-requests":           middleware.UseAuth(api.MakeHandler(routes.FriendRequestDelete)),
		"GET /api/friends":                      middleware.UseAuth(api.MakeHandler(routes.FriendsGet)),
		"POST /api/friends":                     middleware.UseAuth(api.MakeHandler(routes.FriendPost)),
		"DELETE /api/friends":                   middleware.UseAuth(api.MakeHandler(routes.FriendDelete)),
		"GET /api/friend-requestors":            middleware.UseAuth(api.MakeHandler(routes.FriendRequestorsGet)),
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
