package database

import (
	"database/sql"
	"errors"
)

var query_Exist = `SELECT commentID FROM Comment WHERE (messageID,conversationID,userID)=(?,?,?) ;`

func (db *appdbimpl) GetExistComment(messageID int, conversationID int, userID int) (int, error) {
	var commentID int
	err := db.c.QueryRow(query_Exist, messageID, conversationID, userID).Scan(&commentID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	return commentID, err
}
