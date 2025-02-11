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
		BadRequest(w, err, ctx, "Bad Request")
		return
	}

	userID := ctx.UserID

	// check authorization
	if profileUserID != userID {
		Forbidden(w, err, ctx, "Forbidden")
		return
	}
	// check group id
	groupId, err := strconv.Atoi(ps.ByName("group"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	// check if group exist
	exist, err := rt.db.ExistGroup(groupId)
	if err != nil {
		InternalServerError(w, err, "can't check the group", ctx)
		return
	}

	if !exist {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}
	// BODY
	type userlist struct {
		Users []User `json:"users"`
	}
	var newusers userlist
	if err := json.NewDecoder(r.Body).Decode(&newusers); err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	for _, newuser := range newusers.Users {

		// check if user exist
		exist, err = rt.db.ExistName(newuser.Username)
		if err != nil {
			InternalServerError(w, err, "can't check the user", ctx)
			return
		}
		if !exist {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		// take user by name
		newusergroup, err := rt.db.GetUserByName(newuser.Username)
		if err != nil {
			InternalServerError(w, err, "can't get the user", ctx)
			return
		}
		newuser.UserID = newusergroup.UserID

		// check if user is already in the group
		exist, err = rt.db.ExistUserInGroup(groupId, newuser.UserID)
		if err != nil {
			InternalServerError(w, err, "can't check if user is in the group", ctx)
			return
		}
		if exist {
			http.Error(w, "User already in the group", http.StatusConflict)
			return

		}

		// add user to the group
		err = rt.db.AddUserToGroup(groupId, newuser.UserID)
		if err != nil {
			InternalServerError(w, err, "can't add user to the group", ctx)
			return
		}
	}

	newlistusers, err := rt.db.GetMemberGroup(groupId)
	var newAdded userlist
	if err != nil {
		InternalServerError(w, err, "can't get the list of users", ctx)
		return
	}
	for _, u := range newlistusers {
		usDB, err := rt.db.GetUserByID(u)
		if err != nil {
			InternalServerError(w, err, "can't get the user", ctx)
			return
		}
		var us User
		err = us.FromDatabase(usDB)
		if err != nil {
			InternalServerError(w, err, "can't get the user", ctx)
			return
		}

		newAdded.Users = append(newAdded.Users, us)

	}

	// response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode(newAdded); err != nil {
		InternalServerError(w, err, "can't encode the response", ctx)
		return
	}

}
