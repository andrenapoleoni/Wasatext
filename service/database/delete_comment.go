package database

var query_DELETECOMMENT = `DELETE FROM Comment WHERE commentID = ?`

func (db *appdbimpl) DeleteComment(commentID int) error {
	_, err := db.c.Exec(query_DELETECOMMENT, commentID)

	return err
}
