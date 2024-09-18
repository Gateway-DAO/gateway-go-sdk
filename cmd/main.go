package main

import (
	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	dataassets "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_assets"
)

type SDK struct {
	DataAssets *dataassets.DataAssets
	Client     *services.APIClient
	APIKey     string
}

func NewSDK(apiKey string) *SDK {
	client := services.NewAPIClient(apiKey)
	return &SDK{
		DataAssets: dataassets.New(client.Client, apiKey),
		Client:     client,
		APIKey:     apiKey,
	}
}
