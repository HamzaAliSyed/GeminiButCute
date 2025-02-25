package main

import (
	"backend/internal/database"
	"backend/internal/services"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	database.ConnectToMongo()
	collections := database.GetCollections()
	users := collections.Users
	fmt.Printf("%v]n", users)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not defined")
	}

	sampleUserID, idCreationError := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if idCreationError != nil {
		fmt.Println("Cannot create ID")
	}

	chat, chatError := services.CreateAChat(sampleUserID, "Are atheist persecuted in Pakistan")
	if chatError != nil {
		fmt.Printf("Something went wrong\n%v", chatError)
	}

	fmt.Println(chat)

	log.Printf("HTTP server is starting on port %s", port)

	if serverStartError := http.ListenAndServe(":"+port, nil); serverStartError != nil {
		log.Fatalf("Server failed: %v", serverStartError)
	}
}
