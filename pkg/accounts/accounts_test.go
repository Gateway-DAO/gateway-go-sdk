package accounts_test

import (
	"net/http"
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/accounts"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAccountsImpl(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := common.SDKConfig{
		Client: client,
		ApiKey: "test-api-key",
	}

	accountImpl := accounts.NewAccountsImpl(config)

	t.Run("TestCreateAccount", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"token": "jwt-token"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", common.CreateAccount, responder)

		// Test
		accountDetails := common.AccountCreateRequest{
			Signature:     "test",
			WalletAddress: "test",
			Message:       "test",
		}
		token, err := accountImpl.Create(accountDetails)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "jwt-token", token)
	})

	t.Run("TestCreateAccountError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response for error
		errorResponse := `{"error": "Account creation failed"}`
		httpmock.RegisterResponder("POST", common.CreateAccount, httpmock.NewStringResponder(400, errorResponse))

		// Test
		accountDetails := common.AccountCreateRequest{
			Signature:     "test",
			WalletAddress: "test",
			Message:       "test",
		}
		token, err := accountImpl.Create(accountDetails)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("TestGetMe", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"username": "testuser", "email": "test@example.com"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", common.GetMyAccount, responder)

		// Test
		myAccount, err := accountImpl.GetMe()

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "testuser", myAccount.Username)
	})

	t.Run("TestGetMeError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response for error
		errorResponse := `{"error": "Failed to get account"}`
		httpmock.RegisterResponder("GET", common.GetMyAccount, httpmock.NewStringResponder(400, errorResponse))

		// Test
		myAccount, err := accountImpl.GetMe()

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, myAccount)
	})

	t.Run("TestUpdateMe", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"username": "test", "email": "updated@example.com"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PATCH", common.GetMyAccount, responder)

		// Test
		updateDetails := common.AccountUpdateRequest{
			ProfilePicture: "test",
		}
		myAccount, err := accountImpl.UpdateMe(updateDetails)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "test", myAccount.Username)
	})

	t.Run("TestUpdateMeError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response for error
		errorResponse := `{"error": "Failed to update account"}`
		httpmock.RegisterResponder("PATCH", common.GetMyAccount, httpmock.NewStringResponder(400, errorResponse))

		// Test
		updateDetails := common.AccountUpdateRequest{
			ProfilePicture: "test",
		}
		myAccount, err := accountImpl.UpdateMe(updateDetails)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, myAccount)
	})
}
