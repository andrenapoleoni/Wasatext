package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ChangeGroupName changes the name of a group.
func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check autorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	//take data from the body request
	var g Group
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	//get groupID from request
	groupID, err := strconv.Atoi(ps.ByName("group"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	g.GroupID = groupID

	//check if groupname is valid
	if !g.IsValid() {
		http.Error(w, "INVALID GROUPNAME", http.StatusBadRequest)
		return
	}

	//change groupname
	if err := rt.db.ChangeGroupName(g.GroupID, g.Name); err != nil {
		http.Error(w, "groupname already taken, please retry", http.StatusBadRequest)
		return
	}
	//response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode("Groupname changed succesfully"); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
