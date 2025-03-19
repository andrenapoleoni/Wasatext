package database

var query_GETMESSAGE = `SELECT message, userID FROM Message WHERE messageID = ? AND conversationID = ?;`

func (db *appdbimpl) GetMessage(conversationID int, messageID int) (Message, error) {
	// get message from database
	var message Message
	message.MessageID = messageID
	message.ConversationID = conversationID

	err := db.c.QueryRow(query_GETMESSAGE, messageID, conversationID).Scan(&message.MessageTXT, &message.UserID)
	if err != nil {
		return message, err
	}

	return message, nil

}

func (db *appdbimpl) GetAllMessage(conversationID int) ([]Message, error) {
	// get all message from database
	var messages []Message
	rows, err := db.c.Query("SELECT messageID, message, userID FROM Message WHERE conversationID = ?;", conversationID)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var message Message
		err := rows.Scan(&message.MessageID, &message.MessageTXT, &message.UserID)
		if err != nil {
			return messages, err
		}
		if rows.Err() != nil {
			return messages, err
		}
		message.ConversationID = conversationID
		messages = append(messages, message)
	}

	return messages, nil

}
