package api

import (
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) CreateCommentDB(comment Comment) (Comment, error) {
	dbcomment, err := rt.db.CreateComment(comment.ToDatabase())
	if err != nil {
		return comment, err
	}

	err = comment.FromDatabase(dbcomment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check authorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		Forbidden(w, err, ctx, "Forbidden")
		return
	}
	// get conversation id from endopoint
	conversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	// check if conversation exists
	conversation, err := rt.db.GetConversation(conversationID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

	// get message id from endpoint
	messageID, err := strconv.Atoi(ps.ByName("message"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	// check if message exists
	exist, err := rt.db.ExistMessage(messageID, conversation.ConversationID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}
	if !exist {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// take comment from body
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	comment.UserID = userID
	comment.MessageID = messageID
	comment.ConversationID = conversationID

	// check if comment is valid
	if !comment.IsValid() {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	// check if user is in conversation
	// private conversation
	if conversation.GroupID == 0 {
		exist, err := rt.db.ExistUserInConv(userID, conversation.ConversationID)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
		if !exist {
			Forbidden(w, err, ctx, "Forbidden")
			return
		}
	} else {
		// group conversation
		exist, err := rt.db.ExistUserInGroup(conversation.GroupID, userID)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
		if !exist {
			Forbidden(w, err, ctx, "Forbidden")
			return
		}
	}

	// check if user already commented
	commentID, err := rt.db.GetExistComment(messageID, conversationID, userID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

	// if exist update comment
	if commentID != 0 {
		err = rt.db.UpdateComment(commentID, comment.CommentTXT)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}

	} else {
		// else create comment
		comment, err = rt.CreateCommentDB(comment)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
	}

	comment.CommentID = commentID

	// response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

}
