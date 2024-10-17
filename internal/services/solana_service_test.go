package services_test

import (
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestNewSolanaService(t *testing.T) {
	// Known valid Solana private key (Base58-encoded)
	privateKeyBase58 := "5f44d72YmGb68Dm9AYMo6hfG4bPv5URTgfay63yxDFmc1AEWrYyiVuDhktkWUKr5kduCZWyDaJa8rBPUuzW3sqzn" // Replace this with a valid test key

	// Instantiate SolanaService
	solanaService := services.NewSolanaService(privateKeyBase58)

	// Check if the wallet was created successfully
	assert.NotNil(t, solanaService.GetWallet(), "Wallet should be initialized correctly")
}

func TestValidateSolanaWallet(t *testing.T) {
	// Known private key for testing
	privateKeyBase58 := "5f44d72YmGb68Dm9AYMo6hfG4bPv5URTgfay63yxDFmc1AEWrYyiVuDhktkWUKr5kduCZWyDaJa8rBPUuzW3sqzn"
	solanaService := services.NewSolanaService(privateKeyBase58)

	// Valid Solana wallet address
	validWallet := solanaService.GetWallet()
	walletAddress, err := solanaService.ValidateWallet(validWallet)
	assert.Nil(t, err, "Error should be nil when validating wallet")
	assert.Equal(t, validWallet, walletAddress, "Wallet should be valid")

	// Invalid wallet address
	invalidWallet := "invalid_address"
	_, err = solanaService.ValidateWallet(invalidWallet)
	assert.NotNil(t, err, "There should be an error for invalid wallet address")
}
