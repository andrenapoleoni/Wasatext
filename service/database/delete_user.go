package database

var query_DELETEUSER = `DELETE FROM User WHERE userID = ?`

func (db *appdbimpl) DeleteUser(userID int) error {
	_, err := db.c.Exec(query_DELETEUSER, userID)
	return err
}
