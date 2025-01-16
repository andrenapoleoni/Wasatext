package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User

	// read the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		BadRequest(w, err, ctx, "Invalid request body")
		return
	}

	// check if username is valid
	if !user.IsValid() {
		BadRequest(w, nil, ctx, "Invalid username")
		return

	}

	// check if user exist
	exist, err := rt.db.ExistName(user.Username)
	if err != nil {
		InternalServerError(w, err, "Failed to check user existence", ctx)
		return
	}

	if !exist {
		user, err = rt.CreateUser(user)
		if err != nil {
			InternalServerError(w, err, "Failed to create user", ctx)
			return

		}
		w.WriteHeader(http.StatusCreated)
	} else {
		dbUser, err := rt.db.GetUserByName(user.Username)
		if err != nil {
			InternalServerError(w, err, "Failed to get user", ctx)
			return
		}
		err = user.FromDatabase(dbUser)
		if err != nil {
			InternalServerError(w, err, "Failed to get user", ctx)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	type AuthUser struct {
		User  User `json:"user"`
		Token int  `json:"token"`
	}

	authUser := AuthUser{user, user.UserID}
	// response
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(authUser); err != nil {
		InternalServerError(w, err, "Failed to encode response", ctx)
		return
	}
}
