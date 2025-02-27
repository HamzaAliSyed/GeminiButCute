package utils

import (
	"fmt"
	"net/http"
)

func AllowPreflightAndCORSForFrontend(response http.ResponseWriter, request http.Request) error {
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if request.Method == http.MethodOptions {
		response.WriteHeader(http.StatusOK)
		return fmt.Errorf("preflight request handled")
	}

	return nil
}
