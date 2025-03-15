package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type GenerateContentRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GenerateContentResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content struct {
		Parts []Part `json:"parts"`
	} `json:"content"`
}

// Make NewClient variable so it can be replaced in tests
var NewClient = func(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		baseURL:    "https://generativelanguage.googleapis.com/v1beta",
		httpClient: &http.Client{},
	}
}

func (c *Client) GenerateContent(text string) (string, error) {
	endpoint := fmt.Sprintf("%s/models/%s:generateContent?key=%s", c.baseURL, "gemini-2.0-flash", c.apiKey)

	request := GenerateContentRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: text},
				},
			},
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	resp, err := c.httpClient.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var response GenerateContentResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	if len(response.Candidates) == 0 {
		return "", fmt.Errorf("no response candidates received")
	}

	if len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content parts in response")
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}
