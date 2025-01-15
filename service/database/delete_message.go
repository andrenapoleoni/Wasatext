package database

var query_DELETEMESSAGE = `DELETE FROM Message WHERE messageID = ? AND conversationID = ?`

func (db *appdbimpl) DeleteMessage(messageID int, conversationID int) error {
	_, err := db.c.Exec(query_DELETEMESSAGE, messageID, conversationID)
	return err
}
