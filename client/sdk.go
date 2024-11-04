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

func (sdk *SDK) Reinitialize(config SDKConfig) *SDK {
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
