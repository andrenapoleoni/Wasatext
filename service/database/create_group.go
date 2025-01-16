package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"myWasatext/service/api/utils"
	"os"
)

var query_ADDGROUP = ` INSERT INTO Groupchat (groupID,groupName)
								VALUES(?,?);`

var query_MAXIDGROUP = `SELECT MAX(groupID) FROM Groupchat`

func (db *appdbimpl) CreateGroup(g Group, userID int) (Group, error) {
	var group Group
	group.Name = g.Name

	var _maxIDG = sql.NullInt64{Int64: 0, Valid: false}

	row, err := db.c.Query(query_MAXIDGROUP)
	if err != nil {
		return group, err
	}

	var maxIDG int

	for row.Next() {
		if row.Err() != nil {
			return group, err
		}

		err = row.Scan(&_maxIDG)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return group, err
		}

		if !_maxIDG.Valid {
			maxIDG = 0
		} else {
			maxIDG = int(_maxIDG.Int64)
		}
	}
	// set new group id
	group.GroupID = maxIDG + 1

	// ---------create group folder-------- //
	path := "./storage/group/" + fmt.Sprint(group.GroupID) + "/gallery"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return group, err
	}
	// --------set default----------- //
	source, err := os.Open("./storage/default/defaultphoto.jpg")
	if err != nil {
		return group, err
	}
	defer source.Close()

	destination, err := os.Create((utils.GetGroupPath(group.GroupID)))
	if err != nil {
		return group, err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return group, err
	}

	_, err = db.c.Exec(query_ADDGROUP, group.GroupID, group.Name)
	if err != nil {
		return group, err
	}

	err = db.AddUserToGroup(group.GroupID, userID)
	if err != nil {
		return group, err
	}

	return group, nil

}
