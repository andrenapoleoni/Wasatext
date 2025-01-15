package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	//check authorization
	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	//check group id
	groupId, err := strconv.Atoi(ps.ByName("group"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	//check if group exist
	exist, err := rt.db.ExistGroup(groupId)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't check the group")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !exist {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}
	//BODY
	var newuser User
	if err := json.NewDecoder(r.Body).Decode(&newuser); err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	//check if user exist
	exist, err = rt.db.ExistName(newuser.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't check the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exist {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	//take user by name
	newusergroup, err := rt.db.GetUserByName(newuser.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newuser.UserID = newusergroup.UserID

	//check if user is already in the group
	exist, err = rt.db.ExistUserInGroup(groupId, newuser.UserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get the group")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exist {
		http.Error(w, "User already in the group", http.StatusConflict)
		return

	}

	//add user to the group
	err = rt.db.AddUserToGroup(groupId, newuser.UserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't add user to the group")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode("Member added"); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
