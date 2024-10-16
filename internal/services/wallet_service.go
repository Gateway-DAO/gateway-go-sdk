package services

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type WalletTypeEnum string

const (
	Ethereum WalletTypeEnum = "ethereum"
	Solana   WalletTypeEnum = "solana"
)

type WalletSignMessageType struct {
	Signature  string
	SigningKey string
}

type Wallet interface {
	SignMessage(message string) (WalletSignMessageType, error)
}

type WalletService struct {
	wallet        Wallet
	walletType    WalletTypeEnum
	walletPrivKey string
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
	default:
		return nil, fmt.Errorf("unsupported wallet type")
	}

	return &WalletService{
		wallet:        wallet,
		walletType:    walletType,
		walletPrivKey: walletPrivateKey,
	}, nil
}

func (ws *WalletService) SignMessage(message string) (WalletSignMessageType, error) {
	return ws.wallet.SignMessage(message)
}
