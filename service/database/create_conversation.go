package database

import (
	"database/sql"
	"errors"
)

var query_ADDCONVERSATION = `INSERT INTO Conversation (conversationID, groupID)
								VALUES (?,  ?)`

var query_CONVERSATIONMAXID = `SELECT MAX(conversationID) FROM Conversation`

func (db *appdbimpl) CreateConversation(conversation Conversation) (Conversation, error) {

	var c Conversation
	c.GroupID = conversation.GroupID

	// get last id from conversation

	var _maxID = sql.NullInt64{Int64: 0, Valid: false}

	row, err := db.c.Query(query_CONVERSATIONMAXID)

	if err != nil {
		return c, err
	}

	var maxID int

	for row.Next() {
		if row.Err() != nil {
			return c, err
		}

		err = row.Scan(&_maxID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return c, err
		}

		if !_maxID.Valid {
			maxID = 0
		} else {
			maxID = int(_maxID.Int64)
		}
	}

	// --------set new conversationid---------- //
	c.ConversationID = maxID + 1
	// add conversation to database
	if c.GroupID == 0 {
		_, err = db.c.Exec(query_ADDCONVERSATION, c.ConversationID, c.GroupID)
		if err != nil {
			return c, err
		}
	} else {
		_, err = db.c.Exec(query_ADDCONVERSATION, c.ConversationID, c.GroupID)
		if err != nil {
			return c, err
		}

	}

	return c, nil
}
