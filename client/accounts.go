package client

import (
	"errors"
)

type Accounts interface {
	Create(accountDetails AccountCreateRequest) (string, error)
	GetMe() (MyAccountResponse, error)
	UpdateMe(updateDetails AccountUpdateRequest) (MyAccountResponse, error)
}

type AccountsImpl struct {
	Config Config
	Wallet WalletInterface
}

func NewAccountsImpl(config Config) *AccountsImpl {
	wallet := NewWalletImpl(config)
	return &AccountsImpl{
		Config: config,
		Wallet: wallet,
	}
}

func (u *AccountsImpl) Create(accountDetails AccountCreateRequest) (string, error) {
	var jwtTokenResponse TokenResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(&accountDetails).SetResult(&jwtTokenResponse).SetError(&error).Post(CreateAccount)

	if err != nil {
		return jwtTokenResponse.Token, err
	}

	if res.IsError() {
		return jwtTokenResponse.Token, errors.New(error.Error)
	}

	return jwtTokenResponse.Token, nil
}

func (u *AccountsImpl) GetMe() (MyAccountResponse, error) {
	var myAccountResponse MyAccountResponse
	var error Error

	res, err := u.Config.Client.R().SetResult(&myAccountResponse).SetError(&error).Get(GetMyAccount)

	if err != nil {
		return myAccountResponse, err
	}

	if res.IsError() {
		return myAccountResponse, errors.New(error.Error)
	}

	return myAccountResponse, nil

}

func (u *AccountsImpl) UpdateMe(updateDetails AccountUpdateRequest) (MyAccountResponse, error) {
	var myAccountResponse MyAccountResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(&updateDetails).SetResult(&myAccountResponse).SetError(&error).Patch(GetMyAccount)

	if err != nil {
		return myAccountResponse, err
	}

	if res.IsError() {
		return myAccountResponse, errors.New(error.Error)
	}

	return myAccountResponse, nil
}
