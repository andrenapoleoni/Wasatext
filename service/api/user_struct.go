package api

import (
	"myWasatext/service/database"
	"regexp"
)

type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

func (u *User) IsValid() bool {
	username := u.Username
	validUser := regexp.MustCompile(`^.*?$`)
	return validUser.MatchString(username)
}

func (u *User) ToDatabase() database.User {
	return database.User{
		UserID:   u.UserID,
		Username: u.Username,
	}
}

func (u *User) FromDatabase(dbUser database.User) error {
	u.UserID = dbUser.UserID
	u.Username = dbUser.Username

	return nil
}
