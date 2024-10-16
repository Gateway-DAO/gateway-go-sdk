package services

import (
	"fmt"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type Wallet interface {
	SignMessage(message string) (WalletSignMessageType, error)
}

type WalletService struct {
	wallet        Wallet
	walletType    common.WalletTypeEnum
	walletPrivKey string
}

func NewWalletService(walletPrivateKey string, walletType common.WalletTypeEnum) (*WalletService, error) {
	var wallet Wallet

	switch walletType {
	case common.Ethereum:
		wallet = NewEtherumService(walletPrivateKey)
	case common.Solana:
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
