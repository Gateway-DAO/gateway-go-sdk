package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifySolanaMessage_Success(t *testing.T) {
	message := "test message"
	solanaService := NewSolanaService("T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3")
	signedMessage, _ := solanaService.SignMessage(message)

	isValid, err := VerifySolanaMessage(message, signedMessage.Signature, signedMessage.SigningKey)

	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestVerifySolanaMessage_InvalidSignature(t *testing.T) {
	message := "test message"
	solanaService := NewSolanaService("T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3")
	signedMessage, _ := solanaService.SignMessage(message)

	invalidSignature := "invalidsignature123"

	isValid, err := VerifySolanaMessage(message, invalidSignature, signedMessage.SigningKey)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode signature")
	assert.False(t, isValid)
}

func TestVerifySolanaMessage_InvalidPublicKey(t *testing.T) {
	message := "test message"
	solanaService := NewSolanaService("T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3")
	signedMessage, _ := solanaService.SignMessage(message)

	invalidPublicKey := "invalidpublickey123"

	isValid, err := VerifySolanaMessage(message, signedMessage.Signature, invalidPublicKey)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode public key")
	assert.False(t, isValid)
}

func TestValidateSolanaWallet_Success(t *testing.T) {
	validWallet := "AqzrrxaBCXRsq2BaY32djAp38B42asRRahbsYvD5uvSF" // Replace with a valid Base58 public key

	isValid := ValidateSolanaWallet(validWallet)

	assert.True(t, isValid)
}

func TestValidateSolanaWallet_Fail(t *testing.T) {
	invalidWallet := "invalid-base58-wallet-public-key"

	isValid := ValidateSolanaWallet(invalidWallet)

	assert.False(t, isValid)
}

func TestSignMessage_Success(t *testing.T) {
	message := "test message"
	solanaService := NewSolanaService("T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3") // Replace with a valid private key

	signedMessage, err := solanaService.SignMessage(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, signedMessage.Signature)
	assert.NotEmpty(t, signedMessage.SigningKey)

	isValid, err := VerifySolanaMessage(message, signedMessage.Signature, signedMessage.SigningKey)
	assert.NoError(t, err)
	assert.True(t, isValid)
}
