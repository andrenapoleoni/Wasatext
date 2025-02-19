package api

import (
	"myWasatext/service/api/utils"
	"myWasatext/service/database"
	"regexp"
)

type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
}

func (u *User) IsValid() bool {
	username := u.Username
	validUser := regexp.MustCompile(`^.*?$`)
	// check if lenght of username is more than 3 and less than 16
	/*
		if len(username) < 3 || len(username) > 16 {
			return false
		}*/

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
	profilephoto, err := utils.ImageToBase64(utils.GetUserPhotoPath(dbUser.UserID))
	if err != nil {
		return err
	}
	u.Photo = profilephoto

	return nil
}
