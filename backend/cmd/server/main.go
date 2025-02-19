package main

import (
	"backend/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"
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

	log.Printf("HTTP server is starting on port %s", port)

	if serverStartError := http.ListenAndServe(":"+port, nil); serverStartError != nil {
		log.Fatalf("Server failed: %v", serverStartError)
	}
}
