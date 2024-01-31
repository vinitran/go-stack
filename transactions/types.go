package transactions

import (
	"encoding/hex"
	"fmt"
)

type LengthPrefixedList struct {
	Type                StacksMessageType
	LengthPrefixedBytes int
	Values              []StacksMessage
}

type StacksMessage interface {
	GetType() StacksMessageType
}

type MemoString struct {
	Type    StacksMessageType `json:"type"`
	Content string            `json:"content"`
}

func (m MemoString) GetType() StacksMessageType {
	return m.Type
}

func CreateLPList() LengthPrefixedList {
	return LengthPrefixedList{
		Type:                LengthPrefixedListMessageType,
		LengthPrefixedBytes: 4,
		Values:              []StacksMessage{},
	}
}

func SerializeLPList(lpList LengthPrefixedList) ([]byte, error) {
	var bytesArray [][]byte
	lpDataBytes, err := hex.DecodeString(IntToHex(int64(len(lpList.Values)), lpList.LengthPrefixedBytes))
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, lpDataBytes)

	//for (const l of list) {
	//    bytesArray.push(serializeStacksMessage(l));
	//  }

	concatArray, err := ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}
	return concatArray, nil
}

func SerializeAddress(address Address) ([]byte, error) {
	var bytesArray [][]byte
	versionBytes, err := hex.DecodeString(IntToHex(int64(address.Version), 1))
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, versionBytes)

	hash160Bytes, err := hex.DecodeString(address.Hash160)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, hash160Bytes)

	concatArray, err := ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}

	return concatArray, nil
}

func SerializeStacksMessage(message StacksMessage) ([]byte, error) {
	switch message.GetType() {
	case MemoStringMessageType:
		return SerializeMemoString(message.(MemoString))
	default:
		return nil, fmt.Errorf("invalid stack message type")
	}
}

func SerializeMemoString(memoString MemoString) ([]byte, error) {
	var bytesArray [][]byte
	contentHex := hex.EncodeToString([]byte(memoString.Content))
	paddedContent := RightPadHexToLength(contentHex, MEMO_MAX_LENGTH_BYTES*2)
	paddedContentBytes, err := hex.DecodeString(paddedContent)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, paddedContentBytes)

	concatArray, err := ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}
	return concatArray, nil
}

func AddressFromPublicKeys(version AddressVersion, hashMode AddressHashMode, numSigs int, publicKeys []StacksPublicKey) (Address, error) {
	if len(publicKeys) == 0 {
		return Address{}, fmt.Errorf("invalid number of public keys")
	}

	if hashMode == SerializeP2PKH || hashMode == SerializeP2WPKH {
		if len(publicKeys) != 1 || numSigs != 1 {
			return Address{}, fmt.Errorf("invalid number of public keys or signatures")
		}
	}

	//if hashMode == SerializeP2WPKH || hashMode == SerializeP2WSH {
	//	if len(publicKeys) != 1 || numSigs != 1 {
	//		return Address{}, fmt.Errorf("invalid number of public keys or signatures")
	//	}
	//}

	switch hashMode {
	case SerializeP2PKH:
		return AddressFromVersionHash(version, HashP2PKH(publicKeys[0].Data.SerializeUncompressed())), nil
	default:
		return Address{}, fmt.Errorf("invalid hashmode")
	}
}

func CreateMemotring(content string) (MemoString, error) {
	if ExceedsMaxLengthBytes(content, MEMO_MAX_LENGTH_BYTES) {
		return MemoString{}, fmt.Errorf("memo exceeds maximum length of %d bytes", MEMO_MAX_LENGTH_BYTES)
	}
	return MemoString{
		Type:    MemoStringMessageType,
		Content: content,
	}, nil
}
