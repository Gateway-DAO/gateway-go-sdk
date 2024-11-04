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

func TestACLSuite(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := gateway.Config{
		Client: client,
	}

	aclImpl := gateway.NewACLImpl(config)

	t.Run("TestAddACL", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"id": 1, "roles": ["share"], "granted": true}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.AssignACLItemsToDataAsset, responder)

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		result, err := aclImpl.Add(1, aclList)

		assert.NoError(t, err)
		assert.Contains(t, result.Roles[0], "share")
	})

	t.Run("TestAddACLError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Invalid ACL request"}`)
		httpmock.RegisterResponder("POST", gateway.AssignACLItemsToDataAsset, responder)

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		_, err := aclImpl.Add(1, aclList)

		assert.Error(t, err)
	})

	t.Run("TestAddACLHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("POST", gateway.AssignACLItemsToDataAsset, httpmock.NewErrorResponder(errors.New("http request error")))

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		_, err := aclImpl.Add(1, aclList)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "http request error")
	})

	t.Run("TestUpdateACL", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"id": 1, "roles": ["share"], "granted": true}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", gateway.UpdateACLItemsToDataAsset, responder)

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		result, err := aclImpl.Update(1, aclList)

		assert.NoError(t, err)
		assert.Contains(t, result.Roles[0], "share")
	})

	t.Run("TestUpdateACLError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Invalid ACL update"}`)
		httpmock.RegisterResponder("PUT", gateway.UpdateACLItemsToDataAsset, responder)

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		_, err := aclImpl.Update(1, aclList)

		assert.Error(t, err)
	})

	t.Run("TestUpdateACLHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("PUT", gateway.UpdateACLItemsToDataAsset, httpmock.NewErrorResponder(errors.New("http request error")))

		// Test
		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		_, err := aclImpl.Update(1, aclList)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "http request error")
	})

	t.Run("TestDeleteACL", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"message": "ACL deleted successfully"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("DELETE", gateway.DeleteAssignedRoleByACL, responder)

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		message, err := aclImpl.Delete(1, aclList)

		assert.NoError(t, err)
		assert.Equal(t, "ACL deleted successfully", message)
	})

	t.Run("TestDeleteACLError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("DELETE", gateway.DeleteAssignedRoleByACL, responder)

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		message, err := aclImpl.Delete(1, aclList)

		assert.Error(t, err)
		assert.Empty(t, message)
	})

	t.Run("TestDeleteACLHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("DELETE", gateway.DeleteAssignedRoleByACL, httpmock.NewErrorResponder(errors.New("http request error")))

		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{
				gateway.RoleShare,
			}},
		}
		message, err := aclImpl.Delete(1, aclList)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "http request error")
		assert.Empty(t, message)
	})
}
