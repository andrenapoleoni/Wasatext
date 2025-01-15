package utils

import (
	"fmt"
)

func GetProfilePath(userID int) string {
	return fmt.Sprintf("./storage/profiles/%d/user_photo.jpg", userID)
}
