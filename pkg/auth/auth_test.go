package auth_test

import (
	"net/http"
	"testing"

	"gateway-go-sdk/pkg/auth"
	"gateway-go-sdk/pkg/common"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthSuite(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := common.SDKConfig{
		Client: client,
		ApiKey: "test-api-key",
	}

	authImpl := auth.NewAuthImpl(config)

	t.Run("TestLogin", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", common.AuthenticateAccount, responder)

		// Test
		token, err := authImpl.Login("test-message", "test-signature", "test-wallet")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "test-token", token)
	})

	t.Run("TestLoginError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Invalid credentials"}`)
		httpmock.RegisterResponder("POST", common.AuthenticateAccount, responder)

		// Test
		token, err := authImpl.Login("test-message", "test-signature", "test-wallet")

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestGetMessage", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"message": "test-message"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", common.GenerateSignMessage, responder)

		// Test
		message, err := authImpl.GetMessage()

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "test-message", message)
	})

	t.Run("TestGetRefreshToken", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"token": "refresh-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", common.RefreshToken, responder)

		// Test
		token, err := authImpl.GetRefreshToken()

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "refresh-token", token)
	})
}
