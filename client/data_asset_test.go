package client_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	gateway "github.com/Gateway-DAO/gateway-go-sdk/client"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDataAssetSuite(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := gateway.Config{
		Client: client,
	}

	dataAssetImpl := gateway.NewDataAssetImpl(config)

	t.Run("TestGetDataAsset", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"id": 1, "title": "Test Asset", "description": "Test Description"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", "/data-assets/1", responder)

		result, err := dataAssetImpl.Get(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, result.Id)
	})

	t.Run("TestGetDataAssetHttpRequestError", func(t *testing.T) {
		httpmock.RegisterResponder("GET", gateway.GetDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		asset, err := dataAssetImpl.Get(123)

		assert.Error(t, err)
		assert.Empty(t, asset)
	})

	t.Run("TestUploadDataAsset", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"id": 1}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.CreateANewDataAsset, responder)

		input := gateway.CreateDataAssetRequest{
			Name: "New Asset",
		}
		result, err := dataAssetImpl.Upload(input)

		assert.NoError(t, err)
		assert.Equal(t, int(1), result.Id)
	})

	t.Run("TestUploadDataAssetError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Invalid data asset request"}`)
		httpmock.RegisterResponder("POST", gateway.CreateANewDataAsset, responder)

		input := gateway.CreateDataAssetRequest{
			Name: "New Asset",
		}
		_, err := dataAssetImpl.Upload(input)

		assert.Error(t, err)
	})

	t.Run("TestUploadDataAssetHttpRequestError", func(t *testing.T) {
		httpmock.RegisterResponder("POST", gateway.CreateANewDataAsset,
			httpmock.NewErrorResponder(errors.New("http request error")))

		input := gateway.CreateDataAssetRequest{
			Name: "New Asset",
		}
		id, err := dataAssetImpl.Upload(input)

		assert.Error(t, err)
		assert.Empty(t, id)
	})

	t.Run("TestGetCreatedByMe", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"data": [{"id": 1, "title": "Created Asset"}]}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", gateway.GetCreatedDataAssets, responder)

		result, err := dataAssetImpl.GetCreatedByMe(1, 10)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(result.Data))
	})

	t.Run("TestGetCreatedByMeHttpRequestError", func(t *testing.T) {
		httpmock.RegisterResponder("GET", gateway.GetCreatedDataAssets,
			httpmock.NewErrorResponder(errors.New("http request error")))

		assets, err := dataAssetImpl.GetCreatedByMe(1, 10)

		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestGetCreatedByMeServerError", func(t *testing.T) {
		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("GET", gateway.GetCreatedDataAssets,
			httpmock.NewStringResponder(500, fixture))

		assets, err := dataAssetImpl.GetCreatedByMe(1, 10)

		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestGetReceivedByMe", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"data": [{"id": 1, "title": "Received Asset"}]}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", gateway.GetReceivedDataAssets, responder)

		result, err := dataAssetImpl.GetReceivedByMe(1, 10)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(result.Data))
	})

	t.Run("TestGetReceivedByMeHttpRequestError", func(t *testing.T) {
		httpmock.RegisterResponder("GET", gateway.GetReceivedDataAssets,
			httpmock.NewErrorResponder(errors.New("http request error")))

		assets, err := dataAssetImpl.GetReceivedByMe(1, 10)

		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestGetReceivedByMeServerError", func(t *testing.T) {
		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("GET", gateway.GetReceivedDataAssets,
			httpmock.NewStringResponder(500, fixture))

		assets, err := dataAssetImpl.GetReceivedByMe(1, 10)

		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestUpdateAsset", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"id": 1, "Name": "Updated Asset", "description": "Updated Description"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", "/data-assets/1", responder)

		input := gateway.UpdateDataAssetRequest{
			Name: "New Asset",
		}
		result, err := dataAssetImpl.UpdateAsset("1", input)

		assert.NoError(t, err)
		assert.Equal(t, "Updated Asset", result.Name)
	})

	t.Run("TestUpdateAssetApiError", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("PUT", gateway.UpdateDataAssetByID,
			httpmock.NewStringResponder(500, fixture))

		input := gateway.UpdateDataAssetRequest{
			Name: "New Asset",
		}
		asset, err := dataAssetImpl.UpdateAsset("1", input)

		assert.Error(t, err)
		assert.Empty(t, asset)
	})

	t.Run("TestUpdateAssetHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("PUT", gateway.UpdateDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		input := gateway.UpdateDataAssetRequest{
			Name: "New Asset",
		}
		asset, err := dataAssetImpl.UpdateAsset("1", input)

		assert.Error(t, err)
		assert.Empty(t, asset)
	})

	t.Run("TestDeleteAsset", func(t *testing.T) {
		httpmock.Reset()

		// Set up mock response
		fixture := `{"message": "Asset deleted successfully"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("DELETE", "/data-assets/1", responder)

		// Test
		message, err := dataAssetImpl.DeleteAsset(1)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "Asset deleted successfully", message.Message)
	})

	t.Run("TestDownload", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("GET", gateway.DownloadDataAssetByID, httpmock.NewStringResponder(200, ""))

		fixture := "File content goes here"
		httpmock.RegisterResponder("GET", gateway.DownloadDataAssetByID, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Disposition", `attachment; filename="testfile.txt"`)
			return resp, nil
		})

		result, err := dataAssetImpl.Download(1)

		assert.NoError(t, err)
		assert.Equal(t, "testfile.txt", result.FileName)
	})

	t.Run("TestShare", func(t *testing.T) {
		httpmock.Reset()

		fixture := `[{"id": 1, "roles": ["share"]}]`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", "/data-assets/1/share", responder)

		shareDetails := []gateway.ShareDataAssetRequest{
			{Addresses: []string{"test"}},
		}
		result, err := dataAssetImpl.Share(1, shareDetails)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("TestShareHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("POST", gateway.ShareDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		shareDetails := []gateway.ShareDataAssetRequest{
			{Addresses: []string{"test"}},
		}
		acl, err := dataAssetImpl.Share(1, shareDetails)

		assert.Error(t, err)
		assert.Empty(t, acl)
	})

	t.Run("TestShareApiError", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("POST", gateway.ShareDataAssetByID,
			httpmock.NewStringResponder(500, fixture))

		shareDetails := []gateway.ShareDataAssetRequest{
			{Addresses: []string{"test"}},
		}

		acl, err := dataAssetImpl.Share(1, shareDetails)

		assert.Error(t, err)
		assert.Empty(t, acl)
	})

	t.Run("TestDeleteAssetApiError", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"error": "Asset not found"}`
		httpmock.RegisterResponder("DELETE", gateway.DeleteDataAssetByID,
			httpmock.NewStringResponder(404, fixture))

		message, err := dataAssetImpl.DeleteAsset(1)

		assert.Error(t, err)
		assert.Empty(t, message)
	})

	t.Run("TestDeleteAssetHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("DELETE", gateway.DeleteDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		message, err := dataAssetImpl.DeleteAsset(1)

		assert.Error(t, err)     
		assert.Empty(t, message) 
	})

	t.Run("TestUploadFileSuccess", func(t *testing.T) {
		httpmock.RegisterResponder("POST", gateway.CreateANewDataAsset,
			httpmock.NewJsonResponderOrPanic(200, gateway.DataAssetIDRequestAndResponse{Id: 123}))

		expirationDate := time.Now().Add(24 * time.Hour)
		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{gateway.RoleShare}},
		}
		result, err := dataAssetImpl.UploadFile("testfile.txt", []byte("file content"), &aclList, &expirationDate)

		assert.NoError(t, err)
		assert.Equal(t, 123, result.Id)
	})

	t.Run("TestUploadFileError", func(t *testing.T) {
		httpmock.RegisterResponder("POST", gateway.CreateANewDataAsset,
			httpmock.NewJsonResponderOrPanic(400, gateway.Error{Error: "Upload failed"}))

		result, err := dataAssetImpl.UploadFile("testfile.txt", []byte("file content"), nil, nil)

		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})

	t.Run("TestUploadFileHttpRequestError", func(t *testing.T) {
		httpmock.RegisterResponder("POST", gateway.CreateANewDataAsset,
			httpmock.NewErrorResponder(errors.New("http request error")))

		result, err := dataAssetImpl.UploadFile("testfile.txt", []byte("file content"), nil, nil)

		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})

	t.Run("TestUpdateFileSuccess", func(t *testing.T) {
		httpmock.RegisterResponder("PUT", "/data-assets/123",
			httpmock.NewJsonResponderOrPanic(200, gateway.DataAssetIDRequestAndResponse{Id: 123}))

		fixture := `{"id":123}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", "/data-assets/123", responder)

		expirationDate := time.Now().Add(24 * time.Hour)
		aclList := []gateway.ACLRequest{
			{Address: "test", Roles: []gateway.TypesAccessLevel{gateway.RoleShare}},
		}
		result, err := dataAssetImpl.UpdateFile("123", "testfile.txt", []byte("file content"), &aclList, &expirationDate)

		assert.NoError(t, err)
		assert.Equal(t, 123, result.Id)
	})

	t.Run("TestUpdateFileError", func(t *testing.T) {
		httpmock.Reset()
		httpmock.RegisterResponder("PUT", gateway.UpdateDataAssetByID,
			httpmock.NewJsonResponderOrPanic(400, gateway.Error{Error: "Update failed"}))

		result, err := dataAssetImpl.UpdateFile("123", "testfile.txt", []byte("new content"), nil, nil)

		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})

	t.Run("TestUpdateFileHttpRequestError", func(t *testing.T) {
		httpmock.RegisterResponder("PUT", gateway.UpdateDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		result, err := dataAssetImpl.UpdateFile("123", "testfile.txt", []byte("new content"), nil, nil)

		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})
}
