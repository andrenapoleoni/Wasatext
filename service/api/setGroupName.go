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
	// check autorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		Forbidden(w, err, ctx, "Forbidden")
		return
	}
	// take data from the body request
	var g Group
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		BadRequest(w, err, ctx, "Invalid request body")
		return
	}
	// get groupID from request
	groupID, err := strconv.Atoi(ps.ByName("group"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid group id")
		return
	}
	g.GroupID = groupID

	// check if groupname is valid
	if !g.IsValid() {
		BadRequest(w, nil, ctx, "Invalid group name")
		return
	}

	// change groupname
	if err := rt.db.ChangeGroupName(g.GroupID, g.Name); err != nil {
		InternalServerError(w, err, "Failed to change groupname", ctx)
		return
	}
	// response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode("Groupname changed succesfully"); err != nil {
		InternalServerError(w, err, "Failed to encode response", ctx)
		return
	}

}
