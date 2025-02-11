package api

import (
	"myWasatext/service/database"
)

type Conversation struct {
	ConversationID int `json:"conversationID"`

	GroupID int `json:"groupID"`
}

func (c *Conversation) ToDatabase() database.Conversation {
	return database.Conversation{
		ConversationID: c.ConversationID,
		GroupID:        c.GroupID,
	}
}

func (c *Conversation) FromDatabase(dbConversation database.Conversation) error {
	c.ConversationID = dbConversation.ConversationID
	c.GroupID = dbConversation.GroupID

	return nil
}
