package database

import (
	"database/sql"
	"errors"
)

var query_EXISTMESSAGE = `SELECT messageID FROM Message WHERE messageID = ? AND conversationID = ?`

func (db *appdbimpl) ExistMessage(messageID int, conversationID int) (bool, error) {
	var existmessageID int

	err := db.c.QueryRow(query_EXISTMESSAGE, messageID, conversationID).Scan(&existmessageID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existmessageID != 0, err
}
