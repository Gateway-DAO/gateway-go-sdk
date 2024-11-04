package client

import (
	"errors"
	"fmt"
)

type Auth interface {
	Login(message string, signature string, wallet_address string) (string, error)
	GetMessage() (string, error)
	GetRefreshToken() (string, error)
}

type AuthImpl struct {
	Config Config
}

func NewAuthImpl(config Config) *AuthImpl {
	return &AuthImpl{
		Config: config,
	}
}

func (u *AuthImpl) Login(message string, signature string, wallet_address string) (string, error) {
	var isValid bool
	var err error
	if ValidateEtherumWallet(wallet_address) {
		isValid, err = VerifyEtherumMessage(signature, message, wallet_address)
		if err != nil {
			return "", fmt.Errorf("ethereum signature verification failed: %v", err)
		}
		if !isValid {
			return "", errors.New("invalid Ethereum signature")
		}
	} else if ValidateSuiWallet(wallet_address) {
		isValid, err = VerifySuiMessage(message, signature, wallet_address)
		if err != nil {
			return "", fmt.Errorf("sui signature verification failed: %v", err)
		}
		if !isValid {
			return "", errors.New("invalid sui signature")
		}

	} else if ValidateSolanaWallet(wallet_address) {
		isValid, err = VerifySolanaMessage(message, signature, wallet_address)
		if err != nil {
			return "", fmt.Errorf("solana signature verification failed: %v", err)
		}
		if !isValid {
			return "", errors.New("invalid Solana signature")
		}
	}

	var jwtTokenResponse TokenResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(&AuthRequest{Message: message, Signature: signature, WalletAddress: wallet_address}).SetResult(&jwtTokenResponse).SetError(&error).Post(AuthenticateAccount)

	if err != nil {
		return jwtTokenResponse.Token, err
	}

	if res.IsError() {
		return jwtTokenResponse.Token, errors.New(error.Error)
	}

	return jwtTokenResponse.Token, nil
}

func (u *AuthImpl) GetMessage() (string, error) {

	var messageResponse MessageResponse
	var error Error

	res, err := u.Config.Client.R().SetResult(&messageResponse).SetError(&error).Get(GenerateSignMessage)
	if err != nil {
		return messageResponse.Message, err
	}

	if res.IsError() {
		return messageResponse.Message, errors.New(error.Error)
	}

	return messageResponse.Message, nil
}

func (u *AuthImpl) GetRefreshToken() (string, error) {

	var jwtTokenResponse TokenResponse
	var error Error

	res, err := u.Config.Client.R().SetResult(&jwtTokenResponse).SetError(&error).Get(RefreshToken)

	if err != nil {
		return jwtTokenResponse.Token, err
	}

	if res.IsError() {
		return jwtTokenResponse.Token, errors.New(error.Error)
	}

	return jwtTokenResponse.Token, nil
}
