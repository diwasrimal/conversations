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
		"GET /api/login-status":      mw.UseAuth(api.MakeHandler(routes.LoginStatusGet)),
		"GET /api/users/{id}":        mw.UseAuth(mw.UseJson(api.MakeHandler(routes.UsersGet))),
		"GET /api/chat-partners":     mw.UseAuth(mw.UseJson(api.MakeHandler(routes.ChatPartnersGet))),
		"GET /api/search":            mw.UseAuth(mw.UseJson(api.MakeHandler(routes.SearchGet))),
		"GET /api/messages/{pairId}": mw.UseAuth(mw.UseJson(api.MakeHandler(routes.MessagesGet))),
		"GET /api/friends":           mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendsGet))),
		"GET /api/friend-requestors": mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendRequestorsGet))),
		"GET /ws":                    mw.UseAuth(http.HandlerFunc(routes.WSHandleFunc)),
		"GET /api/tmp":               mw.UseJson(api.MakeHandler(routes.TmpGet)),

		"GET /api/friendship-status/{targetId}": mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendshipStatusGet))),

		"POST /api/login":           mw.UseJson(api.MakeHandler(routes.LoginPost)),
		"POST /api/register":        mw.UseJson(api.MakeHandler(routes.RegisterPost)),
		"POST /api/friend-requests": mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendRequestPost))),
		"POST /api/friends":         mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendPost))),

		"DELETE /api/friend-requests": mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendRequestDelete))),
		"DELETE /api/friends":         mw.UseAuth(mw.UseJson(api.MakeHandler(routes.FriendDelete))),
	}
	mux := http.NewServeMux()
	for route, handler := range handlers {
		mux.Handle(route, handler)
	}

	// File server to serve frontend build files
	mux.Handle("/", http.FileServer(http.Dir("./dist")))

	finalHandler := cors.AllowAll().Handler(mux)
	port := 3030
	addr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, finalHandler))
}
