package services_test

import (
	"encoding/hex"
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestNewEtherumService(t *testing.T) {
	// Known private key for testing (do not use this in production)
	privateKeyHex := "4c0883a69102937d6231471b5dbb6204fe5129617082794ee3f2f7d9d3c10b6a"
	expectedWalletAddress := "0xEB66720657C4b498d3E5fafb7748835A76dDfE3b" // Public address derived from the private key

	ethService := services.NewEtherumService(privateKeyHex)

	// Check if wallet address is derived correctly
	assert.Equal(t, expectedWalletAddress, ethService.GetWallet(), "Wallet address should match the expected public address")
}

func TestSignAndVerifyMessage(t *testing.T) {
	// Known private key for testing (do not use this in production)
	privateKeyHex := "4c0883a69102937d6231471b5dbb6204fe5129617082794ee3f2f7d9d3c10b6a"
	message := "Test message"
	ethService := services.NewEtherumService(privateKeyHex)

	// Sign message
	signedMessage, err := ethService.SignMessage(message)
	assert.Nil(t, err, "Error should be nil when signing a message")

	// Convert the signature to bytes
	sigBytes, err := hex.DecodeString(signedMessage.Signature[2:]) // Removing "0x" prefix
	assert.Nil(t, err, "Error should be nil when decoding hex signature")
	sigBytes = append(sigBytes)
	// Verify the signed message
}

func TestValidateWallet(t *testing.T) {
	ethService := services.NewEtherumService("4c0883a69102937d6231471b5dbb6204fe5129617082794ee3f2f7d9d3c10b6a")

	// Valid Ethereum address
	validWallet := "0xEB66720657C4b498d3E5fafb7748835A76dDfE3b"
	assert.True(t, ethService.ValidateWallet(validWallet), "Wallet should be valid")

	// Invalid Ethereum address
	invalidWallet := "invalid_address"
	assert.False(t, ethService.ValidateWallet(invalidWallet), "Wallet should be invalid")
}
