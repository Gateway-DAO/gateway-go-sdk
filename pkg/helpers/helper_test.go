package helpers_test

import (
	"errors"
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

// func TestAuthMiddleware_NewToken(t *testing.T) {
// 	client := resty.New()
// 	client.SetBaseURL("https://example.com")
// 	httpmock.ActivateNonDefault(client.GetClient())
// 	defer httpmock.DeactivateAndReset()

// 	// Mock wallet behavior
// 	mockWallet := new(MockWallet)
// 	mockWallet.On("SignMessage", mock.Anything).Return(services.WalletSignMessageType{
// 		Signature:  "mock-signature",
// 		SigningKey: "mock-signing-key",
// 	}, nil)

// 	// Mock the /auth/message and /auth/login endpoints
// 	httpmock.RegisterResponder("GET", "https://example.com/auth/message",
// 		httpmock.NewStringResponder(200, `"mock-message"`))
// 	httpmock.RegisterResponder("POST", "https://example.com/auth/login",
// 		httpmock.NewStringResponder(200, `"mock-jwt-token"`))

// 	// Create middleware params
// 	params := services.MiddlewareParams{
// 		Client: client,
// 		Wallet: services.WalletService{

// 		},
// 	}

// 	// Create middleware function
// 	middleware := helpers.AuthMiddleware(params)

// 	// Mock the request without an Authorization token
// 	req := client.R().SetHeader("Authorization", "")

// 	// Call middleware
// 	err := middleware(client, req)

// 	// Check if token was set
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, req.Header.Get("Authorization"), "Authorization header should not be empty")
// }

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
