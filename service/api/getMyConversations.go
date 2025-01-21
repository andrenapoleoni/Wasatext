package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check authorization
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
	// get conversations
	conversations, err := rt.db.GetListConversations(userID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}
	// convert from conv id to username or groupname
	type Response struct {
		Conversation Conversation `json:"conversation"`
		User         User         `json:"user"`
		Group        Group        `json:"group"`
		//Message      Message      `json:"message"`
	}
	// Response
	response := make([]Response, len(conversations))

	for i, conversationDB := range conversations {
		if conversationDB.GroupID != 0 {
			groupDB, err := rt.db.GetGroupByID(conversationDB.GroupID)
			if err != nil {
				InternalServerError(w, err, "Internal Server Errorbraerera", ctx)
				return
			}
			var group Group
			err = group.FromDatabase(groupDB)
			if err != nil {
				InternalServerError(w, err, "Internal Server Error", ctx)
				return
			}

			response[i].Group = group
		} else {
			userDB, err := rt.db.GetUserInConversationPrivate(conversationDB.ConversationID, userID)
			if err != nil {
				InternalServerError(w, err, "Internal Server Errorrere", ctx)
				return
			}
			var user User
			err = user.FromDatabase(userDB)
			if err != nil {
				InternalServerError(w, err, "Internal Server Error", ctx)
				return
			}

			response[i].User = user
		}
		var conversation Conversation
		err = conversation.FromDatabase(conversationDB)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}

		response[i].Conversation = conversation

		/*message, err := rt.db.GetLastMessage(conversation.ConversationID)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
		response[i].Message = message*/
	}

	// return conversations
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

}
