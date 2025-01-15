package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User

	//read the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request"+err.Error(), http.StatusBadRequest)
		return
	}

	//check if username is valid
	if !user.IsValid() {
		http.Error(w, "INVALID USERNAME", http.StatusBadRequest)
		return

	}

	/*check if user exist
	if exist return user
	else create a new user*/
	exist, err := rt.db.ExistName(user.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't check the username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !exist {
		user, err = rt.CreateUser(user)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't create the user")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusCreated)
	} else {
		dbUser, err := rt.db.GetUserByName(user.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't get the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = user.FromDatabase(dbUser)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't  convert the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	type AuthUser struct {
		User  User `json:"user"`
		Token int  `json:"token"`
	}

	authUser := AuthUser{user, user.UserID}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(authUser); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
