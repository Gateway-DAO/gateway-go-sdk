package client

import (
	"github.com/go-resty/resty/v2"
)

type SDK struct {
	DataAssets DataAsset
	DataModel  DataModel
	Account    *AccountsImpl
	ACL        ACL
	Auth       Auth
}

type SDKConfig struct {
	ApiKey        string
	WalletDetails WalletDetails
	URL           string
}

type WalletDetails struct {
	PrivateKey string
	WalletType WalletTypeEnum
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
		wallet, _ := NewWalletService(config.WalletDetails.PrivateKey, config.WalletDetails.WalletType)
		params := MiddlewareParams{
			Client: client,
			Wallet: *wallet,
		}
		client.OnBeforeRequest(AuthMiddleware(params))
	}

	sdkClient := Config{
		Client: client,
	}

	return &SDK{
		DataAssets: NewDataAssetImpl(sdkClient),
		DataModel:  NewDataModelImpl(sdkClient),
		Auth:       NewAuthImpl(sdkClient),
		ACL:        NewACLImpl(sdkClient),
		Account:    NewAccountsImpl(sdkClient),
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

// 	sdkClient :=  SDKConfig{
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
