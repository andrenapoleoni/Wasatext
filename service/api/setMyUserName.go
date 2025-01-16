package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid user id")
		return
	}

	userID := ctx.UserID

	// check authorization
	if profileUserID != userID {
		Forbidden(w, err, ctx, "Forbidden")
		return
	}
	// take data from the body request
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		BadRequest(w, err, ctx, "Invalid request body")
		return
	}
	// check if username is valid
	if !user.IsValid() {
		BadRequest(w, nil, ctx, "Invalid username")
		return
	}

	// change username, if username already taken return an error
	if err := rt.db.ChangeUsername(userID, user.Username); err != nil {
		InternalServerError(w, err, "Failed to change username: already taken", ctx)
		return
	}

	// response success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode("Username changed succesfully"); err != nil {
		InternalServerError(w, err, "Failed to encode response", ctx)
		return
	}

}
