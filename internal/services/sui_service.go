package services

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/bech32"
	"golang.org/x/crypto/blake2b"
)

type SuiService struct {
	walletPrivateKey ed25519.PrivateKey
	walletAddress    string
}

const (
	SUI_PRIVATE_KEY_PREFIX = "suiprivkey" // replace with actual prefix
	PRIVATE_KEY_SIZE       = 32
	SUI_ADDRESS_LENGTH     = 20
)

var SIGNATURE_FLAG_TO_SCHEME = map[byte]string{
	0x00: "ED25519",
	0x01: "Secp256k1",
}

type ParsedKeypair struct {
	Schema    string
	SecretKey []byte
}

type SigFlag byte

const (
	SigFlagEd25519   SigFlag = 0x00
	SigFlagSecp256k1 SigFlag = 0x01
)

func Ed25519PublicKeyToSuiAddress(pubKey []byte) string {
	newPubkey := []byte{byte(SigFlagEd25519)}
	newPubkey = append(newPubkey, pubKey...)

	addrBytes := blake2b.Sum256(newPubkey)
	return fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])
}

func decodeSuiPrivateKey(value string) (ParsedKeypair, error) {
	prefix, words, err := bech32.Decode(value)
	if err != nil {
		return ParsedKeypair{}, err
	}
	if prefix != SUI_PRIVATE_KEY_PREFIX {
		return ParsedKeypair{}, errors.New("invalid private key prefix")
	}

	extendedSecretKey, err := bech32.ConvertBits(words, 5, 8, false)
	if err != nil {
		return ParsedKeypair{}, err
	}

	signatureScheme := SIGNATURE_FLAG_TO_SCHEME[extendedSecretKey[0]]
	secretKey := extendedSecretKey[1:]

	return ParsedKeypair{
		Schema:    signatureScheme,
		SecretKey: secretKey,
	}, nil
}

func fromSecretKey(secretKey []byte) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	secretKeyLength := len(secretKey)
	if secretKeyLength != PRIVATE_KEY_SIZE {
		return nil, nil, fmt.Errorf("wrong secretKey size. Expected %d bytes, got %d", PRIVATE_KEY_SIZE, secretKeyLength)
	}

	// Generate keypair from seed
	seed := make([]byte, PRIVATE_KEY_SIZE)
	copy(seed, secretKey[:PRIVATE_KEY_SIZE])
	keypair := ed25519.NewKeyFromSeed(seed)

	// Validation
	signData := []byte("sui validation")
	signature := ed25519.Sign(keypair, signData)
	if !ed25519.Verify(keypair.Public().(ed25519.PublicKey), signData, signature) {
		return nil, nil, errors.New("provided secretKey is invalid")
	}

	return keypair.Public().(ed25519.PublicKey), keypair, nil
}

func messageWithIntent(message []byte) []byte {
	intent := []byte{3, 0, 0}
	intentMessage := make([]byte, len(intent)+len(message))
	copy(intentMessage, intent)
	copy(intentMessage[len(intent):], message)
	return intentMessage
}

func toSerializedSignature(signature []byte, scheme byte, publicKey ed25519.PublicKey) string {
	serialized := make([]byte, 1+len(signature)+len(publicKey))
	serialized[0] = scheme
	copy(serialized[1:], signature)
	copy(serialized[1+len(signature):], publicKey)
	return base64.StdEncoding.EncodeToString(serialized)
}

const SignatureScheme = 3

func SuiGenerateSignature(privateKey ed25519.PrivateKey, message string) (string, error) {
	// Prepare message with intent
	messageBytes := []byte(message)
	bcsBytes := append([]byte{uint8(len(messageBytes))}, messageBytes...)
	intentMessage := messageWithIntent(bcsBytes)

	// Hash the message
	digest := blake2b.Sum256(intentMessage)

	// Sign the message
	signature := ed25519.Sign(privateKey, digest[:])

	// Serialize signature with scheme and public key
	return toSerializedSignature(signature, SignatureScheme, privateKey.Public().(ed25519.PublicKey)), nil
}

func NewSuiService(walletPrivateKey string) *SuiService {
	log.Println(walletPrivateKey)
	decoded, err := decodeSuiPrivateKey(walletPrivateKey)
	if err != nil {
		fmt.Println("Error decoding private key:", err)
	}

	// Check if schema is ED25519
	if decoded.Schema != "ED25519" {
		fmt.Printf("Expected an ED25519 keypair, got %s\n", decoded.Schema)
	}

	// Create keypair from secret key
	pub, private, err := fromSecretKey(decoded.SecretKey)
	if err != nil {
		fmt.Println("Error creating keypair:", err)
	}

	message := "hello world"

	signature, err := SuiGenerateSignature(private, message)
	if err != nil {
		fmt.Println("Error signing message:", err)
	}
	fmt.Println("Signature:", signature)
	publicKeyHex := hex.EncodeToString(pub)
	a := Ed25519PublicKeyToSuiAddress(pub)
	log.Println(a)
	log.Println(err)
	return &SuiService{
		walletPrivateKey: private,
		walletAddress:    publicKeyHex,
	}
}

func (es *SuiService) SignMessage(message string) (WalletSignMessageType, error) {
	// messageHash := accounts.TextHash([]byte(message))

	// signature, err := crypto.Sign(messageHash, es.walletPrivateKey)

	// if err != nil {
	// 	return WalletSignMessageType{}, fmt.Errorf("failed to sign message: %v", err)
	// }

	// if signature[64] < 27 {
	// 	signature[64] += 27
	// }

	return WalletSignMessageType{
		Signature:  "",
		SigningKey: es.walletAddress,
	}, nil
}

// func VerifySuiMessage(signature string, message, walletAddress string) (bool, error) {
// 	signatureHex, err := hexutil.Decode(signature)
// 	if err != nil {
// 		return false, err
// 	}

// 	signatureHex[crypto.RecoveryIDOffset] -= 27

// 	messageHash := accounts.TextHash([]byte(message))

// 	pubKey, err := crypto.SigToPub(messageHash, signatureHex)
// 	if err != nil {
// 		return false, err
// 	}

// 	if common.HexToAddress(walletAddress) != crypto.PubkeyToAddress(*pubKey) {
// 		return false, fmt.Errorf("invalid signature")
// 	}

// 	return true, nil
// }

// func ValidateSuiWallet(wallet string) bool {
// 	return common.IsHexAddress(wallet)
// }

// func (es *SuiService) GetWallet() string {
// 	return es.walletAddress
// }
