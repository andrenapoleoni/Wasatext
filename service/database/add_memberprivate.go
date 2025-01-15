package database

var sql_ADDMEMBERPRIVATE = `INSERT INTO MemberPrivate (conversationID, userID) VALUES (?, ?)`

func (db *appdbimpl) AddMemberPrivate(conversationID int, userID int) error {
	_, err := db.c.Exec(sql_ADDMEMBERPRIVATE, conversationID, userID)
	return err
}
