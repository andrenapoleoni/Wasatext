package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// get conversationID from endpoint
	conversationID, err := strconv.Atoi(ps.ByName("conversation"))
	if err != nil {
		BadRequest(w, err, ctx, "Bad Request")
		return
	}

	// check if conversation exists

	// get messages of conversation
	messageDB, err := rt.db.GetAllMessage(conversationID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

	type MessageResponse struct {
		MessageR Message `json:"message"`
		UserR    User    `json:"user"`
	}

	response := make([]MessageResponse, len(messageDB))
	for i, message := range messageDB {
		userDB, err := rt.db.GetUserByID(message.UserID)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
		var user User
		err = user.FromDatabase(userDB)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
		var rsp MessageResponse
		var msg Message
		err = msg.FromDatabase(message)
		if err != nil {
			InternalServerError(w, err, "Internal Server Error", ctx)
			return
		}
		rsp.MessageR = msg
		rsp.UserR = user

		response[i] = rsp
	}
	// return
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

}
