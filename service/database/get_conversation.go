package database

var query_GETCONVERSATION = `SELECT groupID FROM Conversation WHERE conversationID = ?;`

var query_GETCONVERSATIONPRIVATE = `SELECT conversationID FROM MemberPrivate WHERE conversationID = ? AND  userID = ?;`

func (db *appdbimpl) GetConversation(conversationID int) (Conversation, error) {
	// get conversation from database
	var conversation Conversation
	conversation.ConversationID = conversationID

	err := db.c.QueryRow(query_GETCONVERSATION, conversationID).Scan(&conversation.GroupID)
	if err != nil {
		return conversation, err
	}

	return conversation, nil

}

func (db *appdbimpl) GetConversationPrivate(conversationID int, userID int) (Conversation, error) {
	// get conversation from database
	var conversation Conversation

	err := db.c.QueryRow(query_GETCONVERSATIONPRIVATE, conversationID, userID).Scan(&conversation.ConversationID)
	if err != nil {
		return conversation, err
	}

	return conversation, nil

}

func (db *appdbimpl) GetConversationIDfrom2Users(userID1 int, userID2 int) (Conversation, error) {
	// get conversation from database
	var conversation Conversation

	err := db.c.QueryRow(`SELECT c1.conversation_id
		FROM Conversation c1
		JOIN Conversation c2 ON c1.conversation_id = c2.conversation_id
		WHERE c1.user_id = ?
		  AND c2.user_id = ?;`, userID1, userID2).Scan(&conversation.ConversationID)

	if err != nil {
		return conversation, err
	}

	return conversation, nil

}
