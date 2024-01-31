package transactions

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

type SingleSigSpendingCondition struct {
	HashMode     SingleSigHashMode `json:"hashMode"`
	Signer       string            `json:"signer"`
	Nonce        *big.Int          `json:"nonce"`
	Fee          *big.Int          `json:"fee"`
	KeyEncodeing PubKeyEncoding    `json:"keyEncodeing"`
	Signature    MessageSignature  `json:"signature"`
}

type StandardAuthorization struct {
	AuthType          AuthType                   `json:"authType"`
	SpendingCondition SingleSigSpendingCondition `json:"spendingCondition"`
}

func (s *StandardAuthorization) SetFee(fee *big.Int) {
	s.SpendingCondition.Fee = fee
}

func (s *StandardAuthorization) SetNonce(nonce *big.Int) {
	s.SpendingCondition.Nonce = nonce
}

func SerializeAuthorization(auth StandardAuthorization) ([]byte, error) {
	var bytesArray [][]byte
	bytesArray = append(bytesArray, []byte{byte(auth.AuthType)})

	switch auth.AuthType {
	case StandardAuthType:
		dataByte, err := SerializeSingleSigSpendingCondition(auth.SpendingCondition)
		if err != nil {
			log.Println("asd")
			return nil, err
		}
		bytesArray = append(bytesArray, dataByte)
	case SponsoredAuthType:
		log.Println("not supported")
	}

	concatArray, err := ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}

	return concatArray, nil
}

func SerializeSingleSigSpendingCondition(condition SingleSigSpendingCondition) ([]byte, error) {
	var bytesArray [][]byte
	bytesArray = append(bytesArray, []byte{byte(condition.HashMode)})

	signerBytes, err := hex.DecodeString(condition.Signer)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, signerBytes)
	bytesArray = append(bytesArray, BigIntToBytes(condition.Nonce, 8))
	bytesArray = append(bytesArray, BigIntToBytes(condition.Fee, 8))
	bytesArray = append(bytesArray, []byte{byte(condition.KeyEncodeing)})

	signatureBytes, err := hex.DecodeString(condition.Signature.Data)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, signatureBytes)
	concatArray, err := ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}

	return concatArray, nil
}

func CreateSingleSigSpendingCondition(hashMode SingleSigHashMode, publicKey StacksPublicKey, nonce *big.Int, fee *big.Int) (SingleSigSpendingCondition, error) {
	// address version arg doesn't matter for signer hash generation

	signer, err := AddressFromPublicKeys(TestnetSingleSig, AddressHashMode(hashMode), 1, []StacksPublicKey{publicKey})
	if err != nil {
		return SingleSigSpendingCondition{}, err
	}

	keyEncoding := Uncompressed

	return SingleSigSpendingCondition{
		HashMode:     hashMode,
		Signer:       signer.Hash160,
		Nonce:        nonce,
		Fee:          fee,
		KeyEncodeing: keyEncoding,
		Signature:    EmptyMessageSignature(),
	}, nil
}

func EmptyMessageSignature() MessageSignature {
	dataBytes := make([]byte, RECOVERABLE_ECDSA_SIG_LENGTH_BYTES)
	return MessageSignature{
		Type: MessageSignatureMessageType,
		Data: hex.EncodeToString(dataBytes),
	}
}

func CreateStandardAuth(spendingCondition SingleSigSpendingCondition) StandardAuthorization {
	return StandardAuthorization{
		AuthType:          StandardAuthType,
		SpendingCondition: spendingCondition,
	}
}

func IntoInitialSighashAuth(auth StandardAuthorization) (StandardAuthorization, error) {
	switch auth.AuthType {
	case StandardAuthType:
		return CreateStandardAuth(ClearCondition(auth.SpendingCondition)), nil
	default:
		return StandardAuthorization{}, fmt.Errorf("unexpected authorization type for signing")
	}
}

func ClearCondition(condition SingleSigSpendingCondition) SingleSigSpendingCondition {
	return SingleSigSpendingCondition{
		HashMode:     condition.HashMode,
		Signer:       condition.Signer,
		Nonce:        big.NewInt(0),
		Fee:          big.NewInt(0),
		KeyEncodeing: condition.KeyEncodeing,
		Signature:    EmptyMessageSignature(),
	}
}

func NextSignature(curSignHash string, authType AuthType, fee *big.Int, nonce *big.Int, privateKey StacksPrivateKey) (MessageSignature, string, error) {
	sigHashPreSign, err := MakeSigHashPreSign(curSignHash, authType, fee, nonce)
	if err != nil {
		return MessageSignature{}, "", err
	}

	signature, err := SignWithKey(privateKey, sigHashPreSign)
	if err != nil {
		return MessageSignature{}, "", err
	}

	publicKey := GetPublicKey(privateKey)
	nextSigHash, err := MakeSigHashPostSign(sigHashPreSign, publicKey, signature)
	if err != nil {
		return MessageSignature{}, "", err
	}
	return signature, nextSigHash, nil
}

func MakeSigHashPostSign(curSigHash string, pubKey StacksPublicKey, signature MessageSignature) (string, error) {
	hashLength := 32 + 1 + RECOVERABLE_ECDSA_SIG_LENGTH_BYTES
	sigHash := fmt.Sprintf("%s%s%s", curSigHash, LeftPadHex(strconv.FormatInt(int64(1), 16)), signature.Data)
	sigHashBytes, err := hex.DecodeString(sigHash)
	if err != nil {
		return "", err
	}

	if len(sigHashBytes) > hashLength {
		return "", fmt.Errorf("invalid signature hash length")
	}

	return TxidFromData(sigHashBytes), nil
}

func MakeSigHashPreSign(curSignHash string, authType AuthType, fee *big.Int, nonce *big.Int) (string, error) {
	// new hash combines the previous hash and all the new data this signature will add. This
	// includes:
	// * the previous hash
	// * the auth flag
	// * the tx fee (big-endian 8-byte number)
	// * nonce (big-endian 8-byte number)
	hashLength := 32 + 1 + 8 + 8

	signHash := fmt.Sprintf("%s%s%s%s",
		curSignHash,
		hex.EncodeToString([]byte{byte(authType)}),
		hex.EncodeToString(BigIntToBytes(fee, 8)),
		hex.EncodeToString(BigIntToBytes(nonce, 8)),
	)

	sigHashBytes, err := hex.DecodeString(signHash)
	if err != nil {
		return "", err
	}

	if len(sigHashBytes) != hashLength {
		return "", fmt.Errorf("invalid signature hash length")
	}

	return TxidFromData(sigHashBytes), nil
}
