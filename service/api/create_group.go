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

func (rt *_router) CreateaGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check authorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	//var
	var group Group
	type msg struct {
		Groupname    string   `json:"groupname"`
		Usernamelist []string `json:"usernamelist"`
	}
	var Msg msg
	//take data from the body request
	if err := json.NewDecoder(r.Body).Decode(&Msg); err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	group.Name = Msg.Groupname
	//check if name of group is valid
	if !group.IsValid() {
		http.Error(w, "INVALID NAME", http.StatusBadRequest)
		return
	}

	//set name
	group, err = rt.CreateGroupDB(group, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't create group")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	//list of users
	userList := Msg.Usernamelist
	for i := 0; i < len(userList); i++ {
		//check if user exyst by username
		exist, err := rt.db.ExistName(userList[i])
		if err != nil {
			ctx.Logger.WithError(err).Error("can't check the username")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !exist {
			http.Error(w, "INVALID USERNAME", http.StatusBadRequest)
			return

		}
		//get user by username
		dbUser, err := rt.db.GetUserByName(userList[i])
		if err != nil {
			ctx.Logger.WithError(err).Error("can't get the user")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		//add user to group
		err = rt.db.AddUserToGroup(group.GroupID, dbUser.UserID)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't add user to group")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
	}
	//create conversation of group
	var c Conversation
	c.GroupID = group.GroupID

	_, err = rt.CreateConversationDB(c)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't create conversation")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// response
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(group); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response for group creation")
		http.Error(w, "Internal Server Error: unable to encode response", http.StatusInternalServerError)
		return
	}

}
