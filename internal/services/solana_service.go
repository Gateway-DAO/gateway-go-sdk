package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/ed25519"

	"github.com/blocto/solana-go-sdk/types"
	"github.com/gagliardetto/solana-go"
)

type SolanaService struct {
	walletPrivateKey []byte
	wallet           types.Account
}

func NewSolanaService(walletPrivateKey string) *SolanaService {
	privateKey, err := base64.StdEncoding.DecodeString(walletPrivateKey)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		log.Fatalf("Failed to create wallet from private key: %v", err)
	}

	return &SolanaService{
		walletPrivateKey: privateKey,
		wallet:           wallet,
	}
}

func (ss *SolanaService) SignMessage(message string) (WalletSignMessageType, error) {
	messageBytes := []byte(message)

	signedMessage := ed25519.Sign(ss.wallet.PrivateKey, messageBytes)
	signature := base64.StdEncoding.EncodeToString(signedMessage)

	return WalletSignMessageType{
		Signature:  []byte(signature),
		SigningKey: ss.wallet.PublicKey.ToBase58(),
	}, nil
}

func (ss *SolanaService) VerifyMessage(message, signature, publicKey string) (bool, error) {
	messageBytes := []byte(message)
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %v", err)
	}

	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false, fmt.Errorf("failed to decode public key: %v", err)
	}

	isValid := ed25519.Verify(publicKeyBytes, messageBytes, signatureBytes)
	return isValid, nil
}

func (ss *SolanaService) ValidateWallet(wallet string) (string, error) {
	pubKey, err := solana.PublicKeyFromBase58(wallet)
	if err != nil {
		return "", errors.New("invalid wallet address: " + err.Error())
	}
	return pubKey.String(), nil
}

func (ss *SolanaService) GetWallet() string {
	return ss.wallet.PublicKey.String()
}
