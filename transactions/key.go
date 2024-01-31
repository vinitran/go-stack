package transactions

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

type StacksPrivateKey struct {
	Compressed bool                  `json:"compressed"`
	Data       *secp256k1.PrivateKey `json:"data"`
}

func CreateStacksPrivateKey(key string) (StacksPrivateKey, error) {
	data, err := PrivateKeyToBytes(key)
	if err != nil {
		return StacksPrivateKey{}, err
	}

	compressed := false
	if len(data) == PRIVATE_KEY_COMPRESSED_LENGTH {
		compressed = true
	}

	return StacksPrivateKey{
		Compressed: compressed,
		Data:       secp256k1.PrivKeyFromBytes(data),
	}, nil
}

func PrivateKeyToBytes(privateKey string) ([]byte, error) {
	privateKeyBuffer, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	if len(privateKeyBuffer) != 32 && len(privateKeyBuffer) != 33 {
		return nil, fmt.Errorf("improperly formatted private-key. Private-key byte length should be 32 or 33. Length provided: %d", len(privateKeyBuffer))
	}

	if len(privateKeyBuffer) == 33 && privateKeyBuffer[32] != 1 {
		return nil, fmt.Errorf("improperly formatted private-key. 33 bytes indicate compressed key, but the last byte must be == 01")
	}

	return privateKeyBuffer, nil
}

type StacksPublicKey struct {
	Data *secp256k1.PublicKey
}

func (s StacksPublicKey) GetType() int {
	return int(PublicKeyMessageType) // PublicKeyMessageType
}

func (s StacksPublicKey) String() string {
	return hex.EncodeToString(s.Data.SerializeUncompressed())
}

func GetPublicKey(privateKey StacksPrivateKey) StacksPublicKey {
	return StacksPublicKey{Data: privateKey.Data.PubKey()}
}

func CreateStacksPublickey(key string) (StacksPublicKey, error) {
	publicKeyBytes, err := hex.DecodeString(key)
	if err != nil {
		return StacksPublicKey{}, err
	}

	publicKey, err := secp256k1.ParsePubKey(publicKeyBytes)
	if err != nil {
		return StacksPublicKey{}, err
	}
	return StacksPublicKey{Data: publicKey}, nil
}

func SignWithKey(privateKey StacksPrivateKey, messageHash string) (MessageSignature, error) {
	messageBytes, err := hex.DecodeString(messageHash)
	if err != nil {
		return MessageSignature{}, err
	}
	signature := ecdsa.SignCompact(privateKey.Data, messageBytes, false)
	signature[0] -= 27
	return MessageSignature{
		Type: MessageSignatureMessageType,
		Data: hex.EncodeToString(signature),
	}, nil
}
