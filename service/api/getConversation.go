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
	exist, err := rt.db.ExistConversationByID(conversationID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}
	if !exist {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var membergroup []User
	// get conversation to check if groupId !=0
	conversation, err := rt.db.GetConversation(conversationID)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}
	if conversation.GroupID != 0 {
		allusers, err := rt.db.GetMemberGroup(conversation.GroupID)

		if err != nil {
			InternalServerError(w, err, "can't get the list of users", ctx)
			return
		}
		for _, u := range allusers {
			usDB, err := rt.db.GetUserByID(u)
			if err != nil {
				InternalServerError(w, err, "can't get the user", ctx)
				return
			}
			var us User
			err = us.FromDatabase(usDB)
			if err != nil {
				InternalServerError(w, err, "can't get the user", ctx)
				return
			}

			membergroup = append(membergroup, us)

		}

	}

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

	type MessageListResponse struct {
		Messages   []MessageResponse `json:"messages"`
		MemberList []User            `json:"memberlist"`
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

	finalResponse := MessageListResponse{
		Messages:   response,
		MemberList: membergroup,
	}
	// return
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(finalResponse)
	if err != nil {
		InternalServerError(w, err, "Internal Server Error", ctx)
		return
	}

}
