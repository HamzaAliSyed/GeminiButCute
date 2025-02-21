package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collections struct {
	Users       *mongo.Collection
	Chats       *mongo.Collection
	ChatThreads *mongo.Collection
}

func GetCollections() *Collections {
	databaseCollections := MongoDBClient.Database("GeminiButCute")
	usersCollections := databaseCollections.Collection("users")
	chatCollections := databaseCollections.Collection("chats")
	chatThreadCollections := databaseCollections.Collection("chatThreads")

	return &Collections{
		usersCollections,
		chatCollections,
		chatThreadCollections,
	}
}
