package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// User routes
	//--------LOGIN AND REGISTER---------//
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))
	//-----SET USERNAME-----//
	rt.router.PUT("/profiles/:user/username", rt.wrap(rt.setMyUserName, true))
	//--------DELETE USER------//
	rt.router.DELETE("/profiles/:user", rt.wrap(rt.DeleteUser, true))
	//--------SEARCH USER------//
	rt.router.GET("/profiles", rt.wrap(rt.SearchUser, true))

	//Group routes
	//--------CREATE GROUP------//
	rt.router.POST("/profiles/:user/groups", rt.wrap(rt.CreateaGroup, true))
	//--------ADD TO GROUP------//
	rt.router.POST("/profiles/:user/groups/:group", rt.wrap(rt.addToGroup, true))
	//--------LEAVE GROUP------//
	rt.router.DELETE("/profiles/:user/groups/:group", rt.wrap(rt.leaveGroup, true))
	//--------SET GROUPNAME------//
	rt.router.PUT("/profiles/:user/groups/:group/groupname", rt.wrap(rt.setGroupName, true))

	//Conversation routes
	//--------CREATE CONVERSATION------//
	rt.router.PUT("/profiles/:user/conversations/:conversation", rt.wrap(rt.CreateConversation, true))
	//--------SEND MESSAGE------//
	rt.router.POST("/profiles/:user/conversations/:conversation/messages", rt.wrap(rt.sendMessage, true))
	//--------GET CONVERSATION------//
	rt.router.GET("/profiles/:user/conversations/:conversation", rt.wrap(rt.getConversation, true))
	//--------GET CONVERSATIONS------//
	rt.router.GET("/profiles/:user/conversations", rt.wrap(rt.getMyConversations, true))
	//Message routes
	//--------FORWARD MESSAGE------//
	rt.router.POST("/profiles/:user/conversations/:conversation/messages/:message", rt.wrap(rt.forwardMessage, true))
	//------COMMENT MESSAGE------//
	rt.router.PUT("/profiles/:user/conversations/:conversation/messages/:message/comments", rt.wrap(rt.commentMessage, true))
	//------DELETE MESSAGE------//
	rt.router.DELETE("/profiles/:user/conversations/:conversation/messages/:message", rt.wrap(rt.deleteMessage, true))
	//------DELETE COMMENT------//
	rt.router.DELETE("/profiles/:user/conversations/:conversation/messages/:message/comments/:comment", rt.wrap(rt.uncommentMessage, true))

	//special routes
	rt.router.GET("/liveness", rt.liveness)

	//return router
	return rt.router
}
