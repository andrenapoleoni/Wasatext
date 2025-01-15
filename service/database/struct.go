package database

// user struct
type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

// group struct
type Group struct {
	GroupID int    `json:"groupID"`
	Name    string `json:"name"`
}

// conversation struct
type Conversation struct {
	ConversationID int `json:"conversationID"`
	GroupID        int `json:"groupID"`
}

// message struct
type Message struct {
	MessageID      int    `json:"messageID"`
	ConversationID int    `json:"conversationID"`
	UserID         int    `json:"userID"`
	MessageTXT     string `json:"message"`
}

// membergroup struct
type MemberGroup struct {
	GroupID int `json:"groupID"`
	UserID  int `json:"userID"`
}

// comment struct
type Comment struct {
	CommentID      int    `json:"commentID"`
	UserID         int    `json:"userID"`
	ConversationID int    `json:"conversationID"`
	MessageID      int    `json:"messageID"`
	CommentTXT     string `json:"comment"`
}
