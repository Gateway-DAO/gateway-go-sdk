package accounts

import (
	"errors"

	"gateway-go-sdk/pkg/common"
)

type Accounts interface {
	Create(accountDetails common.AccountCreateRequest) (string, error)
	GetMe() (common.MyAccountResponse, error)
	UpdateMe(updateDetails common.AccountUpdateRequest) (common.MyAccountResponse, error)
}

type AccountsImpl struct {
	Config common.SDKConfig
	Wallet WalletInterface
}

func NewAccountsImpl(config common.SDKConfig) *AccountsImpl {
	wallet := NewWalletImpl(config)
	return &AccountsImpl{
		Config: config,
		Wallet: wallet,
	}
}

func (u *AccountsImpl) Create(accountDetails common.AccountCreateRequest) (string, error) {
	var jwtTokenResponse common.TokenResponse
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&accountDetails).SetResult(&jwtTokenResponse).SetError(&error).Post(common.CreateAccount)

	if err != nil {
		return jwtTokenResponse.Token, err
	}

	if res.IsError() {
		return jwtTokenResponse.Token, errors.New(error.Error)
	}

	return jwtTokenResponse.Token, nil
}

func (u *AccountsImpl) GetMe() (common.MyAccountResponse, error) {
	var myAccountResponse common.MyAccountResponse
	var error common.Error

	res, err := u.Config.Client.R().SetResult(&myAccountResponse).SetError(&error).Get(common.GetMyAccount)

	if err != nil {
		return myAccountResponse, err
	}

	if res.IsError() {
		return myAccountResponse, errors.New(error.Error)
	}

	return myAccountResponse, nil

}

func (u *AccountsImpl) UpdateMe(updateDetails common.AccountUpdateRequest) (common.MyAccountResponse, error) {
	var myAccountResponse common.MyAccountResponse
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&updateDetails).SetResult(&myAccountResponse).SetError(&error).Patch(common.GetMyAccount)

	if err != nil {
		return myAccountResponse, err
	}

	if res.IsError() {
		return myAccountResponse, errors.New(error.Error)
	}

	return myAccountResponse, nil
}
