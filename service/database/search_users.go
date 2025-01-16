package database

var query_SEARCHUSER = `SELECT userID, username FROM User WHERE username LIKE ?`

func (db *appdbimpl) SearchUser(usersearch string) ([]User, error) {

	var users []User

	rows, err := db.c.Query(query_SEARCHUSER, "%"+usersearch+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}

		var user User
		err := rows.Scan(&user.UserID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
