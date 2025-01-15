package database

var query_CHANGEUSERNAME = `UPDATE User SET username = ? WHERE userID = ?;`

var query_CHANGEGROUPNAME = `UPDATE Groupchat SET groupName = ? WHERE groupID = ?;`

func (db *appdbimpl) ChangeUsername(userID int, newUsername string) error {
	_, err := db.c.Exec(query_CHANGEUSERNAME, newUsername, userID)
	return err
}

func (db *appdbimpl) ChangeGroupName(groupID int, newGroupName string) error {
	_, err := db.c.Exec(query_CHANGEGROUPNAME, newGroupName, groupID)
	return err
}
