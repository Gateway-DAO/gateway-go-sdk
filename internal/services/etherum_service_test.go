package services_test

import (
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestNewEtherumService(t *testing.T) {
	privateKeyHex := "4c0883a69102937d6238470e2e8d3e2a5b6a3d06c4c7a4d8f1fbd9e6a9d6b2c3" // Sample private key
	service := services.NewEtherumService(privateKeyHex)

	// Validate the wallet address
	expectedAddress := "0x171844F2Dc3E20237710f35672f4d76A83EBd3f4" // Adjust based on your private key
	assert.Equal(t, expectedAddress, service.GetWallet())

}

func TestSignMessage(t *testing.T) {
	privateKeyHex := "4c0883a69102937d6238470e2e8d3e2a5b6a3d06c4c7a4d8f1fbd9e6a9d6b2c3" // Sample private key
	service := services.NewEtherumService(privateKeyHex)

	message := "Hello, Ethereum!"
	signature, err := service.SignMessage(message)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, signature.Signature)
	assert.Equal(t, service.GetWallet(), signature.SigningKey)
}

func TestVerifyMessage(t *testing.T) {
	privateKeyHex := "4c0883a69102937d6238470e2e8d3e2a5b6a3d06c4c7a4d8f1fbd9e6a9d6b2c3" // Sample private key
	service := services.NewEtherumService(privateKeyHex)

	message := "Hello, Ethereum!"
	signature, err := service.SignMessage(message)
	assert.NoError(t, err)

	// Verify the message
	isValid, err := service.VerifyMessage(signature.Signature, message, service.GetWallet())
	assert.NoError(t, err)
	assert.True(t, isValid)

	// Test with an invalid address
	isValid, err = service.VerifyMessage(signature.Signature, message, "0xInvalidAddress")
	assert.NoError(t, err)
	assert.False(t, isValid)
}

func TestValidateWallet(t *testing.T) {
	service := services.EtherumService{}

	// Valid address
	validAddress := "0x6c37A84A397C12F3E0A8C6E2B0B5E49D2502E0E4"
	assert.True(t, service.ValidateWallet(validAddress))

	// Invalid address
	invalidAddress := "InvalidAddress"
	assert.False(t, service.ValidateWallet(invalidAddress))
}
