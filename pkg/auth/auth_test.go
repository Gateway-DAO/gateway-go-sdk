package auth_test

import (
	"net/http"
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/auth"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)


type MockClient struct {
	mock.Mock
}

func (m *MockClient) R() *resty.Request {
	args := m.Called()
	return args.Get(0).(*resty.Request)
}

// TestAuthImpl_Login tests the Login method
func TestAuthImpl_Login(t *testing.T) {
	// Mock the HTTP client
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Success case
	httpmock.RegisterResponder("POST", common.AuthenticateAccount, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, common.TokenResponse{Token: "mock-token"})
	})

	mockClient := new(MockClient)

	// Create a mock request
	mockRequest := new(resty.Request)
	mockClient.On("R").Return(mockRequest)

	// Setup your SDKConfig with the mock client
	config := common.SDKConfig{
		Client: mockClient,
	}

	config := common.SDKConfig{
		Client: &http.Client{},
	}

	authImpl := auth.NewAuthImpl(config)

	token, err := authImpl.Login("mock-message", "mock-signature", "mock-wallet")
	assert.NoError(t, err)
	assert.Equal(t, "mock-token", token)

	// Error case
	httpmock.RegisterResponder("POST", common.AuthenticateAccount, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(400, common.Error{Error: "mock-error"})
	})

	token, err = authImpl.Login("mock-message", "mock-signature", "mock-wallet")
	assert.Error(t, err)
	assert.Equal(t, "mock-error", err.Error())
	assert.Empty(t, token)
}

// TestAuthImpl_GetMessage tests the GetMessage method
func TestAuthImpl_GetMessage(t *testing.T) {
	// Mock the HTTP client
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Success case
	httpmock.RegisterResponder("GET", common.GenerateSignMessage, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, common.MessageResponse{Message: "mock-message"})
	})

	config := common.SDKConfig{
		Client: &http.Client{},
	}

	authImpl := auth.NewAuthImpl(config)

	message, err := authImpl.GetMessage()
	assert.NoError(t, err)
	assert.Equal(t, "mock-message", message)

	// Error case
	httpmock.RegisterResponder("GET", common.GenerateSignMessage, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(400, common.Error{Error: "mock-error"})
	})

	message, err = authImpl.GetMessage()
	assert.Error(t, err)
	assert.Equal(t, "mock-error", err.Error())
	assert.Empty(t, message)
}

// TestAuthImpl_GetRefreshToken tests the GetRefreshToken method
func TestAuthImpl_GetRefreshToken(t *testing.T) {
	// Mock the HTTP client
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Success case
	httpmock.RegisterResponder("GET", common.RefreshToken, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, common.TokenResponse{Token: "mock-refresh-token"})
	})

	config := common.SDKConfig{
		Client: ,
	}

	authImpl := auth.NewAuthImpl(config)

	token, err := authImpl.GetRefreshToken()
	assert.NoError(t, err)
	assert.Equal(t, "mock-refresh-token", token)

	// Error case
	httpmock.RegisterResponder("GET", common.RefreshToken, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(400, common.Error{Error: "mock-error"})
	})

	token, err = authImpl.GetRefreshToken()
	assert.Error(t, err)
	assert.Equal(t, "mock-error", err.Error())
	assert.Empty(t, token)
}
