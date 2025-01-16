package api

import (
	"net/http"
	"strconv"

	"myWasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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
	// take conversation id from endpoint
	conversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid conversation id")
		return
	}
	// check if conversation exists
	conversation, err := rt.db.GetConversation(conversationID)
	if err != nil {
		InternalServerError(w, err, "Failed to get conversation", ctx)
		return
	}

	// take messageID from endpoint
	messageID, err := strconv.Atoi(ps.ByName("message"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid message id")
		return
	}
	// check if message exist
	exist, err := rt.db.ExistMessage(messageID, conversation.ConversationID)
	if err != nil {
		InternalServerError(w, err, "Failed to check message existence", ctx)
		return
	}
	if !exist {
		http.Error(w, "message Not Found", http.StatusNotFound)
		return
	}

	// take commentId from endpoint
	commentID, err := strconv.Atoi(ps.ByName("comment"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid comment id")
		return
	}
	// check if comment exists
	exist, err = rt.db.ExistComment(commentID, messageID)
	if err != nil {
		InternalServerError(w, err, "Failed to check comment existence", ctx)
		return
	}
	if !exist {
		http.Error(w, "comment Not Found", http.StatusNotFound)
		return
	}

	// delete comment
	err = rt.db.DeleteComment(commentID)
	if err != nil {
		InternalServerError(w, err, "Failed to delete comment", ctx)
		return
	}

	// response
	w.WriteHeader(http.StatusOK)

}
