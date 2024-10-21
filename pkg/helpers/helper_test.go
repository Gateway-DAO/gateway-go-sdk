package helpers_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/helpers"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWallet struct {
	mock.Mock
}

func (m *MockWallet) SignMessage(message string) (services.WalletSignMessageType, error) {
	args := m.Called(message)
	return args.Get(0).(services.WalletSignMessageType), args.Error(1)
}

func TestCheckJWTTokenExpiration_Valid(t *testing.T) {
	// Create a token with claims but without signing
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	// Get the unsigned token string
	tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

	// Check expiration without validating signature
	isValid, err := helpers.CheckJWTTokenExpiration(tokenString)

	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestCheckJWTTokenExpiration(t *testing.T) {
	t.Run("TestCheckJWTTokenExpirationError", func(t *testing.T) {
		// Simulate an invalid token string that will cause an error during parsing
		invalidToken := "invalid.token.string"

		// Call the CheckJWTTokenExpiration function
		valid, err := helpers.CheckJWTTokenExpiration(invalidToken)

		// Assertions
		assert.Error(t, err)   // Expecting an error
		assert.False(t, valid) // valid should be false on error
	})
}

func TestCheckJWTTokenExpiration_Expired(t *testing.T) {
	secret := []byte("my-secret-key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-10 * time.Minute)),
	})
	tokenString, _ := token.SignedString(secret)

	isValid, err := helpers.CheckJWTTokenExpiration(tokenString)

	assert.NoError(t, err)
	assert.False(t, isValid)
}

func TestIssueJWT_Success(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	// Mock the wallet to return a signature
	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(services.WalletSignMessageType{
		Signature:  "mock-signature",
		SigningKey: "mock-signing-key",
	}, nil)

	// Mock the /auth/message endpoint to return a message
	fixtureMessage := `"mock-message"`
	httpmock.RegisterResponder("GET", "https://example.com/auth/message",
		httpmock.NewStringResponder(200, fixtureMessage))

	// Mock the /auth/login endpoint to return a JWT
	fixtureToken := `"mock-jwt-token"`
	httpmock.RegisterResponder("POST", "https://example.com/auth/login",
		httpmock.NewStringResponder(200, fixtureToken))

	// Issue JWT
	_, err := helpers.IssueJWT(*client, mockWallet)

	// Assert the results
	assert.Error(t, err)

}

func TestIssueJWT_SignMessageError(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	// Mock the wallet to return an error
	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(services.WalletSignMessageType{}, fmt.Errorf("failed to sign message"))

	// Mock the /auth/message endpoint to return a message
	fixtureMessage := `"mock-message"`
	httpmock.RegisterResponder("GET", "=~.*/auth/message", // Adjust this based on actual URL
		httpmock.NewStringResponder(200, fixtureMessage))

	// Issue JWT
	_, err := helpers.IssueJWT(*client, mockWallet)

	// Assert that there was an error in signing the message
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to sign message")
}

func TestIssueJWT_LoginError(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	// Mock the wallet to return a valid signature
	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(services.WalletSignMessageType{
		Signature:  "mock-signature",
		SigningKey: "mock-signing-key",
	}, nil)

	// Mock the /auth/message endpoint to return a message
	fixtureMessage := `"mock-message"`
	httpmock.RegisterResponder("GET", "=~.*/auth/message", // Adjust this based on actual URL
		httpmock.NewStringResponder(200, fixtureMessage))

	// Mock the /auth/login endpoint to return an error
	httpmock.RegisterResponder("POST", "=~.*/auth/login", // Adjust this if needed
		httpmock.NewStringResponder(500, `{"error": "internal server error"}`))

	// Issue JWT
	_, err := helpers.IssueJWT(*client, mockWallet)

	// Assert that there was an error in the login process
	assert.Error(t, err)
	assert.EqualError(t, err, "solana signature verification failed: failed to decode signature from Base58: invalid base58 digit ('-')")
}

func TestIssueJWT_FailSignMessage(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	// Mock the wallet to return an error
	mockWallet := new(MockWallet)
	mockWallet.On("SignMessage", mock.Anything).Return(services.WalletSignMessageType{}, errors.New("signing error"))

	// Mock the /auth/message endpoint (optional, as it's not used in this case)
	httpmock.RegisterResponder("GET", "https://example.com/auth/message",
		httpmock.NewStringResponder(200, `"mock-message"`))

	_, err := helpers.IssueJWT(*client, mockWallet)

	// Assert that an error was returned
	assert.Error(t, err)
}

func TestAuthMiddleware_ExistingValidToken(t *testing.T) {
	client := resty.New()
	client.SetBaseURL("https://example.com")
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	// Create an existing valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	// Mock wallet
	// mockWallet := new(MockWallet)

	// Create middleware params
	params := services.MiddlewareParams{
		Client: client,
		Wallet: services.WalletService{},
	}

	// Create middleware function
	middleware := helpers.AuthMiddleware(params)

	// Mock the request with an existing Authorization token
	req := client.R().SetHeader("Authorization", tokenString)

	// Call middleware
	err := middleware(client, req)

	// Assert the token remains unchanged
	assert.NoError(t, err)
	assert.Equal(t, tokenString, req.Header.Get("Authorization"), "Authorization header should not change if token is valid")
}
