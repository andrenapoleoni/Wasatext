package api

import (
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check authentication
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid user id")
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		Forbidden(w, err, ctx, "Forbidden")
		return
	}
	// get groupID from request
	groupId, err := strconv.Atoi(ps.ByName("group"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid group id")
		return
	}
	// check if group exist
	exist, err := rt.db.ExistGroup(groupId)
	if err != nil {
		InternalServerError(w, err, "Failed to get group", ctx)
		return
	}

	if !exist {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	// check if user is a member of the group
	_, err = rt.db.ExistUserInGroup(groupId, userID)
	if err != nil {
		InternalServerError(w, err, "Failed to check user existence in group", ctx)
		return
	}

	// delete user from group
	err = rt.db.DeleteUserFromGroup(groupId, userID)
	if err != nil {
		InternalServerError(w, err, "Failed to delete user from group", ctx)
		return
	}
	// check if user is the last member of the group
	member, err := rt.db.GetMemberGroup(groupId)
	if err != nil {
		InternalServerError(w, err, "Failed to get member group", ctx)
		return
	}
	// delete group if user is the last member
	if len(member) == 0 {
		err = rt.db.DeleteGroup(groupId)
		if err != nil {
			InternalServerError(w, err, "Failed to delete group", ctx)
			return
		}
	}
	// response
	w.WriteHeader(http.StatusOK)

}
