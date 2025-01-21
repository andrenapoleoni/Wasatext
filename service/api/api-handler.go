package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// User routes
	//	--------LOGIN AND REGISTER---------	//
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))
	//	-----SET USERNAME-----	//
	rt.router.PUT("/user/:user/username", rt.wrap(rt.setMyUserName, true))
	//	--------DELETE USER------	//
	rt.router.DELETE("/user/:user", rt.wrap(rt.DeleteUser, true))
	//	--------SEARCH USER------	//
	rt.router.GET("/user", rt.wrap(rt.searchUser, true))
	//	--------SET PHOTO------	//
	rt.router.PUT("/user/:user/photo", rt.wrap(rt.SetMyPhoto, true))

	// Group routes
	// --------CREATE GROUP------ //
	rt.router.POST("/user/:user/groups", rt.wrap(rt.createGroup, true))
	//  --------ADD TO GROUP------  //
	rt.router.POST("/user/:user/groups/:group", rt.wrap(rt.addToGroup, true))
	// --------LEAVE GROUP------ //
	rt.router.DELETE("/user/:user/groups/:group", rt.wrap(rt.leaveGroup, true))
	// --------SET GROUPNAME------ //
	rt.router.PUT("/user/:user/groups/:group/groupname", rt.wrap(rt.setGroupName, true))
	// ---------SET GROUP PHOTO--------- //
	rt.router.PUT("/user/:user/groups/:group/groupphoto", rt.wrap(rt.setGroupPhoto, true))

	// Conversation routes
	// --------CREATE CONVERSATION------ //
	rt.router.PUT("/user/:user/conversation/:conversation", rt.wrap(rt.createConversation, true))
	// --------SEND MESSAGE------ //
	rt.router.POST("/user/:user/conversation/:conversation/messages", rt.wrap(rt.sendMessage, true))
	// --------GET CONVERSATION------ //
	rt.router.GET("/user/:user/conversation/:conversation", rt.wrap(rt.getConversation, true))
	// --------GET CONVERSATIONS------ //
	rt.router.GET("/user/:user/conversation", rt.wrap(rt.getMyConversations, true))

	// Message routes
	// --------FORWARD MESSAGE------ //
	rt.router.POST("/user/:user/conversation/:conversation/messages/:message", rt.wrap(rt.forwardMessage, true))
	// ------COMMENT MESSAGE------ //
	rt.router.PUT("/user/:user/conversation/:conversation/messages/:message/comments", rt.wrap(rt.commentMessage, true))
	// ------DELETE MESSAGE------ //
	rt.router.DELETE("/user/:user/conversation/:conversation/messages/:message", rt.wrap(rt.deleteMessage, true))
	// ------DELETE COMMENT------ //
	rt.router.DELETE("/user/:user/conversation/:conversation/messages/:message/comments/:comment", rt.wrap(rt.uncommentMessage, true))

	// special routes
	rt.router.GET("/liveness", rt.liveness)

	// return router
	return rt.router
}
