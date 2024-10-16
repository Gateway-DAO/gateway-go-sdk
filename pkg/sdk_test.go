package pkg_test

import (
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewSDK(t *testing.T) {
	apiKey := "test-api-key"

	// Call NewSDK to create a new instance
	sdkInstance := pkg.NewSDK(apiKey)

	// Assertions to check if the SDK is initialized correctly
	assert.NotNil(t, sdkInstance)
	assert.Equal(t, apiKey, sdkInstance.APIKey)

	// Check that the resty client is created and configured
	assert.NotNil(t, sdkInstance.DataAssets)
	assert.NotNil(t, sdkInstance.DataModel)
	assert.NotNil(t, sdkInstance.Auth)
	assert.NotNil(t, sdkInstance.ACL)
	assert.NotNil(t, sdkInstance.Account)

	// Additional checks on the resty client if needed
	client := sdkInstance.Account.Config.Client
	assert.IsType(t, &resty.Client{}, client)
	assert.Equal(t, "https://dev.api.gateway.tech", client.BaseURL)
}
