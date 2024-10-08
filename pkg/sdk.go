package pkg

import (
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/accounts"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/auth"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
	dataassets "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_assets"
	datamodels "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_models"
	"github.com/go-resty/resty/v2"
)

type SDK struct {
	DataAssets dataassets.DataAsset
	DataModel  datamodels.DataModel
	Account    *accounts.AccountsImpl
	Auth       auth.Auth
	APIKey     string
}

func NewSDK(apiKey string) *SDK {
	client := resty.New()
	client.SetBaseURL("https://dev.api.gateway.tech")
	client.SetAuthToken(apiKey)
	// client.SetDebug(true)

	config := common.SDKConfig{
		Client: client,
		ApiKey: apiKey,
	}

	return &SDK{
		DataAssets: dataassets.NewDataAssetImpl(config),
		DataModel:  datamodels.NewDataModelImpl(config),
		Auth:       auth.NewAuthImpl(config),
		Account:    accounts.NewAccountsImpl(config),
		APIKey:     apiKey,
	}
}
