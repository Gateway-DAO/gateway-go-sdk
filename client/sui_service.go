package client

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
	"golang.org/x/crypto/blake2b"
)

type SuiService struct {
	walletPrivateKey ed25519.PrivateKey
	walletAddress    string
}

const (
	SUI_PRIVATE_KEY_PREFIX = "suiprivkey"
	PRIVATE_KEY_SIZE       = 32
	SUI_ADDRESS_LENGTH     = 20
)

type SignaturePubkeyPair struct {
	SignatureScheme string
	Signature       []byte
	PubKey          []byte
}

var SIGNATURE_FLAG_TO_SCHEME = map[byte]string{
	0x00: "ED25519",
}

type ParsedKeypair struct {
	Schema    string
	SecretKey []byte
}

type SigFlag byte

const SignatureScheme = 3

const (
	SigFlagEd25519 SigFlag = 0x00
)

func ed25519PublicKeyToSuiAddress(pubKey []byte) string {
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

	seed := make([]byte, PRIVATE_KEY_SIZE)
	copy(seed, secretKey[:PRIVATE_KEY_SIZE])
	keypair := ed25519.NewKeyFromSeed(seed)

	signData := []byte("sui validation")
	signature := ed25519.Sign(keypair, signData)
	if !ed25519.Verify(keypair.Public().(ed25519.PublicKey), signData, signature) {
		return nil, nil, errors.New("provided secretKey is invalid")
	}

	return keypair.Public().(ed25519.PublicKey), keypair, nil
}

func parseSerializedSignature(serializedSignature string) (*SignaturePubkeyPair, error) {
	_bytes, err := base64.StdEncoding.DecodeString(serializedSignature)
	if err != nil {
		return nil, err
	}

	signatureScheme := "ED25519"
	if strings.EqualFold(serializedSignature, "") {
		return nil, fmt.Errorf("multiSig is not supported")
	}

	signature := _bytes[1 : len(_bytes)-32]
	pubKeyBytes := _bytes[1+len(signature):]

	keyPair := &SignaturePubkeyPair{
		SignatureScheme: signatureScheme,
		Signature:       signature,
		PubKey:          pubKeyBytes,
	}
	return keyPair, nil
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

func isHex(value string) bool {
	match, _ := regexp.MatchString("^(0x|0X)?[a-fA-F0-9]+$", value)
	return match && len(value)%2 == 0
}

func getHexByteLength(value string) int {
	if strings.HasPrefix(value, "0x") || strings.HasPrefix(value, "0X") {
		return (len(value) - 2) / 2
	}
	return len(value) / 2
}

// Note this is a custom implementation of Sui Wallet in go.
func NewSuiService(walletPrivateKey string) *SuiService {
	decoded, err := decodeSuiPrivateKey(walletPrivateKey)
	if err != nil {
		fmt.Println("Error decoding private key:", err)
	}

	if decoded.Schema != "ED25519" {
		fmt.Printf("Expected an ED25519 keypair, got %s\n", decoded.Schema)
	}

	pub, private, err := fromSecretKey(decoded.SecretKey)
	if err != nil {
		fmt.Println("Error creating keypair:", err)
	}

	publicKeyHex := hex.EncodeToString(pub)

	return &SuiService{
		walletPrivateKey: private,
		walletAddress:    publicKeyHex,
	}
}

func (es *SuiService) SignMessage(message string) (WalletSignMessageType, error) {
	messageBytes := []byte(message)
	bcsBytes := append([]byte{uint8(len(messageBytes))}, messageBytes...)
	intentMessage := messageWithIntent(bcsBytes)

	digest := blake2b.Sum256(intentMessage)

	signature := ed25519.Sign(es.walletPrivateKey, digest[:])

	serializedSignature := toSerializedSignature(signature, SignatureScheme, es.walletPrivateKey.Public().(ed25519.PublicKey))

	return WalletSignMessageType{
		Signature:  serializedSignature,
		SigningKey: es.walletAddress,
	}, nil
}

func VerifySuiMessage(signature string, message, walletAddress string) (bool, error) {
	serializedSignature, err := parseSerializedSignature(signature)
	if err != nil {
		return false, err
	}

	messageBytes := []byte(message)
	bcsBytes := append([]byte{uint8(len(messageBytes))}, messageBytes...)

	messageBytes = messageWithIntent(bcsBytes)

	digest := blake2b.Sum256(messageBytes)
	pass := ed25519.Verify(serializedSignature.PubKey[:], digest[:], serializedSignature.Signature)

	if !pass {
		return false, errors.New("signature verification failed")
	}

	derivedAddress := ed25519PublicKeyToSuiAddress(serializedSignature.PubKey)

	return strings.EqualFold(derivedAddress, walletAddress), nil
}

func ValidateSuiWallet(walletAddress string) bool {
	return isHex(walletAddress) && getHexByteLength(walletAddress) == SUI_ADDRESS_LENGTH
}

func (es *SuiService) GetWallet() string {
	return es.walletAddress
}
