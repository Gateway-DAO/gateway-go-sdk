package dataassets_test

import (
	"net/http"
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
	dataassets "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_assets"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestACLSuite(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := common.SDKConfig{
		Client: client,
		ApiKey: "test-api-key",
	}

	aclImpl := dataassets.NewACLImpl(config)

	t.Run("TestAddACL", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 1, "roles": ["share"], "granted": true}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", common.AssignACLItemsToDataAsset, responder)

		// Test
		aclList := []common.ACLRequest{
			{Address: "test", Roles: []common.AccessLevel{
				common.RoleShare,
			}},
		}
		result, err := aclImpl.Add(1, aclList)

		// Assertions
		assert.NoError(t, err)
		assert.Contains(t, result.Roles[0], "share")
	})

	t.Run("TestAddACLError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Invalid ACL request"}`)
		httpmock.RegisterResponder("POST", common.AssignACLItemsToDataAsset, responder)

		// Test
		aclList := []common.ACLRequest{
			{Address: "test", Roles: []common.AccessLevel{
				common.RoleShare,
			}},
		}
		_, err := aclImpl.Add(1, aclList)

		// Assertions
		assert.Error(t, err)
	})

	t.Run("TestUpdateACL", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 1, "roles": ["share"], "granted": true}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", common.UpdateACLItemsToDataAsset, responder)

		// Test
		aclList := []common.ACLRequest{
			{Address: "test", Roles: []common.AccessLevel{
				common.RoleShare,
			}},
		}
		result, err := aclImpl.Update(1, aclList)

		// Assertions
		assert.NoError(t, err)
		assert.Contains(t, result.Roles[0], "share")
	})

	t.Run("TestUpdateACLError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Invalid ACL update"}`)
		httpmock.RegisterResponder("PUT", common.UpdateACLItemsToDataAsset, responder)

		// Test
		aclList := []common.ACLRequest{
			{Address: "test", Roles: []common.AccessLevel{
				common.RoleShare,
			}},
		}
		_, err := aclImpl.Update(1, aclList)

		// Assertions
		assert.Error(t, err)
	})

	t.Run("TestDeleteACL", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"message": "ACL deleted successfully"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("DELETE", common.DeleteAssignedRoleByACL, responder)

		// Test
		aclList := []common.ACLRequest{
			{Address: "test", Roles: []common.AccessLevel{
				common.RoleShare,
			}},
		}
		message, err := aclImpl.Delete(1, aclList)
		
		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "ACL deleted successfully", message)
	})

	t.Run("TestDeleteACLError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("DELETE", common.DeleteAssignedRoleByACL, responder)

		// Test
		aclList := []common.ACLRequest{
			{Address: "test", Roles: []common.AccessLevel{
				common.RoleShare,
			}},
		}
		message, err := aclImpl.Delete(1, aclList)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, message)
	})
}
