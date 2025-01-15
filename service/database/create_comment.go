package database

import (
	"database/sql"
	"errors"
)

var sql_CREATECOMMENT = ` INSERT INTO Comment (commentID,messageID,commentTXt,conversationID,userID) VALUES (?, ?, ?, ?,?)`

var sql_MAXIDCOMMENT = `SELECT MAX(commentID) FROM Comment`

func (db *appdbimpl) CreateComment(comment Comment) (Comment, error) {
	var c Comment
	c.ConversationID = comment.ConversationID
	c.MessageID = comment.MessageID
	c.UserID = comment.UserID
	c.CommentTXT = comment.CommentTXT
	//get last id from message
	var _maxID = sql.NullInt64{Int64: 0, Valid: false}

	row, err := db.c.Query(sql_MAXIDCOMMENT)

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

	//--------set new messageid----------//
	c.CommentID = maxID + 1

	//add message to database
	_, err = db.c.Exec(sql_CREATECOMMENT, c.CommentID, c.MessageID, c.CommentTXT, c.ConversationID, c.UserID)
	if err != nil {
		return c, err
	}

	return c, nil
}
