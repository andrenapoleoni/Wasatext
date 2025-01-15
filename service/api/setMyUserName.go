package api

import (
	"encoding/json"
	"myWasatext/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/* description
.
.
.
*/

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	profileUserID, err := strconv.Atoi(ps.ByName("user"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	//check authorization
	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	//take data from the body request
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	//check if username is valid
	if !user.IsValid() {
		http.Error(w, "INVALID USERNAME", http.StatusBadRequest)
		return
	}

	//change username, if username already taken return an error
	if err := rt.db.ChangeUsername(userID, user.Username); err != nil {
		http.Error(w, "username already taken, please retry", http.StatusBadRequest)
		return
	}

	//response success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "plain/text")
	if err := json.NewEncoder(w).Encode("Username changed succesfully"); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
