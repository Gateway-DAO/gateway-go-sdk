package services

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type EtherumService struct {
	walletPrivateKey *ecdsa.PrivateKey
	walletAddress    string
}

func NewEtherumService(walletPrivateKey string) *EtherumService {
	privateKey, err := crypto.HexToECDSA(walletPrivateKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to get the public key from private key")
	}

	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &EtherumService{
		walletPrivateKey: privateKey,
		walletAddress:    walletAddress,
	}
}

func (es *EtherumService) SignMessage(message string) (WalletSignMessageType, error) {
	messageHash := accounts.TextHash([]byte(message))

	signature, err := crypto.Sign(messageHash, es.walletPrivateKey)

	if err != nil {
		return WalletSignMessageType{}, fmt.Errorf("failed to sign message: %v", err)
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	return WalletSignMessageType{
		Signature:  hexutil.Encode(signature),
		SigningKey: es.walletAddress,
	}, nil
}

func (es *EtherumService) VerifyMessage(signature []byte, message, walletAddress string) (bool, error) {
	msgHash := crypto.Keccak256Hash([]byte(message))

	publicKey, err := crypto.SigToPub(msgHash.Bytes(), signature)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*publicKey).Hex()

	return recoveredAddr == walletAddress, nil
}

func (es *EtherumService) ValidateWallet(wallet string) bool {
	return common.IsHexAddress(wallet)
}

func (es *EtherumService) GetWallet() string {
	return es.walletAddress
}
