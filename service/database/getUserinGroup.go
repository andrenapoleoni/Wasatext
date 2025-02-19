package database

var query_MEMBERSINGROUP = `SELECT userID FROM MemberGroup where groupID = ?`

func (db *appdbimpl) GetUsersInGroup(groupID int) ([]int, error) {
	var users []int
	rows, err := db.c.Query(query_MEMBERSINGROUP, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user int
		err := rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		if rows.Err() != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
