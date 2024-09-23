package accounts

import (
	"errors"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type Accounts interface {
	CreateAccount(accountDetails common.AccountCreateRequest) (string, error)
	Me() (common.MyAccountResponse, error)
	UpdateAccount(updateDetails common.AccountUpdateRequest) (common.MyAccountResponse, error)
}

type AccountsImpl struct {
	Config common.SDKConfig
}

func NewAccountsImpl(config common.SDKConfig) *AccountsImpl {
	return &AccountsImpl{
		Config: config,
	}
}

func (u *AccountsImpl) CreateAccount(accountDetails common.AccountCreateRequest) (string, error) {
	var jwtTokenResponse common.TokenResponse = common.TokenResponse{Token: ""}
	var error common.Error

	res, err := u.Config.Client.R().SetBody(accountDetails).SetResult(&jwtTokenResponse).SetError(&error).Post(common.CreateAccount)

	if err != nil {
		return jwtTokenResponse.Token, err
	}

	if res.IsError() {
		return jwtTokenResponse.Token, errors.New(error.Error)
	}

	return jwtTokenResponse.Token, nil
}

func (u *AccountsImpl) Me() (common.MyAccountResponse, error) {
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

func (u *AccountsImpl) UpdateAccount(updateDetails common.AccountUpdateRequest) (common.MyAccountResponse, error) {
	var myAccountResponse common.MyAccountResponse
	var error common.Error

	res, err := u.Config.Client.R().SetBody(updateDetails).SetResult(&myAccountResponse).SetError(&error).Patch(common.GetMyAccount)

	if err != nil {
		return myAccountResponse, err
	}

	if res.IsError() {
		return myAccountResponse, errors.New(error.Error)
	}

	return myAccountResponse, nil
}
