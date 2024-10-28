package pkg

import (
	"github.com/Gateway-DAO/gateway-go-sdk/client/accounts"
	"github.com/Gateway-DAO/gateway-go-sdk/client/auth"
	"github.com/Gateway-DAO/gateway-go-sdk/client/common"
	dataassets "github.com/Gateway-DAO/gateway-go-sdk/client/data_assets"
	datamodels "github.com/Gateway-DAO/gateway-go-sdk/client/data_models"
	"github.com/Gateway-DAO/gateway-go-sdk/client/helpers"
	"github.com/Gateway-DAO/gateway-go-sdk/client/services"
	"github.com/go-resty/resty/v2"
)

type SDK struct {
	DataAssets dataassets.DataAsset
	DataModel  datamodels.DataModel
	Account    *accounts.AccountsImpl
	ACL        dataassets.ACL
	Auth       auth.Auth
}

type SDKConfig struct {
	ApiKey        string
	WalletDetails WalletDetails
	URL           string
}

type WalletDetails struct {
	PrivateKey string
	WalletType services.WalletTypeEnum
}

func NewSDK(config SDKConfig) *SDK {
	client := resty.New()
	if config.URL != "" {
		client.SetBaseURL(config.URL)
	} else {
		client.SetBaseURL("https://dev.api.gateway.tech")
	}

	if config.ApiKey != "" {
		client.SetAuthToken(config.ApiKey)
	} else {
		wallet, _ := services.NewWalletService(config.WalletDetails.PrivateKey, config.WalletDetails.WalletType)
		params := services.MiddlewareParams{
			Client: client,
			Wallet: *wallet,
		}
		client.OnBeforeRequest(helpers.AuthMiddleware(params))
	}

	sdkClient := common.SDKConfig{
		Client: client,
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
// 	client := resty.New()

// 	if config.URL != "" {
// 		client.SetBaseURL(config.URL)
// 	} else {
// 		client.SetBaseURL("https://dev.api.gateway.tech")
// 	}

// 	if config.ApiKey != "" {
// 		client.SetAuthToken(config.ApiKey)
// 	} else {
// 		wallet, _ := services.NewWalletService(config.WalletDetails.PrivateKey, config.WalletDetails.WalletType)
// 		params := services.MiddlewareParams{
// 			Client: client,
// 			Wallet: *wallet,
// 		}
// 		client.OnBeforeRequest(helpers.AuthMiddleware(params))
// 	}

// 	sdkClient := common.SDKConfig{
// 		Client: client,
// 	}

// 	return &SDK{
// 		DataAssets: dataassets.NewDataAssetImpl(sdkClient),
// 		DataModel:  datamodels.NewDataModelImpl(sdkClient),
// 		Auth:       auth.NewAuthImpl(sdkClient),
// 		ACL:        dataassets.NewACLImpl(sdkClient),
// 		Account:    accounts.NewAccountsImpl(sdkClient),
// 	}
// }
