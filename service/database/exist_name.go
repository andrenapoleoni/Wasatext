package database

import (
	"database/sql"
	"errors"
)

var query_FINDUSERNAME = `SELECT username FROM User
						WHERE username = ?`

var query_FINDUSERID = `SELECT userID FROM User WHERE userID = ?`

func (db *appdbimpl) ExistName(username string) (bool, error) {
	var existName string
	err := db.c.QueryRow(query_FINDUSERNAME, username).Scan(&existName)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existName != "", err
}

func (db *appdbimpl) ExistUserID(userID int) (bool, error) {
	var existUserID int
	err := db.c.QueryRow(query_FINDUSERID, userID).Scan(&existUserID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existUserID != 0, err
}
