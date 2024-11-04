package client_test

import (
	"testing"

	gateway "github.com/Gateway-DAO/gateway-go-sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletService_Ethereum(t *testing.T) {
	mockPrivateKey := "edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a"

	walletService, err := gateway.NewWalletService(mockPrivateKey, gateway.Ethereum)

	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, walletService, "WalletService should not be nil")
	assert.Equal(t, gateway.Ethereum, walletService.WalletType, "Wallet type should be Ethereum")
	assert.Equal(t, mockPrivateKey, walletService.WalletPrivKey, "Private key should match")
}

func TestNewWalletService_Solana(t *testing.T) {
	mockPrivateKey := "T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3"

	walletService, err := gateway.NewWalletService(mockPrivateKey, gateway.Solana)

	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, walletService, "WalletService should not be nil")
	assert.Equal(t, gateway.Solana, walletService.WalletType, "Wallet type should be Solana")
	assert.Equal(t, mockPrivateKey, walletService.WalletPrivKey, "Private key should match")
}

func TestWalletService_SignMessage_Ethereum(t *testing.T) {
	mockPrivateKey := "edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a"

	walletService, err := gateway.NewWalletService(mockPrivateKey, gateway.Ethereum)
	assert.NoError(t, err)

	message := "test message"
	signature, err := walletService.SignMessage(message)

	assert.NoError(t, err, "Error should be nil")
	assert.NotEmpty(t, signature.Signature, "Signature should not be empty")
	assert.NotEmpty(t, signature.SigningKey, "SigningKey should not be empty")
}

func TestWalletService_SignMessage_Solana(t *testing.T) {
	mockPrivateKey := "T8HMDTLmyQgY6VjvLdEwSSZsexAtiFvfiKBzEsT3ajNQg7jJgnTBK2qDSShz98ND3ihtrwrQcUWokdQr4ozPQt3" // Solana private key in base58 format

	walletService, err := gateway.NewWalletService(mockPrivateKey,gateway. Solana)
	assert.NoError(t, err)

	message := "test message"
	signature, err := walletService.SignMessage(message)

	assert.NoError(t, err, "Error should be nil")
	assert.NotEmpty(t, signature.Signature, "Signature should not be empty")
	assert.NotEmpty(t, signature.SigningKey, "SigningKey should not be empty")
}
