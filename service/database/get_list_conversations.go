package database

var query_GETGROUPIDWHEREUSERISMEMBER = `SELECT groupID FROM MemberGroup WHERE userID = ?;`
var query_GETCONVIDFORUSER = `SELECT conversationID FROM MemberPrivate WHERE userID = ?;`

func (db *appdbimpl) GetListConversations(userID int) ([]Conversation, error) {
	var conversations []Conversation
	//get conversation of user by member private table
	rows, err := db.c.Query(query_GETCONVIDFORUSER, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var conversation Conversation
		err := rows.Scan(&conversation.ConversationID)
		if err != nil {
			return nil, err
		}
		if rows.Err() != nil {
			return conversations, err
		}

		conversation, err = db.GetConversation(conversation.ConversationID)
		if err != nil {
			return nil, err
		}

		conversations = append(conversations, conversation)
	}

	// get conversation of user by groupmember table
	rows, err = db.c.Query(query_GETGROUPIDWHEREUSERISMEMBER, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var conversation Conversation

		err := rows.Scan(&conversation.GroupID)
		if err != nil {
			return nil, err
		}

		if rows.Err() != nil {
			return conversations, err
		}

		conversation, err = db.GetConversationIDfromGroup(conversation.GroupID)
		if err != nil {
			return nil, err
		}

		conversations = append(conversations, conversation)

	}
	return conversations, nil

}
