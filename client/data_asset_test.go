package client_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	dataassets "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_assets"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDataAssetSuite(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := SDKConfig{
		Client: client,
	}

	dataAssetImpl := dataassets.NewDataAssetImpl(config)

	t.Run("TestGetDataAsset", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 1, "title": "Test Asset", "description": "Test Description"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", "/data-assets/1", responder)

		// Test
		result, err := dataAssetImpl.Get(1)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 1, result.Id)
	})

	t.Run("TestGetDataAssetHttpRequestError", func(t *testing.T) {
		// Simulate an HTTP request error
		httpmock.RegisterResponder("GET", GetDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		// Call the function
		asset, err := dataAssetImpl.Get(123)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, asset)
	})

	t.Run("TestUploadDataAsset", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 1}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", CreateANewDataAsset, responder)

		// Test
		input := CreateDataAssetRequest{
			Name: "New Asset",
		}
		result, err := dataAssetImpl.Upload(input)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, int(1), result.Id)
	})

	t.Run("TestUploadDataAssetError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Invalid data asset request"}`)
		httpmock.RegisterResponder("POST", CreateANewDataAsset, responder)

		// Test
		input := CreateDataAssetRequest{
			Name: "New Asset",
		}
		_, err := dataAssetImpl.Upload(input)

		// Assertions
		assert.Error(t, err)
	})

	t.Run("TestUploadDataAssetHttpRequestError", func(t *testing.T) {
		// Simulate an HTTP request error
		httpmock.RegisterResponder("POST", CreateANewDataAsset,
			httpmock.NewErrorResponder(errors.New("http request error")))

		input := CreateDataAssetRequest{
			Name: "New Asset",
		}
		// Call the Upload method
		id, err := dataAssetImpl.Upload(input)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, id) // Assuming that id should be empty on error
	})

	t.Run("TestGetCreatedByMe", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"data": [{"id": 1, "title": "Created Asset"}]}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", GetCreatedDataAssets, responder)

		// Test
		result, err := dataAssetImpl.GetCreatedByMe(1, 10)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result.Data))
	})

	t.Run("TestGetCreatedByMeHttpRequestError", func(t *testing.T) {
		// Simulate an HTTP request error
		httpmock.RegisterResponder("GET", GetCreatedDataAssets,
			httpmock.NewErrorResponder(errors.New("http request error")))

		// Call the function
		assets, err := dataAssetImpl.GetCreatedByMe(1, 10)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestGetCreatedByMeServerError", func(t *testing.T) {
		// Setup the fixture for a server error response
		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("GET", GetCreatedDataAssets,
			httpmock.NewStringResponder(500, fixture))

		// Call the function
		assets, err := dataAssetImpl.GetCreatedByMe(1, 10)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestGetReceivedByMe", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"data": [{"id": 1, "title": "Received Asset"}]}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", GetReceivedDataAssets, responder)

		// Test
		result, err := dataAssetImpl.GetReceivedByMe(1, 10)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result.Data))
	})

	t.Run("TestGetReceivedByMeHttpRequestError", func(t *testing.T) {
		// Simulate an HTTP request error
		httpmock.RegisterResponder("GET", GetReceivedDataAssets,
			httpmock.NewErrorResponder(errors.New("http request error")))

		// Call the function
		assets, err := dataAssetImpl.GetReceivedByMe(1, 10)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestGetReceivedByMeServerError", func(t *testing.T) {
		// Setup the fixture for a server error response
		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("GET", GetReceivedDataAssets,
			httpmock.NewStringResponder(500, fixture))

		// Call the function
		assets, err := dataAssetImpl.GetReceivedByMe(1, 10)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, assets)
	})

	t.Run("TestUpdateAsset", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 1, "Name": "Updated Asset", "description": "Updated Description"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", "/data-assets/1", responder)

		// Test
		input := UpdateDataAssetRequest{
			Name: "New Asset",
		}
		result, err := dataAssetImpl.UpdateAsset("1", input)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "Updated Asset", result.Name)
	})

	t.Run("TestUpdateAssetApiError", func(t *testing.T) {
		httpmock.Reset()

		// Simulate an API error response
		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("PUT", UpdateDataAssetByID,
			httpmock.NewStringResponder(500, fixture))

		input := UpdateDataAssetRequest{
			Name: "New Asset",
		}
		// Call the UpdateAsset method
		asset, err := dataAssetImpl.UpdateAsset("1", input)

		// Assertions
		assert.Error(t, err)   // This should be true now
		assert.Empty(t, asset) // Assert that the asset is empty on error
	})

	t.Run("TestUpdateAssetHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		// Simulate an HTTP request error
		httpmock.RegisterResponder("PUT", UpdateDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		input := UpdateDataAssetRequest{
			Name: "New Asset",
		}
		// Call the UpdateAsset method
		asset, err := dataAssetImpl.UpdateAsset("1", input)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, asset) // Assuming that asset should be empty on error
	})

	t.Run("TestDeleteAsset", func(t *testing.T) {
		// Reset mock
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
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		httpmock.RegisterResponder("GET", DownloadDataAssetByID, httpmock.NewStringResponder(200, ""))

		fixture := "File content goes here" // Replace with actual file content if needed
		httpmock.RegisterResponder("GET", DownloadDataAssetByID, func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Disposition", `attachment; filename="testfile.txt"`)
			return resp, nil
		})

		// Test
		result, err := dataAssetImpl.Download(1)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "testfile.txt", result.FileName)
	})

	t.Run("TestShare", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `[{"id": 1, "roles": ["share"]}]`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", "/data-assets/1/share", responder)

		// Test
		shareDetails := []ShareDataAssetRequest{
			{Addresses: []string{"test"}},
		}
		result, err := dataAssetImpl.Share(1, shareDetails)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("TestShareHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		// Simulate an HTTP request error
		httpmock.RegisterResponder("POST", ShareDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		shareDetails := []ShareDataAssetRequest{
			{Addresses: []string{"test"}},
		}
		// Call the Share method
		acl, err := dataAssetImpl.Share(1, shareDetails)

		// Assertions
		assert.Error(t, err) // Expecting an error
		assert.Empty(t, acl) // Assert that the acl slice is empty on error
	})

	t.Run("TestShareApiError", func(t *testing.T) {
		httpmock.Reset()

		// Simulate an API error response
		fixture := `{"error": "Internal server error"}`
		httpmock.RegisterResponder("POST", ShareDataAssetByID,
			httpmock.NewStringResponder(500, fixture))

		shareDetails := []ShareDataAssetRequest{
			{Addresses: []string{"test"}},
		}
		// Call the Share method
		acl, err := dataAssetImpl.Share(1, shareDetails)

		// Assertions
		assert.Error(t, err) // Expecting an error
		assert.Empty(t, acl) // Assert that the acl slice is empty on error
	})

	t.Run("TestDeleteAssetApiError", func(t *testing.T) {
		httpmock.Reset()

		// Simulate an API error response
		fixture := `{"error": "Asset not found"}`
		httpmock.RegisterResponder("DELETE", DeleteDataAssetByID,
			httpmock.NewStringResponder(404, fixture))

		// Call the DeleteAsset method
		message, err := dataAssetImpl.DeleteAsset(1)

		// Assertions
		assert.Error(t, err)     // Expecting an error
		assert.Empty(t, message) // Assert that the message response is empty on error
	})

	t.Run("TestDeleteAssetHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		// Simulate an HTTP request error
		httpmock.RegisterResponder("DELETE", DeleteDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		// Call the DeleteAsset method
		message, err := dataAssetImpl.DeleteAsset(1)

		// Assertions
		assert.Error(t, err)     // Expecting an error
		assert.Empty(t, message) // Assert that the message response is empty on error
	})

	t.Run("TestUploadFileSuccess", func(t *testing.T) {
		// Setup
		httpmock.RegisterResponder("POST", CreateANewDataAsset,
			httpmock.NewJsonResponderOrPanic(200, DataAssetIDRequestAndResponse{Id: 123}))

		// Test
		expirationDate := time.Now().Add(24 * time.Hour)
		aclList := []ACLRequest{
			{Address: "test", Roles: []TypesAccessLevel{RoleShare}},
		}
		result, err := dataAssetImpl.UploadFile("testfile.txt", []byte("file content"), &aclList, &expirationDate)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 123, result.Id)
	})

	t.Run("TestUploadFileError", func(t *testing.T) {
		// Setup
		httpmock.RegisterResponder("POST", CreateANewDataAsset,
			httpmock.NewJsonResponderOrPanic(400, Error{Error: "Upload failed"}))

		// Test
		result, err := dataAssetImpl.UploadFile("testfile.txt", []byte("file content"), nil, nil)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})

	t.Run("TestUploadFileHttpRequestError", func(t *testing.T) {
		// Register an error responder to simulate HTTP request error
		httpmock.RegisterResponder("POST", CreateANewDataAsset,
			httpmock.NewErrorResponder(errors.New("http request error")))

		// Test
		result, err := dataAssetImpl.UploadFile("testfile.txt", []byte("file content"), nil, nil)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})

	t.Run("TestUpdateFileSuccess", func(t *testing.T) {
		// Setup
		httpmock.RegisterResponder("PUT", "/data-assets/123",
			httpmock.NewJsonResponderOrPanic(200, DataAssetIDRequestAndResponse{Id: 123}))

		fixture := `{"id":123}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", "/data-assets/123", responder)

		// Test
		expirationDate := time.Now().Add(24 * time.Hour)
		aclList := []ACLRequest{
			{Address: "test", Roles: []TypesAccessLevel{RoleShare}},
		}
		result, err := dataAssetImpl.UpdateFile("123", "testfile.txt", []byte("file content"), &aclList, &expirationDate)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 123, result.Id)
	})

	t.Run("TestUpdateFileError", func(t *testing.T) {
		// Setup
		httpmock.Reset()
		httpmock.RegisterResponder("PUT", UpdateDataAssetByID,
			httpmock.NewJsonResponderOrPanic(400, Error{Error: "Update failed"}))

		// Test
		result, err := dataAssetImpl.UpdateFile("123", "testfile.txt", []byte("new content"), nil, nil)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})

	t.Run("TestUpdateFileHttpRequestError", func(t *testing.T) {
		// Register an error responder to simulate HTTP request error
		httpmock.RegisterResponder("PUT", UpdateDataAssetByID,
			httpmock.NewErrorResponder(errors.New("http request error")))

		// Test
		result, err := dataAssetImpl.UpdateFile("123", "testfile.txt", []byte("new content"), nil, nil)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result.Id)
	})
}
