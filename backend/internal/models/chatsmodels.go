package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Message   string             `bson:"message" json:"message"`
	Response  string             `bson:"response" json:"response"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type ChatThread struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Title     string             `bson:"title" json:"title"`
	Chats     []Chat             `bson:"chats" json:"chats"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

func (chat Chat) String() string {
	return fmt.Sprintf(
		"User ID: %s\nMessage: %s\nResponse: %s\nCreated At: %s",
		chat.UserID.Hex(),
		chat.Message,
		chat.Response,
		chat.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
