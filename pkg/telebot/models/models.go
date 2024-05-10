package models

import "fmt"

type Update struct {
	UpdateID      int      `json:"update_id"`
	Message       *Message `json:"message,omitempty"`
	EditedMessage *Message `json:"edited_message,omitempty"`
}

type Message struct {
	MessageID int             `json:"message_id"`
	From      *User           `json:"from,omitempty"`
	Date      int             `json:"date"`
	Chat      *Chat           `json:"chat,omitempty"`
	Text      string          `json:"text,omitempty"`
	Entities  []MessageEntity `json:"entities,omitempty"`
}

type MessageEntityType string

const (
	MessageEntityTypeMention    MessageEntityType = "mention"
	MessageEntityTypeHashtag    MessageEntityType = "hashtag"
	MessageEntityTypeCashtag    MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand MessageEntityType = "bot_command"
)

type MessageEntity struct {
	Type   MessageEntityType
	Offset int
	Url    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

type Chat struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title,omitempty"`
	Username string `json:"username,omitempty"`
	//TODO: photo(ChatPhoto)
}

func (c Chat) String() string {
	return fmt.Sprintf("ID: %d, Type: %s, Title: %s, Username: %s", c.ID, c.Type, c.Title, c.Username)

}

type User struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
}
