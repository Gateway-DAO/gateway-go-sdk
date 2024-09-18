package pkg

import (
	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	dataassets "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_assets"
)

type SDK struct {
	DataAssets dataassets.DataAssets
	APIKey     string
}

func NewSDK(apiKey string) *SDK {
	client := services.NewAPIClient(apiKey)
	return &SDK{
		DataAssets: dataassets.DataAssets{Client: client.Client, ApiKey: apiKey},
		APIKey:     apiKey,
	}
}
