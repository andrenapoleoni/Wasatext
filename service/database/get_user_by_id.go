package database

var query_GETUSERBYID = `SELECT username FROM User WHERE userID = ?;`

var query_GETUSERINCONVERSATIONPRIVATE = `SELECT userID FROM MemberPrivate WHERE conversationID = ? AND  userID != ?;`

func (db *appdbimpl) GetUserByID(userID int) (User, error) {
	// get user from database
	var user User
	user.UserID = userID

	err := db.c.QueryRow(query_GETUSERBYID, userID).Scan(&user.Username)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (db *appdbimpl) GetUserInConversationPrivate(conversationID int, userID int) (User, error) {

	var user User

	err := db.c.QueryRow(query_GETUSERINCONVERSATIONPRIVATE, conversationID, userID).Scan(&user.UserID)
	if err != nil {
		return user, err
	}
	user, err = db.GetUserByID(user.UserID)
	if err != nil {
		return user, err
	}

	return user, nil

}
