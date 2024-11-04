package client

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type EtherumService struct {
	WalletPrivateKey *ecdsa.PrivateKey
	WalletAddress    string
}

func NewEtherumService(walletPrivateKey string) *EtherumService {
	privateKey, err := crypto.HexToECDSA(walletPrivateKey)
	if err != nil {
		log.Printf("Failed to load private key: %v", err)
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to get the public key from private key")
	}

	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &EtherumService{
		WalletPrivateKey: privateKey,
		WalletAddress:    walletAddress,
	}
}

func (es *EtherumService) SignMessage(message string) (WalletSignMessageType, error) {
	messageHash := accounts.TextHash([]byte(message))

	signature, err := crypto.Sign(messageHash, es.WalletPrivateKey)

	if err != nil {
		return WalletSignMessageType{}, fmt.Errorf("failed to sign message: %v", err)
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	return WalletSignMessageType{
		Signature:  hexutil.Encode(signature),
		SigningKey: es.WalletAddress,
	}, nil
}

func VerifyEtherumMessage(signature string, message, walletAddress string) (bool, error) {
	signatureHex, err := hexutil.Decode(signature)
	if err != nil {
		return false, err
	}

	signatureHex[crypto.RecoveryIDOffset] -= 27

	messageHash := accounts.TextHash([]byte(message))

	pubKey, err := crypto.SigToPub(messageHash, signatureHex)
	if err != nil {
		return false, err
	}

	if common.HexToAddress(walletAddress) != crypto.PubkeyToAddress(*pubKey) {
		return false, fmt.Errorf("invalid signature")
	}

	return true, nil
}

func ValidateEtherumWallet(wallet string) bool {
	return common.IsHexAddress(wallet)
}

func (es *EtherumService) GetWallet() string {
	return es.WalletAddress
}
