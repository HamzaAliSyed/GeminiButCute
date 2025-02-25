package services

import (
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type RequestBody struct {
	Contents []Content `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason string  `json:"finishReason"`
		AvgLogprobs  float64 `json:"avgLogprobs"`
	} `json:"candidates"`
}

func CallToGemini(request string) (string, error) {

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API Key not found in .env file")
	}

	requestURL := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", apiKey)

	payload := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: request},
				},
			},
		},
	}

	payloadBytes, marshallingError := json.Marshal(payload)
	if marshallingError != nil {
		return "", fmt.Errorf("error marshalling when trying to send payload\n%v", marshallingError)
	}

	requestToGemini, requestError := http.NewRequest("POST", requestURL, bytes.NewBuffer(payloadBytes))
	if requestError != nil {
		return "", fmt.Errorf("error creating HTTP request: %v", requestError)
	}

	requestToGemini.Header.Set("Content-Type", "application/json")

	singleClient := &http.Client{}
	responseFromGemini, responseError := singleClient.Do(requestToGemini)
	if responseError != nil {
		return "", fmt.Errorf("error sending HTTP request: %v", responseError)
	}

	defer responseFromGemini.Body.Close()

	if responseFromGemini.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(responseFromGemini.Body)
		return "", fmt.Errorf("request failed with status: %s, body: %s", responseFromGemini.Status, string(bodyBytes))
	}

	responseBodyBytes, responseParsingError := io.ReadAll(responseFromGemini.Body)
	if responseParsingError != nil {
		return "", fmt.Errorf("error reading response body: %v", responseParsingError)
	}

	var geminiResponse GeminiResponse
	unmarshalError := json.Unmarshal(responseBodyBytes, &geminiResponse)
	if unmarshalError != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", unmarshalError)
	}

	if len(geminiResponse.Candidates) == 0 || len(geminiResponse.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("unable to response for chat")
	}

	chatString := geminiResponse.Candidates[0].Content.Parts[0].Text
	return chatString, nil
}

func CreateAChat(id primitive.ObjectID, messageRequest string) (*models.Chat, error) {
	response, apiCallError := CallToGemini(messageRequest)
	if apiCallError != nil {
		return nil, fmt.Errorf("error calling Gemini API\n%v", apiCallError)
	}

	chat := &models.Chat{
		UserID:    id,
		Message:   messageRequest,
		Response:  response,
		CreatedAt: time.Now(),
	}

	return chat, nil
}
