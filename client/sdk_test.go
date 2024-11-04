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

func TestSDK_Reinitialize_WithAPIKey(t *testing.T) {
	config := client.SDKConfig{
		ApiKey: "test-api-key",
		URL:    "https://example.com",
	}

	sdk := client.NewSDK(config)
	newConfig := client.SDKConfig{
		ApiKey: "new-test-api-key",
		URL:    "https://new-example.com",
	}

	reinitializedSDK := sdk.Reinitialize(newConfig)

	assert.NotNil(t, reinitializedSDK, "Reinitialized SDK instance should not be nil")
	assert.NotNil(t, reinitializedSDK.Auth, "Auth should not be nil")
	assert.NotNil(t, reinitializedSDK.DataAssets, "DataAssets should not be nil")
	assert.Equal(t, "https://new-example.com", reinitializedSDK.Account.Config.Client.BaseURL, "BaseURL should be updated correctly")
}

func TestSDK_Reinitialize_WithoutAPIKey_UseWallet(t *testing.T) {
	initialWalletDetails := client.WalletDetails{
		PrivateKey: "edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a",
		WalletType: client.Ethereum,
	}
	config := client.SDKConfig{
		WalletDetails: initialWalletDetails,
		URL:           "https://example.com",
	}

	sdk := client.NewSDK(config)

	newWalletDetails := client.WalletDetails{
		PrivateKey: "T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3",
		WalletType: client.Solana,
	}
	newConfig := client.SDKConfig{
		WalletDetails: newWalletDetails,
		URL:           "https://new-example.com",
	}

	reinitializedSDK := sdk.Reinitialize(newConfig)

	assert.NotNil(t, reinitializedSDK, "Reinitialized SDK instance should not be nil")
	assert.NotNil(t, reinitializedSDK.Auth, "Auth should not be nil")
	assert.NotNil(t, reinitializedSDK.DataAssets, "DataAssets should not be nil")
	assert.Equal(t, "https://new-example.com", reinitializedSDK.Account.Config.Client.BaseURL, "BaseURL should be updated correctly")
}
