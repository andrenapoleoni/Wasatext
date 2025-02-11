package api

import (
	"io"
	"net/http"
	"strconv"

	"os"

	"encoding/base64"
	"encoding/json"

	"myWasatext/service/api/utils"

	"myWasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SetMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Check if the size of the image is less than 5MB
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

	response := base64.StdEncoding.EncodeToString(data)
	// Check if the file is a jpeg
	fileType := http.DetectContentType(data)
	if fileType != "image/jpeg" {
		BadRequest(w, err, ctx, "Invalid file type")
		return
	}
	defer func() { err = file.Close() }()

	// save photo
	path := utils.GetProfilePath(userID)
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

	type Response struct {
		Photo string `json:"photo"`
	}

	var res Response
	res.Photo = response

	// response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		InternalServerError(w, err, "Error encoding the response", ctx)
		return
	}

}
