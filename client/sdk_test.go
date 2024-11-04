package client_test

import (
	"testing"

	client "github.com/Gateway-DAO/gateway-go-sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestNewSDK_WithAPIKey(t *testing.T) {
	config := client.SDKConfig{
		ApiKey: "test-api-key",
		URL:    "https://example.com",
	}

	sdk := client.NewSDK(config)

	assert.NotNil(t, sdk, "SDK instance should not be nil")
	assert.NotNil(t, sdk.Auth, "Auth should not be nil")
	assert.NotNil(t, sdk.DataAssets, "DataAssets should not be nil")
}

func TestNewSDK_WithoutAPIKey_UseWallet(t *testing.T) {
	walletDetails := client.WalletDetails{
		PrivateKey: "edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a",
		WalletType: client.Ethereum,
	}

	config := client.SDKConfig{
		WalletDetails: walletDetails,
		URL:           "https://example.com",
	}

	sdk := client.NewSDK(config)

	assert.NotNil(t, sdk, "SDK instance should not be nil")
	assert.NotNil(t, sdk.Auth, "Auth should not be nil")
	assert.NotNil(t, sdk.DataAssets, "DataAssets should not be nil")
	assert.NotNil(t, sdk.Account, "Account should not be nil")

	assert.Equal(t, "https://example.com", sdk.Account.Config.Client.BaseURL, "BaseURL should be set correctly")
}
