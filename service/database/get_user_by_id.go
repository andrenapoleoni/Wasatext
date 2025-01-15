package database

var query_GETUSERBYID = `SELECT username FROM User WHERE userID = ?;`

func (db *appdbimpl) GetUserByID(userID int) (User, error) {
	//get user from database
	var user User
	user.UserID = userID

	err := db.c.QueryRow(query_GETUSERBYID, userID).Scan(&user.Username)
	if err != nil {
		return user, err
	}

	return user, nil

}
