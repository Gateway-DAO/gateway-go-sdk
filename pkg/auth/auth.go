package auth

import (
	"errors"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type Auth interface {
	Login(message string, signature string, wallet_address string) (string, error)
	GenerateSignMessage() (string, error)
	RefreshToken() (string, error)
}

type AuthImpl struct {
	Config common.SDKConfig
}

func NewAuthImpl(config common.SDKConfig) *AuthImpl {
	return &AuthImpl{
		Config: config,
	}
}

// check here for jwt response && walletAddress
func (u *AuthImpl) Login(message string, signature string, wallet_address string) (string, error) {

	var jwtTokenResponse common.TokenResponse = common.TokenResponse{Token: ""}
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&common.AuthRequest{Message: message, Signature: signature, WalletAddress: wallet_address}).SetResult(&jwtTokenResponse).SetError(&error).Post(common.AuthenticateAccount)

	if err != nil {
		return jwtTokenResponse.Token, err
	}

	if res.IsError() {
		return jwtTokenResponse.Token, errors.New(error.Error)
	}

	return jwtTokenResponse.Token, nil
}

func (u *AuthImpl) GenerateSignMessage() (string, error) {

	var messageResponse common.MessageResponse = common.MessageResponse{Message: ""}
	var error common.Error

	res, err := u.Config.Client.R().SetResult(&messageResponse).SetError(&error).Get(common.GenerateSignMessage)

	if err != nil {
		return messageResponse.Message, err
	}

	if res.IsError() {
		return messageResponse.Message, errors.New(error.Error)
	}

	return messageResponse.Message, nil
}

func (u *AuthImpl) RefreshToken() (string, error) {

	var jwtTokenResponse common.TokenResponse = common.TokenResponse{Token: ""}
	var error common.Error

	res, err := u.Config.Client.R().SetResult(&jwtTokenResponse).SetError(&error).Get(common.RefreshToken)

	if err != nil {
		return jwtTokenResponse.Token, err
	}

	if res.IsError() {
		return jwtTokenResponse.Token, errors.New(error.Error)
	}

	return jwtTokenResponse.Token, nil

}
