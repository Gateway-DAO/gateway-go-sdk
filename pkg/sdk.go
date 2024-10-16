package pkg

import (
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/accounts"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/auth"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
	dataassets "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_assets"
	datamodels "github.com/Gateway-DAO/gateway-go-sdk/pkg/data_models"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/helpers"

	"github.com/go-resty/resty/v2"
)

type SDK struct {
	DataAssets dataassets.DataAsset
	DataModel  datamodels.DataModel
	Account    *accounts.AccountsImpl
	ACL        dataassets.ACL
	Auth       auth.Auth
	APIKey     string
}

type SDKConfig struct {
	ApiKey        string
	WalletDetails WalletDetails
	URL           string
}

type WalletDetails struct {
	PrivateKey string
	WalletType common.WalletTypeEnum
}

func NewSDK(config SDKConfig) *SDK {
	client := resty.New()
	client.OnBeforeRequest(helpers.Middleware)
	if config.URL != "" {
		client.SetBaseURL(config.URL)
	} else {
		client.SetBaseURL("https://dev.api.gateway.tech")
	}

	if config.ApiKey != "" {
		client.SetAuthToken(config.ApiKey)
	}

	sdkClient := common.SDKConfig{
		Client: client,
		ApiKey: "",
	}

	return &SDK{
		DataAssets: dataassets.NewDataAssetImpl(sdkClient),
		DataModel:  datamodels.NewDataModelImpl(sdkClient),
		Auth:       auth.NewAuthImpl(sdkClient),
		ACL:        dataassets.NewACLImpl(sdkClient),
		Account:    accounts.NewAccountsImpl(sdkClient),
	}
}

// func (sdk *SDK) Reinitialize(config SDKConfig) {
// 	sdk.APIKey = config.ApiKey

// 	client := resty.New()
// 	client.SetBaseURL(sdk.BaseURL)
// 	if newAPIKey != "" {
// 		client.SetAuthToken(newAPIKey) /
// 	}

// }
