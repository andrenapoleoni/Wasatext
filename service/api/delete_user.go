package api

import (
	"net/http"
	"strconv"

	"myWasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check authorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid user id")
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// delete user
	err = rt.db.DeleteUser(userID)
	if err != nil {
		InternalServerError(w, err, "Failed to delete user", ctx)
		return
	}

	// delete user from all groups

	// resposne
	w.WriteHeader(http.StatusNoContent)

}
