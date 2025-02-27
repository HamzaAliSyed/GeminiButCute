package controller

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

func RegisterUserHandle(response http.ResponseWriter, request *http.Request) {

	if cORSError := utils.AllowPreflightAndCORSForFrontend(response, *request); cORSError != nil {
		if request.Method == http.MethodOptions {
			return
		}
	}

	if request.Method != http.MethodPost {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.Users
	if jsonDecodingError := json.NewDecoder(request.Body).Decode(&newUser); jsonDecodingError != nil {
		http.Error(response, "Invalid Payload from frontend", http.StatusBadRequest)
		return
	}

	if userCreationError := services.CreateUser(newUser); userCreationError != nil {
		http.Error(response, userCreationError.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]string{
		"message": "User created successfully",
	})
}
