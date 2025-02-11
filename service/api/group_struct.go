package api

import (
	"myWasatext/service/api/utils"
	"myWasatext/service/database"
	"regexp"
)

type Group struct {
	GroupID int    `json:"groupID"`
	Name    string `json:"groupname"`
	Photo   string `json:"photo"`
}

func (g *Group) IsValid() bool {
	groupname := g.Name
	validName := regexp.MustCompile(`^.*?$`)
	return validName.MatchString(groupname)
}

func (g *Group) ToDatabase() database.Group {
	return database.Group{
		GroupID: g.GroupID,
		Name:    g.Name,
	}
}

func (g *Group) FromDatabase(dbGroup database.Group) error {
	g.GroupID = dbGroup.GroupID
	g.Name = dbGroup.Name
	profilephoto, err := utils.ImageToBase64(utils.GetGroupPhotoPath(dbGroup.GroupID))
	if err != nil {
		return err
	}
	g.Photo = profilephoto

	return nil
}
