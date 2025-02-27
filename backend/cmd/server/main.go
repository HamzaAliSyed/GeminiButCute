package main

import (
	"backend/internal/database"
	"backend/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	database.ConnectToMongo()
	collections := database.GetCollections()
	database.EnforceUniqueUsernameAndEmailIndexes(collections.Users)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not defined")
	}

	log.Printf("HTTP server is starting on port %s", port)
	routes.RegisterUserRoutes()

	if serverStartError := http.ListenAndServe(":"+port, nil); serverStartError != nil {
		log.Fatalf("Server failed: %v", serverStartError)
	}
}
