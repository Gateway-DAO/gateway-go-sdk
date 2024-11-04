package client_test

import (
	"errors"
	"net/http"
	"testing"

	gateway "github.com/Gateway-DAO/gateway-go-sdk/client"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestWalletImpl(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := gateway.Config{
		Client: client,
	}

	walletImpl := gateway.NewWalletImpl(config)

	t.Run("TestAddWallet", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"username": "testuser", "WalletAddresses": ["0xTestAddress"]}`
		httpmock.RegisterResponder("POST", gateway.AddWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		result, err := walletImpl.Add("0xTestAddress")

		assert.NoError(t, err)
		assert.Equal(t, "testuser", result.Username)
		assert.Contains(t, result.Username, "testuser")
	})

	t.Run("TestAddWalletError", func(t *testing.T) {
		httpmock.Reset()

		errorResponse := `{"error": "Failed to add wallet"}`
		httpmock.RegisterResponder("POST", gateway.AddWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, errorResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		result, err := walletImpl.Add("0xTestAddress")

		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})

	t.Run("TestAddWalletHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("POST", gateway.AddWallet, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		result, err := walletImpl.Add("0xTestAddress")

		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})

	t.Run("TestRemoveWallet", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"username": "testuser", "wallets": []}`
		httpmock.RegisterResponder("DELETE", gateway.RemoveWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		result, err := walletImpl.Remove("0xTestAddress")

		assert.NoError(t, err)
		assert.Equal(t, "testuser", result.Username)
	})

	t.Run("TestRemoveWalletError", func(t *testing.T) {
		httpmock.Reset()

		errorResponse := `{"error": "Failed to remove wallet"}`
		httpmock.RegisterResponder("DELETE", gateway.RemoveWallet, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, errorResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		result, err := walletImpl.Remove("0xTestAddress")

		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})

	t.Run("TestRemoveWalletHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("DELETE", gateway.RemoveWallet, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		result, err := walletImpl.Remove("0xTestAddress")

		assert.Error(t, err)
		assert.Empty(t, result.WalletAddresses)
	})
}
