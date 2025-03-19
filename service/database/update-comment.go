package database

var query_UPDATECOMMENT = `UPDATE Comment SET commentTXT = ? WHERE commentID = ?;`

// UpdateComment updates the comment with the given commentID
func (db *appdbimpl) UpdateComment(commentID int, commentTXT string) error {
	_, err := db.c.Exec(query_UPDATECOMMENT, commentTXT, commentID)
	return err
}
