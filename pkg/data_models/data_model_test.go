package datamodels_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
	datamodels "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_models"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDataModelSuite(t *testing.T) {
	// Setup
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := common.SDKConfig{
		Client: client,
		ApiKey: "test-api-key",
	}

	dataModelImpl := datamodels.NewDataModelImpl(config)

	t.Run("TestGetAll", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `
		{
			"meta": {
				"total_items": 2,
				"page": 1,
				"page_size": 10
			},
			"data": [
				{
					"created_by": "user1",
					"id": 1,
					"created_at": "2023-10-01T00:00:00Z",
					"description": "First Data Model",
					"schema": {
						"field1": "value1"
					},
					"tags": ["tag1"],
					"title": "Data Model 1",
					"updated_at": "2023-10-05T00:00:00Z",
					"deleted_at": null
				},
				{
					"created_by": "user2",
					"id": 2,
					"created_at": "2023-10-02T00:00:00Z",
					"description": "Second Data Model",
					"schema": {
						"field2": "value2"
					},
					"tags": ["tag2"],
					"title": "Data Model 2",
					"updated_at": "2023-10-06T00:00:00Z",
					"deleted_at": null
				}
			],
			"links": {
				"next": "/next-page"
			}
		}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", common.GetDataModels, responder)

		// Test
		page, pageSize := 1, 10
		response, err := dataModelImpl.GetAll(page, pageSize)
		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 2, response.Meta.TotalItems)
		assert.Len(t, response.Data, 2)
		assert.Equal(t, "Data Model 1", response.Data[0].Title)
		assert.Equal(t, "Data Model 2", response.Data[1].Title)
		assert.Equal(t, "user1", response.Data[0].CreatedBy)
		assert.Equal(t, "user2", response.Data[1].CreatedBy)

	})

	t.Run("TestGetALLError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("GET", common.GetDataModels, responder)

		page, pageSize := 1, 10
		result, err := dataModelImpl.GetAll(page, pageSize)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetAllHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Simulate a client-side error (e.g., network error)
		httpmock.RegisterResponder("GET", common.GetDataModels, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		// Test
		page, pageSize := 1, 10
		response, err := dataModelImpl.GetAll(page, pageSize)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("TestGetMy", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `
		{
			"meta": {
				"total_items": 2,
				"page": 1,
				"page_size": 10
			},
			"data": [
				{
					"created_by": "user1",
					"id": 1,
					"created_at": "2023-10-01T00:00:00Z",
					"description": "First Data Model",
					"schema": {
						"field1": "value1"
					},
					"tags": ["tag1"],
					"title": "Data Model 1",
					"updated_at": "2023-10-05T00:00:00Z",
					"deleted_at": null
				},
				{
					"created_by": "user2",
					"id": 2,
					"created_at": "2023-10-02T00:00:00Z",
					"description": "Second Data Model",
					"schema": {
						"field2": "value2"
					},
					"tags": ["tag2"],
					"title": "Data Model 2",
					"updated_at": "2023-10-06T00:00:00Z",
					"deleted_at": null
				}
			],
			"links": {
				"next": "/next-page"
			}
		}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("GET", common.GetDataModelsByUser, responder)

		// Test
		page, pageSize := 1, 10
		result, err := dataModelImpl.GetMy(page, pageSize)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 2, result.Meta.TotalItems)
	})

	t.Run("TestGetMyIDError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("GET", common.GetDataModelsByUser, responder)

		page, pageSize := 1, 10
		result, err := dataModelImpl.GetMy(page, pageSize)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetMyHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Simulate a client-side error
		httpmock.RegisterResponder("GET", common.GetDataModelsByUser, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		// Test
		page, pageSize := 1, 10
		result, err := dataModelImpl.GetMy(page, pageSize)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetById", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"created_by": "user1",
					"id": 1,
					"created_at": "2023-10-01T00:00:00Z",
					"description": "First Data Model",
					"schema": {
						"field1": "value1"
					},
					"tags": ["tag1"],
					"title": "Data Model 1",
					"updated_at": "2023-10-05T00:00:00Z",
					"deleted_at": null}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}

		// Register responder for dynamic ID
		httpmock.RegisterResponder("GET", "/data-models/1", responder)

		// Test
		id := int64(1)
		result, err := dataModelImpl.GetById(id)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, int(1), result.Id)
		assert.Equal(t, "user1", result.CreatedBy)
		assert.Equal(t, "Data Model 1", result.Title)
	})

	t.Run("TestGetByIDError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("GET", "/data-models/1", responder)

		id := int64(1)
		result, err := dataModelImpl.GetById(id)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetByIdHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Simulate a client-side error
		httpmock.RegisterResponder("GET", "/data-models/1", func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		// Test
		id := int64(1)
		result, err := dataModelImpl.GetById(id)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestCreate", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 2, "name": "NewModel"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", common.CreateDataModel, responder)

		// Test
		dataModelInput := common.DataModelRequest{
			Title: "NewModel",
		}
		result, err := dataModelImpl.Create(dataModelInput)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, int(2), result.Id)
	})

	t.Run("TestCreateError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("POST", common.CreateDataModel, responder)

		dataModelInput := common.DataModelRequest{
			Title: "NewModel",
		}
		result, err := dataModelImpl.Create(dataModelInput)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestCreateHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Simulate a client-side error
		httpmock.RegisterResponder("POST", common.CreateDataModel, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		// Test
		dataModelInput := common.DataModelRequest{
			Title: "NewModel",
		}
		result, err := dataModelImpl.Create(dataModelInput)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestUpdate", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 2, "name": "UpdatedModel"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", common.UpdateDataModel, responder)

		// Test
		id := int64(2)
		dataModelInput := common.DataModelRequest{
			Title: "UpdatedModel",
		}
		result, err := dataModelImpl.Update(id, dataModelInput)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, int(2), result.Id)
	})

	t.Run("TestUpdateError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Set up mock response
		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("PUT", common.UpdateDataModel, responder)

		id := int64(2)
		dataModelInput := common.DataModelRequest{
			Title: "UpdatedModel",
		}
		result, err := dataModelImpl.Update(id, dataModelInput)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestUpdateHttpRequestError", func(t *testing.T) {
		// Reset mock
		httpmock.Reset()

		// Simulate a client-side error
		httpmock.RegisterResponder("PUT", common.UpdateDataModel, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		// Test
		id := int64(2)
		dataModelInput := common.DataModelRequest{
			Title: "UpdatedModel",
		}
		result, err := dataModelImpl.Update(id, dataModelInput)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
