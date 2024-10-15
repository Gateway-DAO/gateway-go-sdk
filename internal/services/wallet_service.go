package services

import (
	"fmt"
)

type WalletTypeEnum string

const (
	Ethereum WalletTypeEnum = "ethereum"
	Solana   WalletTypeEnum = "solana"
)

type Wallet interface {
	SignMessage(message string) (WalletSignMessageType, error)
}

type WalletService struct {
	wallet        Wallet
	walletType    WalletTypeEnum
	walletPrivKey string
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
