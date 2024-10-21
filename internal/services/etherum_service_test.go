package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignMessageEth_Success(t *testing.T) {
	// Test data
	message := "test message"
	ethService := NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a") // Replace with a valid private key

	// Act: sign the message
	signedMessage, err := ethService.SignMessage(message)

	// Assert: should sign successfully
	assert.NoError(t, err)
	assert.NotEmpty(t, signedMessage.Signature)
	assert.Equal(t, ethService.walletAddress, signedMessage.SigningKey)

	// Verify the signature
	isValid, err := VerifyEtherumMessage(signedMessage.Signature, message, ethService.walletAddress)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestVerifyEtherumMessage_Success(t *testing.T) {
	// Test data
	message := "test message"
	ethService := NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	signedMessage, _ := ethService.SignMessage(message)

	// Act: verify the signed message
	isValid, err := VerifyEtherumMessage(signedMessage.Signature, message, ethService.walletAddress)

	// Assert: should verify successfully
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestVerifyEtherumMessage_InvalidSignature(t *testing.T) {
	// Test data
	message := "test message"
	ethService := NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	signedMessage, _ := ethService.SignMessage(message)

	signedMessage.Signature = ""
	// Modify the signature to make it invalid
	invalidSignature := "0x1234567890abcdef0x1234567890abcdef"

	// Act: verify with the invalid signature
	isValid, err := VerifyEtherumMessage(invalidSignature, message, ethService.walletAddress)

	// Assert: verification should fail
	assert.Error(t, err)
	assert.False(t, isValid)
}

func TestVerifyEtherumMessage_InvalidAddress(t *testing.T) {
	// Test data
	message := "test message"
	ethService := NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	signedMessage, _ := ethService.SignMessage(message)

	// Use an invalid wallet address
	invalidAddress := "0xInvalidAddress"

	// Act: verify with the invalid address
	isValid, err := VerifyEtherumMessage(signedMessage.Signature, message, invalidAddress)

	// Assert: verification should fail
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid signature")
	assert.False(t, isValid)
}

func TestValidateEtherumWallet_Success(t *testing.T) {
	// Test valid Ethereum wallet address
	validWallet := "0x225e681f7A54c248340f7e714b25Dc1fFd2Fda0E" // Replace with a valid hex Ethereum address

	// Act: validate the wallet
	isValid := ValidateEtherumWallet(validWallet)

	// Assert: validation should succeed
	assert.True(t, isValid)
}

func TestValidateEtherumWallet_Fail(t *testing.T) {
	// Test invalid Ethereum wallet address
	invalidWallet := "0xInvalidEthereumAddress"

	// Act: validate the wallet
	isValid := ValidateEtherumWallet(invalidWallet)

	// Assert: validation should fail
	assert.False(t, isValid)
}

func TestGetWallet(t *testing.T) {
	service := NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	wallet := service.GetWallet()

	assert.NotEmpty(t, wallet)
	assert.Equal(t, service.walletAddress, wallet)
}

func TestVerifyEtherumMessage_InvalidSignatureHex(t *testing.T) {
	message := "test-message"
	walletAddress := "0x225e681f7A54c248340f7e714b25Dc1fFd2Fda0E"
	invalidSignature := "invalid-signature"

	isValid, err := VerifyEtherumMessage(invalidSignature, message, walletAddress)

	assert.False(t, isValid)
	assert.Error(t, err)
}

func TestNewEtherumService_InvalidPrivateKey(t *testing.T) {
	invalidPrivateKey := "invalid-private-key"
	assert.Panics(t, func() {
		NewEtherumService(invalidPrivateKey)
	}, "Expected panic due to invalid private key")
}

func TestNewEtherumService_InvalidPublicKey(t *testing.T) {
	invalidPrivateKey := "your-invalid-private-key"
	assert.Panics(t, func() {
		NewEtherumService(invalidPrivateKey)
	}, "Expected panic due to invalid public key derivation")
}
