package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check authorization
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		http.Error(w, "Bad Request1 "+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	//take chat id from endpoint
	ConversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		http.Error(w, "Bad Request2 "+err.Error(), http.StatusBadRequest)
		return
	}
	//check if conversation exists

	conversation, err := rt.db.GetConversation(ConversationID)
	if err != nil {
		http.Error(w, "Not Found "+err.Error(), http.StatusNotFound)
		return
	}

	//check if conversation is a group or private chat
	if conversation.GroupID == 0 {
		//private chat
		conversation, err = rt.db.GetConversationPrivate(ConversationID, userID)
		if err != nil {
			http.Error(w, "Not Found "+err.Error(), http.StatusNotFound)
			return
		}

	} else {
		//check if user is a member of group
		exist, err := rt.db.ExistUserInGroup(conversation.GroupID, userID)
		if err != nil {
			http.Error(w, "Forbidden "+err.Error(), http.StatusForbidden)
			return
		}
		if !exist {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

	}

	//take message id from endpoint
	MessageID, err := strconv.Atoi(ps.ByName("message"))
	if err != nil {
		http.Error(w, "Bad Request3 "+err.Error(), http.StatusBadRequest)
		return
	}
	//check if message exists
	message, err := rt.db.GetMessage(conversation.ConversationID, MessageID)
	if err != nil {
		http.Error(w, "Not Found "+err.Error(), http.StatusNotFound)
		return
	}
	message.UserID = userID
	//take dest id from query
	dest := r.URL.Query().Get("dest")
	if dest == "" {
		http.Error(w, "Bad Request4", http.StatusBadRequest)
		return
	}

	//check if conversation exists by id from dest id
	destID, err := strconv.Atoi(dest)
	//destID := 3
	if err != nil {
		http.Error(w, "Bad Request5 "+err.Error(), http.StatusBadRequest)
		return
	}

	conversation, err = rt.db.GetConversation(destID)
	if err != nil {
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	//create message

	message.ConversationID = destID
	message, err = rt.db.CreateMessage(message)
	if err != nil {
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
