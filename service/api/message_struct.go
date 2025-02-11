package api

import (
	"myWasatext/service/database"
)

// Message struct
type Message struct {
	MessageID      int    `json:"messageID"`
	ConversationID int    `json:"conversationID"`
	UserID         int    `json:"userID"`
	MessageTXT     string `json:"txt"`
}

// from database
func (m *Message) FromDatabase(dbMessage database.Message) error {
	m.MessageID = dbMessage.MessageID
	m.ConversationID = dbMessage.ConversationID
	m.UserID = dbMessage.UserID
	m.MessageTXT = dbMessage.MessageTXT

	return nil
}

// to database
func (m *Message) ToDatabase() database.Message {
	return database.Message{
		MessageID:      m.MessageID,
		ConversationID: m.ConversationID,
		UserID:         m.UserID,
		MessageTXT:     m.MessageTXT,
	}
}
