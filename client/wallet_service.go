package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type WalletTypeEnum string

const (
	Ethereum WalletTypeEnum = "ethereum"
	Solana   WalletTypeEnum = "solana"
	Sui      WalletTypeEnum = "sui"
)

type Wallet interface {
	SignMessage(message string) (WalletSignMessageType, error)
}

type WalletService struct {
	Wallet        Wallet
	WalletType    WalletTypeEnum
	WalletPrivKey string
}

type MiddlewareParams struct {
	Client *resty.Client
	Wallet WalletService
}

func NewWalletService(walletPrivateKey string, walletType WalletTypeEnum) (*WalletService, error) {
	var wallet Wallet

	switch walletType {
	case Ethereum:
		wallet = NewEtherumService(walletPrivateKey)
	case Solana:
		wallet = NewSolanaService(walletPrivateKey)
	case Sui:
		wallet = NewSuiService(walletPrivateKey)
	default:
		return nil, fmt.Errorf("unsupported wallet type")
	}

	return &WalletService{
		Wallet:        wallet,
		WalletType:    walletType,
		WalletPrivKey: walletPrivateKey,
	}, nil
}

func (ws *WalletService) SignMessage(message string) (WalletSignMessageType, error) {
	return ws.Wallet.SignMessage(message)
}
