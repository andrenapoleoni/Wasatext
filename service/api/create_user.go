package api

// CreateUser function is used to create a new user
func (rt *_router) CreateUser(u User) (User, error) {
	dbUser, err := rt.db.CreateUser(u.ToDatabase())
	if err != nil {
		return u, err
	}

	// convert the database user object to the user object
	err = u.FromDatabase(dbUser)
	if err != nil {
		return u, err
	}

	return u, nil
}
