package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"myWasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) CreateMessageDB(message Message) (Message, error) {
	// create message in database
	messageDB, err := rt.db.CreateMessage(message.ToDatabase())
	if err != nil {
		return message, err
	}

	// convert message from database
	err = message.FromDatabase(messageDB)
	if err != nil {
		return message, err
	}

	return message, nil

}

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	// get conversation id
	conversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	// chech if conversation exist
	conversation, err := rt.db.GetConversation(conversationID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}
	// check if conversation is private
	if conversation.GroupID == 0 {
		conversation, err = rt.db.GetConversationPrivate(conversationID, userID)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
	}

	// get messagetxt from body
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}

	message.UserID = userID
	message.ConversationID = conversationID

	// send message
	message, err = rt.CreateMessageDB(message)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

}
