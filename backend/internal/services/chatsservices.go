package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func CallToGemini(request string) string {

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set")
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
		log.Fatalf("Error marshalling payload: %v", marshallingError)
	}

	requestToGemini, requestError := http.NewRequest("POST", requestURL, bytes.NewBuffer(payloadBytes))
	if requestError != nil {
		log.Fatalf("Error creating HTTP request: %v", requestError)
	}

	requestToGemini.Header.Set("Content-Type", "application/json")

	singleClient := &http.Client{}
	responseFromGemini, responseError := singleClient.Do(requestToGemini)
	if responseError != nil {
		log.Fatalf("Error sending HTTP request: %v", responseError)
	}

	defer responseFromGemini.Body.Close()

	if responseFromGemini.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(responseFromGemini.Body)
		log.Fatalf("Request failed with status: %s, body: %s", responseFromGemini.Status, string(bodyBytes))
	}

	responseBodyBytes, responseParsingError := io.ReadAll(responseFromGemini.Body)
	if responseParsingError != nil {
		log.Fatalf("Error reading response body: %v", responseParsingError)
	}

	var geminiResponse GeminiResponse
	unmarshalError := json.Unmarshal(responseBodyBytes, &geminiResponse)
	if unmarshalError != nil {
		log.Fatalf("Error unmarshalling response: %v", unmarshalError)
	}

	if len(geminiResponse.Candidates) == 0 || len(geminiResponse.Candidates[0].Content.Parts) == 0 {
		log.Fatal("No candidate text found in the response")
	}

	chatString := geminiResponse.Candidates[0].Content.Parts[0].Text
	return chatString
}
