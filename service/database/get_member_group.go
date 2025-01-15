package database

var query_GETMEMBERGROUP = `SELECT userID FROM MemberGroup WHERE groupID = ?;`

func (db *appdbimpl) GetMemberGroup(groupID int) ([]int, error) {
	var userID int
	var userIDs []int
	rows, err := db.c.Query(query_GETMEMBERGROUP, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}
