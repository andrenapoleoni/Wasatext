package database

import (
	"database/sql"
	"errors"
)

var query_EXISTCOMMENT = `SELECT commentID FROM Comment WHERE commentID = ? AND messageID = ?`

func (db *appdbimpl) ExistComment(commentID int, messageID int) (bool, error) {
	var existcommentID int

	err := db.c.QueryRow(query_EXISTCOMMENT, commentID, messageID).Scan(&existcommentID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existcommentID != 0, err
}
