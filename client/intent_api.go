package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ameykpatil/chatbot/model"
)

// IntentClient provides access to intent-service endpoints
type IntentClient struct {
	client *http.Client
	url    string
	apiKey string
}

// NewIntentClient returns an intent-service client that gives facility to call APIs of intent-service
func NewIntentClient(client *http.Client, url, apiKey string) *IntentClient {
	return &IntentClient{
		client: client,
		url:    url,
		apiKey: apiKey,
	}
}

// GetIntentsPayload is a payload struct that holds bot-id & message
type GetIntentsPayload struct {
	BotID   string `json:"botId"`
	Message string `json:"message"`
}

// GetIntentsResponse is response struct to hold response coming from api
type GetIntentsResponse struct {
	Intents  []model.Intent `json:"intents"`
	Entities []struct{}     `json:"entities"`
}

// GetIntents call intent api & returns the response
func (ic *IntentClient) GetIntents(ctx context.Context, botID, message string) (*GetIntentsResponse, error) {

	url := fmt.Sprintf("%s/api/intents", ic.url)

	payload := GetIntentsPayload{
		BotID:   botID,
		Message: message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshalling body for new request to intents-service : %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("creating new request to intent-service : %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", ic.apiKey)

	res, err := ic.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing Request to intent-service : %w", err)
	}

	err = CheckResponse(res)
	if err != nil {
		return nil, fmt.Errorf("checking response from intent-service : %w", err)
	}

	var response GetIntentsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("decoding response from intent-service api : %w", err)
	}

	return &response, nil
}

// CheckResponse returns an error if HTTP response code doesn't satisfy success case
func CheckResponse(resp *http.Response) error {
	sc := resp.StatusCode
	if 200 <= sc && sc <= 299 {
		return nil
	}

	return fmt.Errorf("http response with status code %d", resp.StatusCode)
}
