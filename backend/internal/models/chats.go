package models

import (
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
	Title     string             `bson:"title" json:"title"`
	Chats     []Chat             `bson:"chats" json:"chats"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}
