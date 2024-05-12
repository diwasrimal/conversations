package routes

import (
	"log"
	"net/http"

	"github.com/diwasrimal/gochat/backend/api"
	"github.com/diwasrimal/gochat/backend/db"
	"github.com/diwasrimal/gochat/backend/types"
)

// Searches for a user by their username based on the search query
// provided search type. Search type can be "by-username" or "normal".
// Searching by username is exact, while normal search if fuzzy.
// Should be used with auth
func SearchGet(w http.ResponseWriter, r *http.Request) api.Response {
	searchType := r.URL.Query().Get("type")
	searchQuery := r.URL.Query().Get("query")
	log.Printf("Search request with params: type: %q, query: %q\n", searchType, searchQuery)
	if searchType != "normal" && searchType != "by-username" {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Invalid search type in request url params"},
		}
	}
	if len(searchQuery) == 0 {
		return api.Response{
			Code:    http.StatusBadRequest,
			Payload: types.Json{"message": "Search query cannot be empty"},
		}
	}
	matches, err := db.SearchUser(searchType, searchQuery)
	if err != nil {
		log.Printf("Error getting user search results from db: %v\n", err)
		return api.Response{
			Code:    http.StatusInternalServerError,
			Payload: types.Json{},
		}
	}
	return api.Response{
		Code:    http.StatusOK,
		Payload: types.Json{"matches": matches},
	}
}
