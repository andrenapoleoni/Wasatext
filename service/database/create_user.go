package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"

	"myWasatext/service/api/utils"
)

var query_ADDUSER = ` INSERT INTO User (userID,username)
								VALUES(?,?);`

var query_MAXID = `SELECT MAX(userID) FROM User`

func (db *appdbimpl) CreateUser(u User) (User, error) {
	var user User
	user.Username = u.Username

	var _maxID = sql.NullInt64{Int64: 0, Valid: false}

	row, err := db.c.Query(query_MAXID)

	if err != nil {
		return user, err
	}

	var maxID int

	for row.Next() {
		if row.Err() != nil {
			return user, err
		}

		err = row.Scan(&_maxID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return user, err
		}

		if !_maxID.Valid {
			maxID = 0
		} else {
			maxID = int(_maxID.Int64)
		}
	}

	//--------set new userid----------//
	user.UserID = maxID + 1

	//---------create user folder--------//
	path := "./storage/profiles/" + fmt.Sprint(user.UserID) + "/gallery"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return user, err
	}
	//--------set default-----------//
	source, err := os.Open("./storage/default/defaultphoto.jpg")
	if err != nil {
		return user, err
	}
	defer source.Close()

	destination, err := os.Create((utils.GetProfilePath(user.UserID)))
	if err != nil {
		return user, err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return user, err
	}

	_, err = db.c.Exec(query_ADDUSER, user.UserID, user.Username)
	if err != nil {
		return user, err
	}

	return user, nil
}
