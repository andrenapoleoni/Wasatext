package database

// query_DELETEGROUP is the query to delete a group from the database

var query_DELETEGROUP = `DELETE FROM Groupchat WHERE groupID = ?`

func (db *appdbimpl) DeleteGroup(groupID int) error {
	_, err := db.c.Exec(query_DELETEGROUP, groupID)
	return err
}
