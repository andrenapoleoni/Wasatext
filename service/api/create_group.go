package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) CreateGroupDB(g Group, userID int) (Group, error) {
	dbGroup, err := rt.db.CreateGroup(g.ToDatabase(), userID)
	if err != nil {
		return g, err
	}

	err = g.FromDatabase(dbGroup)
	if err != nil {
		return g, err
	}

	return g, nil
}

func (rt *_router) createGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check authorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid user ID")
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		Forbidden(w, err, ctx, "Forbidden")
		return
	}
	// var
	var group Group
	type msg struct {
		Groupname    string   `json:"groupname"`
		Usernamelist []string `json:"usernamelist"`
	}
	var Msg msg
	// take data from the body request
	if err := json.NewDecoder(r.Body).Decode(&Msg); err != nil {
		BadRequest(w, err, ctx, "Invalid request body")
		return
	}
	group.Name = Msg.Groupname
	// check if name of group is valid
	if !group.IsValid() {
		BadRequest(w, nil, ctx, "Invalid group name")
		return
	}

	// set name
	group, err = rt.CreateGroupDB(group, userID)
	if err != nil {
		InternalServerError(w, err, "Unable to create group", ctx)
		return
	}
	w.WriteHeader(http.StatusCreated)

	// list of users
	userList := Msg.Usernamelist
	for i := 0; i < len(userList); i++ {
		// check if user exyst by username
		exist, err := rt.db.ExistName(userList[i])
		if err != nil {
			InternalServerError(w, err, "can't check if user exist", ctx)
			return
		}
		if !exist {
			BadRequest(w, nil, ctx, "INVALID USERNAME")
			return

		}
		// get user by username
		dbUser, err := rt.db.GetUserByName(userList[i])
		if err != nil {
			InternalServerError(w, err, "can't get user by username", ctx)
			return

		}
		// add user to group
		err = rt.db.AddUserToGroup(group.GroupID, dbUser.UserID)
		if err != nil {
			InternalServerError(w, err, "can't add user to group", ctx)
			return

		}
	}
	// create conversation of group
	var c Conversation
	c.GroupID = group.GroupID
	type response struct {
		Conversation Conversation `json:"conversation"`
		GroupUsers   []User       `json:"groupUsers"`
		Group        Group        `json:"group"`
	}
	// create conversation
	c, err = rt.CreateConversationDB(c)
	if err != nil {
		InternalServerError(w, err, "can't create conversation", ctx)
		return
	}
	var res response
	res.Conversation = c
	res.Group = group
	// get all users in group
	groupUsers, err := rt.db.GetUsersInGroup(group.GroupID)
	if err != nil {
		InternalServerError(w, err, "can't get users in group", ctx)
		return
	}
	var groupUsersResponse []User
	for i := 0; i < len(groupUsers); i++ {
		userD, err := rt.db.GetUserByID(groupUsers[i])
		if err != nil {
			InternalServerError(w, err, "can't get user by id", ctx)
			return
		}
		var user User
		err = user.FromDatabase(userD)
		if err != nil {
			InternalServerError(w, err, "can't convert user from database", ctx)
			return
		}
		groupUsersResponse = append(groupUsersResponse, user)
	}
	res.GroupUsers = groupUsersResponse

	// response
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		InternalServerError(w, err, "can't encode response", ctx)
		return
	}

}
