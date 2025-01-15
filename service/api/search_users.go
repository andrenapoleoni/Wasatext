package api

import (
	"encoding/json"
	"net/http"

	"myWasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	/*//check autorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}*/
	//get search string from request
	search := r.URL.Query().Get("search")
	if search == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//search user
	dbUsers, err := rt.db.SearchUser(search)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't search user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//response
	users := make([]User, len(dbUsers))
	for i, u := range dbUsers {
		var user User
		err := user.FromDatabase(u)
		if err != nil {
			ctx.Logger.Error("Error converting users ", err)
			http.Error(w, "Error converting users ", http.StatusInternalServerError)
			return
		}
		users[i] = user
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		ctx.Logger.Error("Error encoding users ", err)
		http.Error(w, "Error encoding response ", http.StatusInternalServerError)
		return
	}

}
