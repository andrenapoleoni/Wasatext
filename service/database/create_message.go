package database

import (
	"database/sql"
	"errors"
)

var query_ADDMESSAGE = `INSERT INTO Message (conversationID, userID, messageID, message)
						VALUES (?, ?, ?, ?)`

var query_MESSAGEMAXID = `SELECT MAX(messageID) FROM Message WHERE conversationID = ?`

// CreateMessage creates a new message in the database
func (db *appdbimpl) CreateMessage(message Message) (Message, error) {
	var m Message
	m.ConversationID = message.ConversationID
	m.UserID = message.UserID
	m.MessageTXT = message.MessageTXT
	//get last id from message
	var _maxID = sql.NullInt64{Int64: 0, Valid: false}

	row, err := db.c.Query(query_MESSAGEMAXID, m.ConversationID)

	if err != nil {
		return m, err
	}

	var maxID int

	for row.Next() {
		if row.Err() != nil {
			return m, err
		}

		err = row.Scan(&_maxID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return m, err
		}

		if !_maxID.Valid {
			maxID = 0
		} else {
			maxID = int(_maxID.Int64)
		}
	}

	//--------set new messageid----------//
	m.MessageID = maxID + 1

	//add message to database
	_, err = db.c.Exec(query_ADDMESSAGE, m.ConversationID, m.UserID, m.MessageID, m.MessageTXT)
	if err != nil {
		return m, err
	}

	return m, nil
}
