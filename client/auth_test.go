package client_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/client/auth"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthSuite(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := SDKConfig{
		Client: client,
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
		httpmock.RegisterResponder("POST", AuthenticateAccount, responder)

		// Test
		token, err := authImpl.Login("test", "0x1bd2387faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "test-token", token)
	})

	t.Run("TestLoginEtherr", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", AuthenticateAccount, responder)

		// Test
		_, err := authImpl.Login("test", "0x1bd287faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

		// Assertions
		assert.Error(t, err)
	})

	t.Run("TestLoginSol", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", AuthenticateAccount, responder)

		// Test
		token, err := authImpl.Login("test", "4mWRM4VmpkJ4uSCAYTFjYuCZRLD5BGs7gwyrRckCGnF7HZkjdUXJRX7bBXnTpuyiWi2BYvPaTQB9QGHFw9jeTdwT", "AqzrrxaBCXRsq2BaY32djAp38B42asRRahbsYvD5uvSF")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "test-token", token)
	})

	t.Run("TestLoginSolError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", AuthenticateAccount, responder)

		// Test
		_, err := authImpl.Login("test", "mWRM4VmpkJ4uSCAYTFjYuCZRLD5BGs7gwyrRckCGnF7HZkjdUXJRX7bBXnTpuyiWi2BYvPaTQB9QGHFw9jeTdwT", "AqzrrxaBCXRsq2BaY32djAp38B42asRRahbsYvD5uvSF")

		// Assertions
		assert.Error(t, err)
	})

	t.Run("TestLoginError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Invalid credentials"}`)
		httpmock.RegisterResponder("POST", AuthenticateAccount, responder)

		// Test
		token, err := authImpl.Login("test", "0x1bd2387faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestLoginHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Register an error responder to simulate HTTP request error
		httpmock.RegisterResponder("POST", AuthenticateAccount, httpmock.NewErrorResponder(errors.New("http request error")))

		// Test
		token, err := authImpl.Login("test", "0x1bd2387faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

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
		httpmock.RegisterResponder("GET", GenerateSignMessage, responder)

		// Test
		message, err := authImpl.GetMessage()

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "test-message", message)
	})

	t.Run("TestGetMessageError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Invalid credentials"}`)
		httpmock.RegisterResponder("GET", GenerateSignMessage, responder)

		// Test
		message, err := authImpl.GetMessage()

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, message)
	})

	t.Run("TestGetMessageHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Register an error responder to simulate HTTP request error
		httpmock.RegisterResponder("GET", GenerateSignMessage, httpmock.NewErrorResponder(errors.New("http request error")))

		// Test
		message, err := authImpl.GetMessage()

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, message)
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
		httpmock.RegisterResponder("GET", RefreshToken, responder)

		// Test
		token, err := authImpl.GetRefreshToken()

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "refresh-token", token)
	})

	t.Run("TestGetRefreshTokenError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Invalid credentials"}`)
		httpmock.RegisterResponder("GET", RefreshToken, responder)

		// Test
		token, err := authImpl.GetRefreshToken()

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestGetRefreshTokenHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Register an error responder to simulate HTTP request error
		httpmock.RegisterResponder("GET", RefreshToken, httpmock.NewErrorResponder(errors.New("http request error")))

		// Test
		token, err := authImpl.GetRefreshToken()

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, token)
	})
} // Closing bracket for TestAuthSuite
