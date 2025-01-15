package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check authorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	//check authorization
	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	//get receiverID from endpoint

	//get conversationID from endpoint
	conversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	//check if conversation exists
	conversation, err := rt.db.GetConversation(conversationID)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	//return conversation
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conversation)

}
