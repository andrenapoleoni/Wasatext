package database

var query_GETCOMMENTS = `SELECT commentID, commentTXT, userID FROM Comment WHERE (messageID,conversationID)=(?,?);`

func (db *appdbimpl) GetComments(messageID int, conversationID int) ([]Comment, error) {
	// get all comments from database
	var comments []Comment
	rows, err := db.c.Query(query_GETCOMMENTS, messageID, conversationID)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.CommentID, &comment.CommentTXT, &comment.UserID)
		if err != nil {
			return comments, err
		}
		if rows.Err() != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, nil

}
