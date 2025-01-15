package utils

import (
	"fmt"
)

func GetGroupPath(groupID int) string {

	return fmt.Sprintf("./storage/group/%d/group_photo.jpg", groupID)

}
