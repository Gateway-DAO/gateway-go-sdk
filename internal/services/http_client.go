package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type APIClient struct {
	Client *http.Client
}

func NewAPIClient(apiKey string) *APIClient {
	return &APIClient{
		Client: &http.Client{},
	}
}

func MakeRequest(client *http.Client, method, endpoint string, params ...interface{}) (string, error) {
	formattedEndpoint := fmt.Sprintf(endpoint, params...)
	BaseURL := ""
	fullURL, err := url.JoinPath(BaseURL, formattedEndpoint)
	if err != nil {
		return "", fmt.Errorf("error joining URL: %w", err)
	}

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
