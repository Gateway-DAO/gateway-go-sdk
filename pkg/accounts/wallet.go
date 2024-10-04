package accounts

import (
	"errors"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type WalletInterface interface {
	Add(address string) (common.MyAccountResponse, error)
	Remove(address string) (common.MyAccountResponse, error)
}

type WalletImpl struct {
	Config common.SDKConfig
}

func NewWalletImpl(config common.SDKConfig) *WalletImpl {
	return &WalletImpl{
		Config: config,
	}
}

func (u *WalletImpl) Add(address string) (common.MyAccountResponse, error) {
	var myAccount common.MyAccountResponse
	var error common.Error

	res, err := u.Config.Client.R().SetBody(map[string]interface{}{"address": address}).SetResult(&myAccount).SetError(&error).Post(common.AddWallet)

	if err != nil {
		return myAccount, err
	}

	if res.IsError() {
		return myAccount, errors.New(error.Error)
	}

	return myAccount, nil
}

func (u *WalletImpl) Remove(address string) (common.MyAccountResponse, error) {
	var myAccount common.MyAccountResponse
	var error common.Error

	res, err := u.Config.Client.R().SetBody(map[string]interface{}{"address": address}).SetResult(&myAccount).SetError(&error).Delete(common.RemoveWallet)

	if err != nil {
		return myAccount, err
	}

	if res.IsError() {
		return myAccount, errors.New(error.Error)
	}

	return myAccount, nil
}
