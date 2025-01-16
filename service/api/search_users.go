package api

import (
	"encoding/json"
	"net/http"

	"myWasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get search string from request
	search := r.URL.Query().Get("search")
	if search == "" {
		BadRequest(w, nil, ctx, "Invalid search string")
		return
	}

	// search user
	dbUsers, err := rt.db.SearchUser(search)
	if err != nil {
		InternalServerError(w, err, "Failed to search user", ctx)
		return
	}

	// response
	users := make([]User, len(dbUsers))
	for i, u := range dbUsers {
		var user User
		err := user.FromDatabase(u)
		if err != nil {
			InternalServerError(w, err, "Failed to get user", ctx)
			return
		}
		users[i] = user
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		InternalServerError(w, err, "Failed to encode response", ctx)
		return
	}

}
