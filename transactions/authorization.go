package transactions

import (
	"encoding/hex"
	"go-stack/common"
	"log"
	"math/big"
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

	concatArray, err := common.ConcatArray(bytesArray)
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
	bytesArray = append(bytesArray, common.BigIntToBytes(condition.Nonce, 8))
	bytesArray = append(bytesArray, common.BigIntToBytes(condition.Fee, 8))
	bytesArray = append(bytesArray, []byte{byte(condition.KeyEncodeing)})

	signatureBytes, err := hex.DecodeString(condition.Signature.Data)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, signatureBytes)
	concatArray, err := common.ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}

	return concatArray, nil
}
