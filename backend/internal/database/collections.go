package database

import "go.mongodb.org/mongo-driver/mongo"

type Collections struct {
	Users *mongo.Collection
}

func GetCollections() *Collections {
	databaseCollections := MongoDBClient.Database("GeminiButCute")
	usersCollections := databaseCollections.Collection("users")
	return &Collections{
		usersCollections,
	}
}
