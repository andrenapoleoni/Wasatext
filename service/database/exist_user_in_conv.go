package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) ExistUserInConv(userID int, conversationID int) (bool, error) {
	var existUser int
	err := db.c.QueryRow(`SELECT userID FROM MemberPrivate WHERE conversationID = ? AND userID = ?;`, conversationID, userID).Scan(&existUser)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existUser != 0, err
}
