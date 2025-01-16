package database

var query_GETMESSAGE = `SELECT message, userID FROM Message WHERE messageID = ? AND conversationID = ?;`

func (db *appdbimpl) GetMessage(conversationID int, messageID int) (Message, error) {
	// get message from database
	var message Message
	message.MessageID = messageID
	message.ConversationID = conversationID

	err := db.c.QueryRow(query_GETMESSAGE, messageID, conversationID).Scan(&message.MessageTXT)
	if err != nil {
		return message, err
	}

	return message, nil

}
