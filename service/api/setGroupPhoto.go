package api

import (
	"io"
	"myWasatext/service/api/utils"
	"net/http"
	"os"
	"strconv"

	"encoding/json"
	"myWasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check auth
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid user id")
		return
	}

	userID := ctx.UserID

	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// take group id from the request
	groupID, err := strconv.Atoi(ps.ByName("group"))
	if err != nil {
		BadRequest(w, err, ctx, "Invalid group id")
		return
	}

	// check if group exists
	exist, err := rt.db.ExistGroup(groupID)
	if err != nil {
		InternalServerError(w, err, "Error checking if group exists", ctx)
		return
	}
	if !exist {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	// check if the user is a member of the group
	isMember, err := rt.db.ExistUserInGroup(groupID, userID)
	if err != nil {
		InternalServerError(w, err, "Error checking if user is a member of the group", ctx)
		return
	}
	if !isMember {
		Forbidden(w, err, ctx, "User is not a member of the group")
		return
	}

	err = r.ParseMultipartForm(5 << 20)
	if err != nil {
		BadRequest(w, err, ctx, "Image too big")
		return
	}

	// Access the file from the request
	file, _, err := r.FormFile("image")
	if err != nil {
		BadRequest(w, err, ctx, "Bad request")
		return
	}

	// Read the file
	data, err := io.ReadAll(file) // In data we have the image file taked in the request
	if err != nil {
		InternalServerError(w, err, "Error reading the image file", ctx)
		return
	}
	// Check if the file is a jpeg
	fileType := http.DetectContentType(data)
	if fileType != "image/jpeg" {
		BadRequest(w, err, ctx, "Invalid file type")
		return
	}
	defer func() { err = file.Close() }()

	// save photo
	path := utils.GetGroupPath(groupID)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		InternalServerError(w, err, "Error saving the image file", ctx)
		return
	}

	// Crop the image
	err = utils.SaveAndCrop(path, 250, 250)
	if err != nil {
		InternalServerError(w, err, "Error cropping the image", ctx)
		return
	}

	//response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode("Photo changed"); err != nil {
		InternalServerError(w, err, "Error encoding the response", ctx)
		return
	}

}
