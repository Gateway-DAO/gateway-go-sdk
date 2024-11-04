package client_test

import (
	"errors"
	"net/http"
	"testing"

	gateway "github.com/Gateway-DAO/gateway-go-sdk/client"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthSuite(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := gateway.Config{
		Client: client,
	}

	authImpl := gateway.NewAuthImpl(config)

	t.Run("TestLogin", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.AuthenticateAccount, responder)

		token, err := authImpl.Login("test", "0x1bd2387faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

		assert.NoError(t, err)
		assert.Equal(t, "test-token", token)
	})

	t.Run("TestLoginEtherr", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.AuthenticateAccount, responder)

		_, err := authImpl.Login("test", "0x1bd287faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

		assert.Error(t, err)
	})

	t.Run("TestLoginSol", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.AuthenticateAccount, responder)

		token, err := authImpl.Login("test", "4mWRM4VmpkJ4uSCAYTFjYuCZRLD5BGs7gwyrRckCGnF7HZkjdUXJRX7bBXnTpuyiWi2BYvPaTQB9QGHFw9jeTdwT", "AqzrrxaBCXRsq2BaY32djAp38B42asRRahbsYvD5uvSF")

		assert.NoError(t, err)
		assert.Equal(t, "test-token", token)
	})

	t.Run("TestLoginSolError", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"token": "test-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.AuthenticateAccount, responder)

		_, err := authImpl.Login("test", "mWRM4VmpkJ4uSCAYTFjYuCZRLD5BGs7gwyrRckCGnF7HZkjdUXJRX7bBXnTpuyiWi2BYvPaTQB9QGHFw9jeTdwT", "AqzrrxaBCXRsq2BaY32djAp38B42asRRahbsYvD5uvSF")

		assert.Error(t, err)
	})

	t.Run("TestLoginError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Invalid credentials"}`)
		httpmock.RegisterResponder("POST", gateway.AuthenticateAccount, responder)

		token, err := authImpl.Login("test", "0x1bd2387faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestLoginHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("POST", gateway.AuthenticateAccount, httpmock.NewErrorResponder(errors.New("http request error")))

		token, err := authImpl.Login("test", "0x1bd2387faf757527cd96d0461bd3012fec227c0b85045169b3e2d4fbc8b9a2c55580db184c33a3810404aa1787151e28e647847a8dfd4d3195c64749494d18421b", "0x125b968F9ac42F33b0e1f1FBEbeE016Ca24A7116")

		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestGetMessage", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"message": "test-message"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", gateway.GenerateSignMessage, responder)

		message, err := authImpl.GetMessage()

		assert.NoError(t, err)
		assert.Equal(t, "test-message", message)
	})

	t.Run("TestGetMessageError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Invalid credentials"}`)
		httpmock.RegisterResponder("GET", gateway.GenerateSignMessage, responder)

		message, err := authImpl.GetMessage()

		assert.Error(t, err)
		assert.Empty(t, message)
	})

	t.Run("TestGetMessageHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("GET", gateway.GenerateSignMessage, httpmock.NewErrorResponder(errors.New("http request error")))

		message, err := authImpl.GetMessage()

		assert.Error(t, err)
		assert.Empty(t, message)
	})

	t.Run("TestGetRefreshToken", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"token": "refresh-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", gateway.RefreshToken, responder)

		token, err := authImpl.GetRefreshToken()

		assert.NoError(t, err)
		assert.Equal(t, "refresh-token", token)
	})

	t.Run("TestGetRefreshTokenError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Invalid credentials"}`)
		httpmock.RegisterResponder("GET", gateway.RefreshToken, responder)

		token, err := authImpl.GetRefreshToken()

		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestGetRefreshTokenHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("GET", gateway.RefreshToken, httpmock.NewErrorResponder(errors.New("http request error")))

		token, err := authImpl.GetRefreshToken()

		assert.Error(t, err)
		assert.Empty(t, token)
	})
}
