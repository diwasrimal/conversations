package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	mw "github.com/diwasrimal/conversations/backend/middleware"
	"github.com/diwasrimal/conversations/backend/routes"

	"github.com/rs/cors"
)

func main() {
	db.MustInit()
	defer db.Close()

	handlers := map[string]http.Handler{
		"GET /api/logout":            mw.UseJson(api.MakeHandler(routes.LogoutGet)),
		"GET /api/tmp":               mw.UseJson(api.MakeHandler(routes.TmpGet)),
		"POST /api/login":            mw.UseJson(api.MakeHandler(routes.LoginPost)),
		"GET /api/login-status":      mw.UseAuth(api.MakeHandler(routes.LoginStatusGet)),
		"POST /api/register":         mw.UseJson(api.MakeHandler(routes.RegisterPost)),
		"GET /api/profile":           mw.UseAuth(mw.UseJson(api.MakeHandler(routes.ProfileGet))),
		"PUT /api/profile":           mw.UseAuth(mw.UseJson(api.MakeHandler(routes.ProfilePut))),
		"GET /api/chat-partners":     mw.UseAuth(mw.UseJson(api.MakeHandler(routes.ChatPartnersGet))),
		"GET /api/search":            mw.UseAuth(mw.UseJson(api.MakeHandler(routes.SearchGet))),
		"GET /api/messages/{pairId}": mw.UseAuth(mw.UseJson(api.MakeHandler(routes.MessagesGet))),

		"GET /api/friendship-status/{targetId}": mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendshipStatusGet))),
		"POST /api/friend-requests":             mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendRequestPost))),
		"DELETE /api/friend-requests":           mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendRequestDelete))),
		"GET /api/friends":                      mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendsGet))),
		"POST /api/friends":                     mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendPost))),
		"DELETE /api/friends":                   mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendDelete))),
		"GET /api/friend-requestors":            mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendRequestorsGet))),

		"GET /ws": mw.UseAuth(http.HandlerFunc(routes.WSHandleFunc)),
	}
	mux := http.NewServeMux()
	for route, handler := range handlers {
		mux.Handle(route, handler)
	}

	finalHandler := cors.AllowAll().Handler(mux)
	port := 3030
	addr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, finalHandler))
}
