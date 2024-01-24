package transactions

import (
	"encoding/hex"
	"go-stack/common"
	"log"
	"math/big"
)

type SingleSigSpendingCondition struct {
	HashMode     SingleSigHashMode
	Signer       string
	Nonce        *big.Int
	Fee          *big.Int
	KeyEncodeing PubKeyEncoding
	Signature    MessageSignature
}

type StandardAuthorization struct {
	AuthType          AuthType
	SpendingCondition SingleSigSpendingCondition
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

//export function serializeAuthorization(auth: Authorization): Uint8Array {
//  const bytesArray = [];
//  bytesArray.push(auth.authType);
//
//  switch (auth.authType) {
//    case AuthType.Standard:
//      bytesArray.push(serializeSpendingCondition(auth.spendingCondition));
//      break;
//    case AuthType.Sponsored:
//      bytesArray.push(serializeSpendingCondition(auth.spendingCondition));
//      bytesArray.push(serializeSpendingCondition(auth.sponsorSpendingCondition));
//      break;
//  }
//
//  return concatArray(bytesArray);
//}

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

//export function serializeSpendingCondition(condition: SpendingConditionOpts): Uint8Array {
//  if (isSingleSig(condition)) {
//    return serializeSingleSigSpendingCondition(condition);
//  }
//  return serializeMultiSigSpendingCondition(condition);
//}
