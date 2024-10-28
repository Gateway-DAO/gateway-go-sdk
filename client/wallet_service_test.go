package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWalletService_Ethereum(t *testing.T) {
	// Mock Ethereum private key (replace with a valid one for real tests)
	mockPrivateKey := "edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a" // Ethereum private key in hex format

	// Call NewWalletService with Ethereum type
	walletService, err := NewWalletService(mockPrivateKey, Ethereum)

	// Assertions
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, walletService, "WalletService should not be nil")
	assert.Equal(t, Ethereum, walletService.walletType, "Wallet type should be Ethereum")
	assert.Equal(t, mockPrivateKey, walletService.walletPrivKey, "Private key should match")
}

func TestNewWalletService_Solana(t *testing.T) {
	// Mock Solana private key (replace with a valid one for real tests)
	mockPrivateKey := "T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3" // Solana private key in base58 format

	// Call NewWalletService with Solana type
	walletService, err := NewWalletService(mockPrivateKey, Solana)

	// Assertions
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, walletService, "WalletService should not be nil")
	assert.Equal(t, Solana, walletService.walletType, "Wallet type should be Solana")
	assert.Equal(t, mockPrivateKey, walletService.walletPrivKey, "Private key should match")
}

func TestWalletService_SignMessage_Ethereum(t *testing.T) {
	// Mock Ethereum private key (replace with a valid one for real tests)
	mockPrivateKey := "edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a" // Ethereum private key in hex format

	// Call NewWalletService
	walletService, err := NewWalletService(mockPrivateKey, Ethereum)
	assert.NoError(t, err)

	// Sign a message
	message := "test message"
	signature, err := walletService.SignMessage(message)

	// Assertions
	assert.NoError(t, err, "Error should be nil")
	assert.NotEmpty(t, signature.Signature, "Signature should not be empty")
	assert.NotEmpty(t, signature.SigningKey, "SigningKey should not be empty")
}

func TestWalletService_SignMessage_Solana(t *testing.T) {
	// Mock Solana private key (replace with a valid one for real tests)
	mockPrivateKey := "T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3" // Solana private key in base58 format

	// Call NewWalletService
	walletService, err := NewWalletService(mockPrivateKey, Solana)
	assert.NoError(t, err)

	// Sign a message
	message := "test message"
	signature, err := walletService.SignMessage(message)

	// Assertions
	assert.NoError(t, err, "Error should be nil")
	assert.NotEmpty(t, signature.Signature, "Signature should not be empty")
	assert.NotEmpty(t, signature.SigningKey, "SigningKey should not be empty")
}
