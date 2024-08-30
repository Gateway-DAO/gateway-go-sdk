package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	client := &http.Client{}

	// Example 1: Get data asset by ID
	dataAssetID := "123"
	resp, err := makeRequest(client, "GET", GetDataAssetByID, dataAssetID)
	if err != nil {
		fmt.Printf("Error getting data asset: %v\n", err)
	} else {
		fmt.Printf("Data asset response: %s\n", resp)
	}

	// Example 2: Add funds to account
	accountID := "456"
	resp, err = makeRequest(client, "POST", AddFundsToAccount, accountID)
	if err != nil {
		fmt.Printf("Error adding funds: %v\n", err)
	} else {
		fmt.Printf("Add funds response: %s\n", resp)
	}
}

func makeRequest(client *http.Client, method, endpoint string, params ...interface{}) (string, error) {
	formattedEndpoint := fmt.Sprintf(endpoint, params...)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
