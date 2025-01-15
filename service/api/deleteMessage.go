package api

import (
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	//take conversation id from endpoint
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

	//take messageID from endpoint
	messageID, err := strconv.Atoi(ps.ByName("message"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	//check if message exist
	exist, err := rt.db.ExistMessage(messageID, conversation.ConversationID)
	if err != nil {
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !exist {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	//delete message
	err = rt.db.DeleteMessage(messageID, conversation.ConversationID)
	if err != nil {
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusOK)

}
