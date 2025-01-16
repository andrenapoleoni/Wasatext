package database

import (
	"database/sql"
	"errors"
)

var query_FINDGROUP = `SELECT groupID FROM Groupchat WHERE groupID = ?`

func (db *appdbimpl) ExistGroup(groupID int) (bool, error) {
	var existGroup int
	err := db.c.QueryRow(query_FINDGROUP, groupID).Scan(&existGroup)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existGroup != 0, err
}
