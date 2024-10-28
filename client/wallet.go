package client

import (
	"errors"
)

type WalletInterface interface {
	Add(address string) (MyAccountResponse, error)
	Remove(address string) (MyAccountResponse, error)
}

type WalletImpl struct {
	Config Config
}

func NewWalletImpl(config Config) *WalletImpl {
	return &WalletImpl{
		Config: config,
	}
}

func (u *WalletImpl) Add(address string) (MyAccountResponse, error) {
	var myAccount MyAccountResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(map[string]interface{}{"address": address}).SetResult(&myAccount).SetError(&error).Post(AddWallet)

	if err != nil {
		return myAccount, err
	}

	if res.IsError() {
		return myAccount, errors.New(error.Error)
	}

	return myAccount, nil
}

func (u *WalletImpl) Remove(address string) (MyAccountResponse, error) {
	var myAccount MyAccountResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(map[string]interface{}{"address": address}).SetResult(&myAccount).SetError(&error).Delete(RemoveWallet)

	if err != nil {
		return myAccount, err
	}

	if res.IsError() {
		return myAccount, errors.New(error.Error)
	}

	return myAccount, nil
}
