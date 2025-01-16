package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check authorization
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

	// take chat id from endpoint
	ConversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid conversation id")
		return
	}
	// check if conversation exists

	conversation, err := rt.db.GetConversation(ConversationID)
	if err != nil {
		http.Error(w, "Not Found "+err.Error(), http.StatusNotFound)
		return
	}

	// check if conversation is a group or private chat
	if conversation.GroupID == 0 {
		// private chat
		conversation, err = rt.db.GetConversationPrivate(ConversationID, userID)
		if err != nil {
			http.Error(w, "Not Found "+err.Error(), http.StatusNotFound)
			return
		}

	} else {
		// check if user is a member of group
		exist, err := rt.db.ExistUserInGroup(conversation.GroupID, userID)
		if err != nil {
			Forbidden(w, err, ctx, "Forbidden")
			return
		}
		if !exist {
			Forbidden(w, err, ctx, "Forbidden")
			return
		}

	}

	// take message id from endpoint
	MessageID, err := strconv.Atoi(ps.ByName("message"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid message id")
		return
	}
	// check if message exists
	message, err := rt.db.GetMessage(conversation.ConversationID, MessageID)
	if err != nil {
		http.Error(w, "Not Found "+err.Error(), http.StatusNotFound)
		return
	}
	message.UserID = userID
	// take dest id from query
	dest := r.URL.Query().Get("dest")
	if dest == "" {
		BadRequest(w, err, ctx, "Invalid dest id")
		return
	}

	// check if conversation exists by id from dest id
	destID, err := strconv.Atoi(dest)

	if err != nil {
		BadRequest(w, err, ctx, "Invalid dest id")
		return
	}

	conversation, err = rt.db.GetConversation(destID)
	if err != nil {
		InternalServerError(w, err, "Failed to get conversation", ctx)
		return
	}

	// create message

	message.ConversationID = destID
	message, err = rt.db.CreateMessage(message)
	if err != nil {
		InternalServerError(w, err, "Failed to create message", ctx)
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		InternalServerError(w, err, "Failed to encode response", ctx)
		return
	}

}
