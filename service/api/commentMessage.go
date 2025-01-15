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
	//get conversation id from endopoint
	conversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	//check if conversation exists
	conversation, err := rt.db.GetConversation(conversationID)
	if err != nil {
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	//get message id from endpoint
	messageID, err := strconv.Atoi(ps.ByName("message"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	//check if message exists
	exist, err := rt.db.ExistMessage(messageID, conversation.ConversationID)
	if err != nil {
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !exist {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	//take comment from body
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	comment.UserID = userID
	comment.MessageID = messageID
	comment.ConversationID = conversationID

	//check if comment is valid
	if !comment.IsValid() {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//create comment
	comment, err = rt.CreateCommentDB(comment)
	if err != nil {
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response for group creation")
		http.Error(w, "Internal Server Error: unable to encode response", http.StatusInternalServerError)
		return
	}

}
