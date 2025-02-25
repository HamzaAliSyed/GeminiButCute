package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

func ConnectToMongo() {
	environmentImportError := godotenv.Load("../../.env")
	if environmentImportError != nil {
		log.Fatalf("Error loading .env file: %v", environmentImportError)
	}

	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("CONNECTION_STRING is not set in the .env file")
	}

	contextTime, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelContext()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, clientConnectError := mongo.Connect(contextTime, clientOptions)

	if clientConnectError != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", clientConnectError)
	}

	pingError := client.Ping(contextTime, nil)
	if pingError != nil {
		log.Fatalf("Could not ping MongoDB: %v", pingError)
	}

	MongoDBClient = client
	log.Println("Connected to MongoDB successfully!")
}

func EnforceUniqueUsernameAndEmailIndexes(usersCollections *mongo.Collection) {
	indexContext, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelContext()

	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	_, createIndexError := usersCollections.Indexes().CreateMany(indexContext, indexModels)
	if createIndexError != nil {
		log.Fatalf("could not create indexes: %v", createIndexError)
	} else {
		log.Println("Unique indexes for username and email created successfully")
	}
}
