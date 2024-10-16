package services_test

import (
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestNewSolanaService(t *testing.T) {
	// Sample private key (base64 encoded)
	privateKeyHex := "3d4f1506a8c0289186c10a391ef493ce38fa8dca8e1c02e2d1ff3b30f3a8c4d8"

	service := services.NewSolanaService(privateKeyHex)

	// Assertions
	assert.NotNil(t, service)
	assert.NotNil(t, service.GetWallet())
}

func TestSolanaSignMessage(t *testing.T) {
	privateKeyHex := "3d4f1506a8c0289186c10a391ef493ce38fa8dca8e1c02e2d1ff3b30f3a8c4d8" // Replace with a valid base64-encoded private key
	service := services.NewSolanaService(privateKeyHex)

	message := "Hello, Solana!"
	signedMessage, err := service.SignMessage(message)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, signedMessage.Signature)
}

func TestSolanaVerifyMessage(t *testing.T) {
	privateKeyHex := "3d4f1506a8c0289186c10a391ef493ce38fa8dca8e1c02e2d1ff3b30f3a8c4d8" // Replace with a valid base64-encoded private key
	service := services.NewSolanaService(privateKeyHex)

	message := "Hello, Solana!"
	signedMessage, err := service.SignMessage(message)
	assert.NoError(t, err)

	// Verify the message
	isValid, err := service.VerifyMessage(message, string(signedMessage.Signature), signedMessage.SigningKey)
	assert.NoError(t, err)
	assert.True(t, isValid)

	// Test with an invalid signature
	isValid, err = service.VerifyMessage(message, "invalid_signature", signedMessage.SigningKey)
	assert.NoError(t, err)
	assert.False(t, isValid)
}

func TestSolanaValidateWallet(t *testing.T) {
	service := services.SolanaService{}

	// Valid address (base58 encoded)
	validAddress := "5hQWmqxdX2hzjGyK3oT8G9sNREcSgXX9HsbMSfS7gHeR" // Replace with a valid address
	validatedAddress, err := service.ValidateWallet(validAddress)
	assert.NoError(t, err)
	assert.Equal(t, validAddress, validatedAddress)

	// Invalid address
	invalidAddress := "InvalidAddress"
	_, err = service.ValidateWallet(invalidAddress)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid wallet address")
}
