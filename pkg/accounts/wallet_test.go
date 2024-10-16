package accounts_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/accounts"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestWalletImpl(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := common.SDKConfig{
		Client: client,
		ApiKey: "test-api-key",
	}

	walletImpl := accounts.NewWalletImpl(config)

	t.Run("TestAddWallet", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"username": "testuser", "WalletAddresses": ["0xTestAddress"]}`
		httpmock.RegisterResponder("POST", common.AddWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		// Test
		result, err := walletImpl.Add("0xTestAddress")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "testuser", result.Username)
		assert.Contains(t, result.Username, "testuser")
	})

	t.Run("TestAddWalletError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response for error
		errorResponse := `{"error": "Failed to add wallet"}`
		httpmock.RegisterResponder("POST", common.AddWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, errorResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		// Test
		result, err := walletImpl.Add("0xTestAddress")

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})

	t.Run("TestAddWalletHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Simulate a client-side error (e.g., network error)
		httpmock.RegisterResponder("POST", common.AddWallet, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		// Test
		result, err := walletImpl.Add("0xTestAddress")

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})

	t.Run("TestRemoveWallet", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"username": "testuser", "wallets": []}`
		httpmock.RegisterResponder("DELETE", common.RemoveWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		// Test
		result, err := walletImpl.Remove("0xTestAddress")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "testuser", result.Username)
	})

	t.Run("TestRemoveWalletError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response for error
		errorResponse := `{"error": "Failed to remove wallet"}`
		httpmock.RegisterResponder("DELETE", common.RemoveWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, errorResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		// Test
		result, err := walletImpl.Remove("0xTestAddress")

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})

	t.Run("TestRemoveWalletHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Simulate a client-side error (e.g., network error)
		httpmock.RegisterResponder("DELETE", common.RemoveWallet, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		// Test
		result, err := walletImpl.Remove("0xTestAddress")

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})
}
