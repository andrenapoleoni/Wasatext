package api

import (
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check authentication
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
	//get groupID from request
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

	//check if user is a member of the group
	_, err = rt.db.ExistUserInGroup(groupId, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get the group")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//delete user from group
	err = rt.db.DeleteUserFromGroup(groupId, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't delete user from group")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//check if user is the last member of the group
	member, err := rt.db.GetMemberGroup(groupId)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get the group members")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//delete group if user is the last member
	if len(member) == 0 {
		err = rt.db.DeleteGroup(groupId)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't delete group")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//response
	w.WriteHeader(http.StatusOK)

}
