package database

var query_GETLISTCONVERSATIONSBYUSER = "SELECT * FROM MemberPrivate WHERE userID=?"

var query_GETLISTCONVERSATIONSBYGROUP = "SELECT * FROM conversation WHERE groupID = ?"

var query_TAKEGROUP = "SELECT * FROM membergroup WHERE userID = ? "

func (db *appdbimpl) GetListConversations(userID int) ([]Conversation, error) {
	var conversations []Conversation
	//get conversation of group where user is member
	rows, err := db.c.Query(query_TAKEGROUP, userID)
	if err != nil {
		return conversations, err
	}
	defer rows.Close()

	for rows.Next() {
		var groupID int
		err = rows.Scan(&groupID)
		if err != nil {
			return conversations, err
		}

		lines, err := db.c.Query(query_GETLISTCONVERSATIONSBYGROUP, groupID)
		if err != nil {
			return conversations, err
		}
		defer lines.Close()

		for lines.Next() {
			var conversation Conversation

			err = lines.Scan(&conversation.ConversationID, &conversation.GroupID)
			if err != nil {
				return conversations, err
			}

			conversations = append(conversations, conversation)
		}
	}
	//get conversation of private chat where user is member
	rows, err = db.c.Query(query_GETLISTCONVERSATIONSBYUSER, userID)
	if err != nil {
		return conversations, err
	}
	defer rows.Close()

	for rows.Next() {
		var conversation Conversation
		err = rows.Scan(&conversation.ConversationID, &conversation.GroupID)
		if err != nil {
			return conversations, err
		}

		conversations = append(conversations, conversation)
	}

	return conversations, nil

}
