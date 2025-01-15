package api

import (
	"myWasatext/service/database"
	"unicode/utf8"
)

type Comment struct {
	CommentID      int    `json:"commentID"`
	ConversationID int    `json:"conversationID"`
	UserID         int    `json:"userID"`
	MessageID      int    `json:"messageID"`
	CommentTXT     string `json:"comment"`
}

func (c Comment) IsValid() bool {
	// Decodifica il primo carattere Unicode nella stringa
	r, _ := utf8.DecodeRuneInString(c.CommentTXT)
	if r == utf8.RuneError {
		return false
	}

	// Verifica se il carattere appartiene a uno dei range degli emoji
	return (r >= 0x1F600 && r <= 0x1F64F) || // Emoticon
		(r >= 0x1F300 && r <= 0x1F5FF) || // Simboli e pittogrammi vari
		(r >= 0x1F680 && r <= 0x1F6FF) || // Trasporti e simboli mappa
		(r >= 0x1F700 && r <= 0x1F77F) || // Simboli alchemici
		(r >= 0x2600 && r <= 0x26FF) || // Simboli vari
		(r >= 0x2700 && r <= 0x27BF) || // Dingbats
		(r >= 0xFE00 && r <= 0xFE0F) || // Variazioni selettori
		(r >= 0x1F900 && r <= 0x1F9FF) || // Simboli supplementari
		(r >= 0x1FA70 && r <= 0x1FAFF) || // Emoji aggiuntivi
		(r >= 0x1F1E6 && r <= 0x1F1FF) // Bandiere (regioni)
}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		CommentID:      c.CommentID,
		ConversationID: c.ConversationID,
		UserID:         c.UserID,
		MessageID:      c.MessageID,
		CommentTXT:     c.CommentTXT,
	}
}

func (c *Comment) FromDatabase(dbComment database.Comment) error {
	c.CommentID = dbComment.CommentID
	c.ConversationID = dbComment.ConversationID
	c.UserID = dbComment.UserID
	c.MessageID = dbComment.MessageID
	c.CommentTXT = dbComment.CommentTXT

	return nil
}
