package main

import (
	"backend/internal/database"
	"backend/internal/services"
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

	response := services.CallToGemini("Goto the interent and find me the latest score of Babar Azam in an international match which is an ODI")
	fmt.Println(response)

	if serverStartError := http.ListenAndServe(":"+port, nil); serverStartError != nil {
		log.Fatalf("Server failed: %v", serverStartError)
	}
}
