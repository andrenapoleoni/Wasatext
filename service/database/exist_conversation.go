package database

var query_TAKECONVERSATION = ` SELECT conversationID FROM Conversation WHERE conversationID = ?;`

func (db *appdbimpl) ExistConversation(userID1 int, userID2 int) (bool, error) {
	// check if conversation exists in private member table
	var exist bool
	err := db.c.QueryRow(`
    SELECT EXISTS(
        SELECT 1
        FROM MemberPrivate mp1
        JOIN MemberPrivate mp2 ON mp1.conversationID = mp2.conversationID
        WHERE mp1.userID = ?
          AND mp2.userID = ?
    );
`, userID1, userID2).Scan(&exist)

	if err != nil {
		return false, err
	}

	return exist, nil

}

func (db *appdbimpl) ExistConversationByID(conversationID int) (bool, error) {
	// check if conversation exists in private member table
	var exist bool
	exist = false
	var conveID int
	err := db.c.QueryRow(query_TAKECONVERSATION, conversationID).Scan(&conveID)
	if err != nil {
		return false, err
	}
	if conveID == conversationID {
		exist = true
	}

	return exist, nil

}
