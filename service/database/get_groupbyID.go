package database

var query_FINDGROUPBYID = `SELECT groupID, groupName FROM Groupchat WHERE groupID = ?`

func (db *appdbimpl) GetGroupByID(groupID int) (Group, error) {
	var group Group
	err := db.c.QueryRow(query_FINDGROUPBYID, groupID).Scan(&group.GroupID, &group.Name)
	return group, err
}
