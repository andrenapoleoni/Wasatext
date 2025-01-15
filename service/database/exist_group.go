package database

import (
	"database/sql"
	"errors"
)

var query_FINDGROUP = `SELECT groupID, groupName FROM Groupchat WHERE groupID = ?`

func (db *appdbimpl) ExistGroup(groupID int) (bool, error) {
	var existGroup Group
	err := db.c.QueryRow(query_FINDGROUP, groupID).Scan(&existGroup.GroupID, &existGroup.Name)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existGroup.GroupID != 0, err
}
