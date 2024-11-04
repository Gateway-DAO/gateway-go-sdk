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

func TestDataModelSuite(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	config := gateway.Config{
		Client: client,
	}

	dataModelImpl := gateway.NewDataModelImpl(config)

	t.Run("TestGetAll", func(t *testing.T) {
		httpmock.Reset()

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
		httpmock.RegisterResponder("GET", gateway.GetDataModels, responder)

		page, pageSize := 1, 10
		response, err := dataModelImpl.GetAll(page, pageSize)
		assert.NoError(t, err)
		assert.Equal(t, 2, response.Meta.TotalItems)
		assert.Len(t, response.Data, 2)
		assert.Equal(t, "Data Model 1", response.Data[0].Title)
		assert.Equal(t, "Data Model 2", response.Data[1].Title)
	})

	t.Run("TestGetALLError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("GET", gateway.GetDataModels, responder)

		page, pageSize := 1, 10
		result, err := dataModelImpl.GetAll(page, pageSize)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetAllHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("GET", gateway.GetDataModels, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		page, pageSize := 1, 10
		response, err := dataModelImpl.GetAll(page, pageSize)

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("TestGetMy", func(t *testing.T) {
		httpmock.Reset()

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
		httpmock.RegisterResponder("GET", gateway.GetDataModelsByUser, responder)

		page, pageSize := 1, 10
		result, err := dataModelImpl.GetMy(page, pageSize)

		assert.NoError(t, err)
		assert.Equal(t, 2, result.Meta.TotalItems)
	})

	t.Run("TestGetMyIDError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("GET", gateway.GetDataModelsByUser, responder)

		page, pageSize := 1, 10
		result, err := dataModelImpl.GetMy(page, pageSize)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetMyHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("GET", gateway.GetDataModelsByUser, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		page, pageSize := 1, 10
		result, err := dataModelImpl.GetMy(page, pageSize)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetById", func(t *testing.T) {
		httpmock.Reset()

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

		httpmock.RegisterResponder("GET", "/data-models/1", responder)

		id := int64(1)
		result, err := dataModelImpl.GetById(id)

		assert.NoError(t, err)
		assert.Equal(t, int(1), result.Id)
		assert.Equal(t, "Data Model 1", result.Title)
	})

	t.Run("TestGetByIDError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("GET", "/data-models/1", responder)

		id := int64(1)
		result, err := dataModelImpl.GetById(id)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestGetByIdHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("GET", "/data-models/1", func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		id := int64(1)
		result, err := dataModelImpl.GetById(id)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestCreate", func(t *testing.T) {
		httpmock.Reset()

		fixture := `{"id": 2, "name": "NewModel"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("POST", gateway.CreateDataModel, responder)

		dataModelInput := gateway.DataModelCreateRequest{
			Title: "NewModel",
		}
		result, err := dataModelImpl.Create(dataModelInput)

		assert.NoError(t, err)
		assert.Equal(t, int(2), result.Id)
	})

	t.Run("TestCreateError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("POST", gateway.CreateDataModel, responder)

		dataModelInput := gateway.DataModelCreateRequest{
			Title: "NewModel",
		}
		result, err := dataModelImpl.Create(dataModelInput)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestCreateHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("POST", gateway.CreateDataModel, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		dataModelInput := gateway.DataModelCreateRequest{
			Title: "NewModel",
		}
		result, err := dataModelImpl.Create(dataModelInput)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestUpdate", func(t *testing.T) {
		httpmock.Reset()

		// Set up mock response
		fixture := `{"id": 2, "name": "UpdatedModel"}`
		responder := func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fixture)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		httpmock.RegisterResponder("PUT", gateway.UpdateDataModel, responder)

		id := int64(2)
		dataModelInput := gateway.DataModelUpdateRequest{
			Title: "UpdatedModel",
		}
		result, err := dataModelImpl.Update(id, dataModelInput)

		assert.NoError(t, err)
		assert.Equal(t, int(2), result.Id)
	})

	t.Run("TestUpdateError", func(t *testing.T) {
		httpmock.Reset()

		responder := httpmock.NewStringResponder(400, `{"error": "Failed to delete ACL"}`)
		httpmock.RegisterResponder("PUT", gateway.UpdateDataModel, responder)

		id := int64(2)
		dataModelInput := gateway.DataModelUpdateRequest{
			Title: "UpdatedModel",
		}
		result, err := dataModelImpl.Update(id, dataModelInput)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("TestUpdateHttpRequestError", func(t *testing.T) {
		httpmock.Reset()

		httpmock.RegisterResponder("PUT", gateway.UpdateDataModel, func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("client-side error")
		})

		id := int64(2)
		dataModelInput := gateway.DataModelUpdateRequest{
			Title: "UpdatedModel",
		}
		result, err := dataModelImpl.Update(id, dataModelInput)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
