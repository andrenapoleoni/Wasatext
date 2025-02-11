package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// Function convert an image in base64
func ImageToBase64(filename string) (string, error) {
	imageFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { err = imageFile.Close() }()

	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	base64 := base64.StdEncoding.EncodeToString(imageData)
	return base64, err
}

// get path group photo
func GetGroupPhotoPath(groupID int) string {
	return fmt.Sprintf("./storage/group/%d/group_photo.jpg", groupID)
}

// get path user photo
func GetUserPhotoPath(userID int) string {
	return fmt.Sprintf("./storage/profiles/%d/user_photo.jpg", userID)
}
