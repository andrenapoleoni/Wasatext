package database

var query_ADDGROUPUSER = `INSERT INTO MemberGroup (groupID,userID)
											VALUES (?,?);`

func (db *appdbimpl) AddUserToGroup(groupID int, userID int) error {
	_, err := db.c.Exec(query_ADDGROUPUSER, groupID, userID)
	if err != nil {
		return err
	}
	return nil
}
