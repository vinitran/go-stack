package transactions

import (
	"go-stack/common"
	"log"
)

type StacksTransaction struct {
	Version           TransactionVersion
	ChainId           ChainID
	Auth              StandardAuthorization
	AnchorMode        AnchorMode
	Payload           TokenTransferPayload
	PostConditionMode PostConditionMode
	PostConditions    LengthPrefixedList
}

func CreateNewStacksTransaction(
	version TransactionVersion,
	auth StandardAuthorization,
	//postConditions LengthPrefixedList,
	postConditionMode PostConditionMode,
	//anchorMode AnchorMode,
	chainID ChainID,
	payload TokenTransferPayload,
) StacksTransaction {
	return StacksTransaction{
		Version:           version,
		ChainId:           chainID,
		Auth:              auth,
		AnchorMode:        AnyAnchorMode,
		Payload:           payload,
		PostConditionMode: postConditionMode,
		PostConditions:    CreateLPList(),
	}
}

func (s StacksTransaction) serialize() ([]byte, error) {
	var bytesArray [][]byte
	bytesArray = append(bytesArray, []byte{byte(s.Version)})

	chainIdBytes := make([]byte, 4)
	common.WriteUInt32BE(&chainIdBytes, uint32(s.ChainId), 0)
	bytesArray = append(bytesArray, chainIdBytes)

	authByte, err := SerializeAuthorization(s.Auth)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, authByte)

	bytesArray = append(bytesArray, []byte{byte(s.AnchorMode)})
	bytesArray = append(bytesArray, []byte{byte(s.PostConditionMode)})

	postConditionBytes, err := SerializeLPList(s.PostConditions)
	if err != nil {
		log.Println(2)
		return nil, err
	}
	bytesArray = append(bytesArray, postConditionBytes)

	payloadBytes, err := SerializeTokenTransferPayload(s.Payload)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, payloadBytes)
	concatArray, err := common.ConcatArray(bytesArray)
	if err != nil {
		log.Println(3)
		return nil, err
	}

	return concatArray, nil
}
