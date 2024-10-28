package client

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ed25519"

	"github.com/blocto/solana-go-sdk/types"
	"github.com/gagliardetto/solana-go"
	"github.com/mr-tron/base58"
)

type SolanaService struct {
	walletPrivateKey []byte
	wallet           types.Account
}

func NewSolanaService(walletPrivateKey string) *SolanaService {
	privateKey := solana.MustPrivateKeyFromBase58(walletPrivateKey)

	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		log.Printf("Failed to create wallet from private key: %v", err)
		panic(err)
	}

	return &SolanaService{
		walletPrivateKey: privateKey,
		wallet:           wallet,
	}
}

func (ss *SolanaService) SignMessage(message string) (WalletSignMessageType, error) {
	messageBytes := []byte(message)

	signedMessage := ed25519.Sign(ss.wallet.PrivateKey, messageBytes)
	signature := base58.Encode(signedMessage)

	return WalletSignMessageType{
		Signature:  signature,
		SigningKey: ss.wallet.PublicKey.ToBase58(),
	}, nil
}

func VerifySolanaMessage(message, signature, publicKey string) (bool, error) {
	signatureBytes, err := base58.Decode(signature)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature from Base58: %v", err)
	}

	publicKeyBytes, err := base58.Decode(publicKey)
	if err != nil {
		return false, fmt.Errorf("failed to decode public key from Base58: %v", err)
	}

	if len(publicKeyBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid public key length: expected %d bytes, got %d", ed25519.PublicKeySize, len(publicKeyBytes))
	}

	if len(signatureBytes) != ed25519.SignatureSize {
		return false, fmt.Errorf("invalid signature length: expected %d bytes, got %d", ed25519.SignatureSize, len(signatureBytes))
	}

	isValid := ed25519.Verify(publicKeyBytes, []byte(message), signatureBytes)

	return isValid, nil
}

func ValidateSolanaWallet(wallet string) bool {
	_, err := solana.PublicKeyFromBase58(wallet)
	return err == nil
}

func (ss *SolanaService) GetWallet() string {
	return ss.wallet.PublicKey.String()
}
