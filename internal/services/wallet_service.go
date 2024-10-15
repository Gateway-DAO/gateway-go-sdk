package services

import (
	"fmt"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
)

type WalletTypeEnum string

const (
	Ethereum WalletTypeEnum = "ethereum"
	Solana   WalletTypeEnum = "solana"
)

// WalletService holds the wallet type and the appropriate service (Ethereum or Solana).
type WalletService struct {
	wallet        WalletService
	walletType    WalletTypeEnum
	walletPrivKey string
}

func NewWalletService(walletPrivateKey string, walletType WalletTypeEnum) (*WalletService, error) {
	var wallet WalletService

	switch walletType {
	case Ethereum:
		wallet = services.NewEtherumService(walletPrivateKey)
	case Solana:
		wallet = services.NewSolanaService(walletPrivateKey)
	default:
		return nil, fmt.Errorf("unsupported wallet type")
	}

	return &WalletService{
		wallet:        wallet,
		walletType:    walletType,
		walletPrivKey: walletPrivateKey,
	}, nil
}

// SignMessage signs a message using the appropriate wallet service.
func (ws *WalletService) SignMessage(message string) (WalletSignMessageType, error) {
	return ws.wallet.SignMessage(message)
}
