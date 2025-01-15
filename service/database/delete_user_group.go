package database

var query_DELEUSERTOGROUP = `DELETE FROM MemberGroup WHERE groupID = ? AND userID = ?`

func (db *appdbimpl) DeleteUserFromGroup(groupID int, userID int) error {
	_, err := db.c.Exec(query_DELEUSERTOGROUP, groupID, userID)
	return err
}
