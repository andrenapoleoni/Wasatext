package api

import (
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) CreateConversationDB(c Conversation) (Conversation, error) {
	dbConversation, err := rt.db.CreateConversation(c.ToDatabase())
	if err != nil {
		return c, err
	}

	err = c.FromDatabase(dbConversation)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// chcek authorization
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
	// get users of conversation
	var u2 User

	u2.UserID, err = strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}
	// check if user exists
	exist, err := rt.db.ExistUserID(u2.UserID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error1", ctx)
		return
	}
	if !exist {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// check if conversation between users already exists
	exist, err = rt.db.ExistConversation(userID, u2.UserID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error2", ctx)
		return
	}
	if exist {
		http.Error(w, "Conversation already exists", http.StatusConflict)
		return
	}

	// create conversation
	var c Conversation

	c, err = rt.CreateConversationDB(c)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error3", ctx)
		return
	}
	// add user to member private table
	err = rt.db.AddMemberPrivate(c.ConversationID, userID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error4", ctx)
		return
	}
	err = rt.db.AddMemberPrivate(c.ConversationID, u2.UserID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error5", ctx)
		return
	}

	// get messagetxt from body
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}

	message.ConversationID = c.ConversationID
	message.UserID = userID
	// send message
	message, err = rt.CreateMessageDB(message)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error6", ctx)
		return
	}

	// response
	w.WriteHeader(http.StatusCreated)

}
