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

func TestAccountsImpl(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := gateway.Config{
		Client: client,
	}

	accountImpl := gateway.NewAccountsImpl(config)

	t.Run("TestCreateAccount", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"token": "jwt-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.CreateAccount, responder)

		accountDetails := gateway.AccountCreateRequest{
			Signature:     "test",
			WalletAddress: "test",
			Message:       "test",
		}
		token, err := accountImpl.Create(accountDetails)

		assert.NoError(t, err)
		assert.Equal(t, "jwt-token", token)
	})

	t.Run("TestCreateAccountError", func(t *testing.T) {
		httpmock.Reset()

		errorResponse := `{"error": "Account creation failed"}`
		httpmock.RegisterResponder("POST", gateway.CreateAccount, httpmock.NewStringResponder(400, errorResponse))

		accountDetails := gateway.AccountCreateRequest{
			Signature:     "test",
			WalletAddress: "test",
			Message:       "test",
		}
		token, err := accountImpl.Create(accountDetails)

		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestCreateAccountHTTPRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("POST", gateway.CreateAccount, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		accountDetails := gateway.AccountCreateRequest{
			Signature:     "test",
			WalletAddress: "test",
			Message:       "test",
		}
		token, err := accountImpl.Create(accountDetails)

		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestGetMe", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"username": "testuser", "email": "test@example.com"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", gateway.GetMyAccount, responder)

		myAccount, err := accountImpl.GetMe()

		assert.NoError(t, err)
		assert.Equal(t, "testuser", myAccount.Username)
	})

	t.Run("TestGetMeError", func(t *testing.T) {
		httpmock.Reset()

		errorResponse := `{"error": "Failed to get account"}`
		httpmock.RegisterResponder("GET", gateway.GetMyAccount, httpmock.NewStringResponder(400, errorResponse))

		myAccount, err := accountImpl.GetMe()

		assert.Error(t, err)
		assert.Empty(t, myAccount)
	})

	t.Run("TestGetMeHTTPRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("POST", gateway.GetMyAccount, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		myAccount, err := accountImpl.GetMe()

		assert.Error(t, err)
		assert.Empty(t, myAccount)
	})

	t.Run("TestUpdateMe", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"username": "test", "email": "updated@example.com"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PATCH", gateway.GetMyAccount, responder)

		updateDetails := gateway.AccountUpdateRequest{
			ProfilePicture: "test",
		}
		myAccount, err := accountImpl.UpdateMe(updateDetails)

		assert.NoError(t, err)
		assert.Equal(t, "test", myAccount.Username)
	})

	t.Run("TestUpdateMeError", func(t *testing.T) {
		httpmock.Reset()

		errorResponse := `{"error": "Failed to update account"}`
		httpmock.RegisterResponder("PATCH", gateway.GetMyAccount, httpmock.NewStringResponder(400, errorResponse))

		updateDetails := gateway.AccountUpdateRequest{
			ProfilePicture: "test",
		}
		myAccount, err := accountImpl.UpdateMe(updateDetails)

		assert.Error(t, err)
		assert.Empty(t, myAccount)
	})

	t.Run("TestUpdateMeHTTPRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("PATCH", gateway.GetMyAccount, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		updateDetails := gateway.AccountUpdateRequest{
			ProfilePicture: "test",
		}
		myAccount, err := accountImpl.UpdateMe(updateDetails)

		assert.Error(t, err)
		assert.Empty(t, myAccount)
	})
}
