package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/diwasrimal/conversations/backend/api"
	"github.com/diwasrimal/conversations/backend/db"
	mw "github.com/diwasrimal/conversations/backend/middleware"
	"github.com/diwasrimal/conversations/backend/routes"

	"github.com/rs/cors"
)

const addr = ":3030"

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

	env, ok := os.LookupEnv("MODE")
	if !ok {
		panic("Environment variable 'MODE' not set, set it to \"dev\" or \"prod\"")
	}

	switch env {
	case "dev":
		// Allow cross origin requests in dev mode
		finalHandler := cors.AllowAll().Handler(mux)
		log.Printf("Server running on %v...\n", addr)
		log.Fatal(http.ListenAndServe(addr, finalHandler))
	case "prod":
		// Use a file server to serve frontend build files in production
		// also redirect all other routes to /index.html so that react handles it
		distDir := "./dist"
		fileServer := http.FileServer(http.Dir(distDir))
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join(distDir, r.URL.Path)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
				return
			}
			fileServer.ServeHTTP(w, r)
		})
		log.Printf("Server running on %v...\n", addr)
		log.Fatal(http.ListenAndServe(addr, mux))
	default:
		panic("Invalid enviroment variable value for 'MODE'")
	}

}
