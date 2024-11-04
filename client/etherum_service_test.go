package client_test

import (
	"fmt"
	"testing"

	gateway "github.com/Gateway-DAO/gateway-go-sdk/client"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/mock"
)

func TestSignMessageEth_Success(t *testing.T) {
	message := "test message"
	ethService := gateway.NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a") // Replace with a valid private key

	signedMessage, err := ethService.SignMessage(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, signedMessage.Signature)
	assert.NotEmpty(t, signedMessage.SigningKey)

	isValid, err := gateway.VerifyEtherumMessage(signedMessage.Signature, message, ethService.WalletAddress)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestVerifyEtherumMessage_Success(t *testing.T) {
	message := "test message"
	ethService := gateway.NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	signedMessage, _ := ethService.SignMessage(message)

	isValid, err := gateway.VerifyEtherumMessage(signedMessage.Signature, message, ethService.WalletAddress)

	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestVerifyEtherumMessage_InvalidSignature(t *testing.T) {
	message := "test message"
	ethService := gateway.NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	signedMessage, _ := ethService.SignMessage(message)

	signedMessage.Signature = ""
	invalidSignature := "0x1234567890abcdef0x1234567890abcdef"

	isValid, err := gateway.VerifyEtherumMessage(invalidSignature, message, ethService.WalletAddress)

	assert.Error(t, err)
	assert.False(t, isValid)
}

func TestVerifyEtherumMessage_InvalidAddress(t *testing.T) {
	message := "test message"
	ethService := gateway.NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	signedMessage, _ := ethService.SignMessage(message)

	invalidAddress := "0xInvalidAddress"

	isValid, err := gateway.VerifyEtherumMessage(signedMessage.Signature, message, invalidAddress)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid signature")
	assert.False(t, isValid)
}

func TestValidateEtherumWallet_Success(t *testing.T) {
	validWallet := "0x225e681f7A54c248340f7e714b25Dc1fFd2Fda0E"

	isValid := gateway.ValidateEtherumWallet(validWallet)

	assert.True(t, isValid)
}

func TestValidateEtherumWallet_Fail(t *testing.T) {
	invalidWallet := "0xInvalidEthereumAddress"

	isValid := gateway.ValidateEtherumWallet(invalidWallet)

	assert.False(t, isValid)
}

func TestGetWallet(t *testing.T) {
	service := gateway.NewEtherumService("edb0ba5a63c5f9e4f4394560907794fca750704b355413bc04baab896254036a")
	wallet := service.GetWallet()

	assert.NotEmpty(t, wallet)
	assert.Equal(t, service.WalletAddress, wallet)
}

func TestVerifyEtherumMessage_InvalidSignatureHex(t *testing.T) {
	message := "test-message"
	walletAddress := "0x225e681f7A54c248340f7e714b25Dc1fFd2Fda0E"
	invalidSignature := "invalid-signature"

	isValid, err := gateway.VerifyEtherumMessage(invalidSignature, message, walletAddress)

	assert.False(t, isValid)
	assert.Error(t, err)
}

func TestNewEtherumService_InvalidPrivateKey(t *testing.T) {
	invalidPrivateKey := "invalid-private-key"
	assert.Panics(t, func() {
		gateway.NewEtherumService(invalidPrivateKey)
	}, "Expected panic due to invalid private key")
}

func TestNewEtherumService_InvalidPublicKey(t *testing.T) {
	invalidPrivateKey := "your-invalid-private-key"
	assert.Panics(t, func() {
		gateway.NewEtherumService(invalidPrivateKey)
	}, "Expected panic due to invalid public key derivation")
}

type MockEtherumService struct {
	mock.Mock
}

func (m *MockEtherumService) SignMessage(message string) (gateway.WalletSignMessageType, error) {
	args := m.Called(message)
	return args.Get(0).(gateway.WalletSignMessageType), args.Error(1)
}

func (m *MockEtherumService) GetWallet() string {
	args := m.Called()
	return args.String(0)
}

func TestSignMessage_SignError(t *testing.T) {
	mockService := new(MockEtherumService)

	mockService.On("SignMessage", "test-message").Return(gateway.WalletSignMessageType{}, fmt.Errorf("failed to sign message"))

	_, err := mockService.SignMessage("test-message")

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to sign message")
}
