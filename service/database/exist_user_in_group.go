package database

import (
	"database/sql"
	"errors"
)

var query_MEMBERINGROUP = `SELECT groupID, userID FROM MemberGroup
							WHERE groupID = ? AND userID = ?`

func (db *appdbimpl) ExistUserInGroup(groupID int, userID int) (bool, error) {
	var _groupID int
	var _userID int
	err := db.c.QueryRow(query_MEMBERINGROUP, groupID, userID).Scan(&_groupID, &_userID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return _groupID == groupID && _userID == userID, nil
}
