package examples

import (
	"fmt"
	"log"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

func ExampleGetAllDataModels(sdk *client.SDK) {
	page := 1
	pageSize := 10
	dataModels, err := sdk.DataModel.GetAll(page, pageSize)
	if err != nil {
		log.Fatalf("Error fetching data models: %v", err)
	}
	fmt.Printf("Fetched %d data models from GetAll\n", len(dataModels.Data))
}

func ExampleGetMyDataModels(sdk *client.SDK) {
	page := 1
	pageSize := 5
	dataModels, err := sdk.DataModel.GetMy(page, pageSize)
	if err != nil {
		log.Fatalf("Error fetching user's data models: %v", err)
	}
	fmt.Printf("Fetched %d user-specific data models from GetMy\n", len(dataModels.Data))
}

func ExampleGetByIDDataModel(sdk *client.SDK) {
	id := int64(1)
	dataModel, err := sdk.DataModel.GetById(id)
	if err != nil {
		log.Fatalf("Error fetching data model by ID: %v", err)
	}
	fmt.Printf("Fetched data model by ID: %v\n", dataModel)
}

func ExampleCreateDataModel(sdk *client.SDK) {
	dataModelInput := common.DataModelCreateRequest{
		Title:       "New Data Model",
		Description: "A description of the new data model",
	}
	createdModel, err := sdk.DataModel.Create(dataModelInput)
	if err != nil {
		log.Fatalf("Error creating data model: %v", err)
	}
	fmt.Printf("Created data model: %v\n", createdModel)
}

func ExampleUpdateDataModel(sdk *client.SDK) {
	id := int64(1)
	dataModelUpdate := common.DataModelUpdateRequest{
		Title:       "Updated Data Model Title",
		Description: "Updated description of the data model",
	}
	updatedModel, err := sdk.DataModel.Update(id, dataModelUpdate)
	if err != nil {
		log.Fatalf("Error updating data model: %v", err)
	}
	fmt.Printf("Updated data model: %v\n", updatedModel)
}

func RunDataModels() {
	sdk := client.NewSDK(client.SDKConfig{WalletDetails: client.WalletDetails{PrivateKey: "", WalletType: services.Ethereum}})

	ExampleCreateDataModel(sdk)
	ExampleUpdateDataModel(sdk)
	ExampleGetAllDataModels(sdk)
	ExampleGetMyDataModels(sdk)
	ExampleGetByIDDataModel(sdk)
}
