package client_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	gateway "github.com/Gateway-DAO/gateway-go-sdk/client"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWallet struct {
	mock.Mock
}

func (m *MockWallet) SignMessage(message string) (gateway.WalletSignMessageType, error) {
	args := m.Called(message)
	return args.Get(0).(gateway.WalletSignMessageType), args.Error(1)
}

func TestCheckJWTTokenExpiration_Valid(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

	isValid, err := gateway.CheckJWTTokenExpiration(tokenString)

	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestCheckJWTTokenExpiration(t *testing.T) {
	t.Run("TestCheckJWTTokenExpirationError", func(t *testing.T) {
		invalidToken := "invalid.token.string"

		valid, err := gateway.CheckJWTTokenExpiration(invalidToken)

		assert.Error(t, err)
		assert.False(t, valid)
	})
}

func TestCheckJWTTokenExpiration_Expired(t *testing.T) {
	secret := []byte("my-secret-key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-10 * time.Minute)),
	})
	tokenString, _ := token.SignedString(secret)

	isValid, err := gateway.CheckJWTTokenExpiration(tokenString)

	assert.NoError(t, err)
	assert.False(t, isValid)
}

func TestIssueJWT_Success(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(gateway.WalletSignMessageType{
		Signature:  "mock-signature",
		SigningKey: "mock-signing-key",
	}, nil)

	fixtureMessage := `"mock-message"`
	httpmock.RegisterResponder("GET", "https://example.com/auth/message",
		httpmock.NewStringResponder(200, fixtureMessage))

	fixtureToken := `"mock-jwt-token"`
	httpmock.RegisterResponder("POST", "https://example.com/auth/login",
		httpmock.NewStringResponder(200, fixtureToken))

	_, err := gateway.IssueJWT(*client, mockWallet)

	assert.Error(t, err)
}

func TestIssueJWT_SignMessageError(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(gateway.WalletSignMessageType{}, fmt.Errorf("failed to sign message"))

	fixtureMessage := `"mock-message"`
	httpmock.RegisterResponder("GET", "=~.*/auth/message",
		httpmock.NewStringResponder(200, fixtureMessage))

	_, err := gateway.IssueJWT(*client, mockWallet)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to sign message")
}

func TestIssueJWT_LoginError(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(gateway.WalletSignMessageType{
		Signature:  "mock-signature",
		SigningKey: "mock-signing-key",
	}, nil)

	fixtureMessage := `"mock-message"`
	httpmock.RegisterResponder("GET", "=~.*/auth/message",
		httpmock.NewStringResponder(200, fixtureMessage))

	httpmock.RegisterResponder("POST", "=~.*/auth/login",
		httpmock.NewStringResponder(500, `{"error": "internal server error"}`))

	_, err := gateway.IssueJWT(*client, mockWallet)

	assert.Error(t, err)
	assert.EqualError(t, err, "solana signature verification failed: failed to decode signature from Base58: invalid base58 digit ('-')")
}

func TestIssueJWT_FailSignMessage(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(gateway.WalletSignMessageType{}, errors.New("signing error"))

	httpmock.RegisterResponder("GET", "https://example.com/auth/message",
		httpmock.NewStringResponder(200, `"mock-message"`))

	_, err := gateway.IssueJWT(*client, mockWallet)

	assert.Error(t, err)
}

func TestAuthMiddleware_ExistingValidToken(t *testing.T) {
	client := resty.New()
	client.SetBaseURL("https://example.com")
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	params := gateway.MiddlewareParams{
		Client: client,
		Wallet: gateway.WalletService{},
	}

	middleware := gateway.AuthMiddleware(params)

	req := client.R().SetHeader("Authorization", tokenString)

	err := middleware(client, req)

	assert.NoError(t, err)
	assert.Equal(t, tokenString, req.Header.Get("Authorization"), "Authorization header should not change if token is valid")
}
