package transactions

import (
	"encoding/hex"
	"fmt"
	"go-stack/common"
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
	lpDataBytes, err := hex.DecodeString(common.IntToHex(int64(len(lpList.Values)), lpList.LengthPrefixedBytes))
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, lpDataBytes)

	//for (const l of list) {
	//    bytesArray.push(serializeStacksMessage(l));
	//  }

	concatArray, err := common.ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}
	return concatArray, nil
}

func SerializeAddress(address Address) ([]byte, error) {
	var bytesArray [][]byte
	versionBytes, err := hex.DecodeString(common.IntToHex(int64(address.Version), 1))
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, versionBytes)

	hash160Bytes, err := hex.DecodeString(address.Hash160)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, hash160Bytes)

	concatArray, err := common.ConcatArray(bytesArray)
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

	concatArray, err := common.ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}
	return concatArray, nil
}

//export function serializeMemoString(memoString: MemoString): Uint8Array {
//  const bytesArray = [];
//  const contentBytes = utf8ToBytes(memoString.content);
//  const paddedContent = rightPadHexToLength(bytesToHex(contentBytes), MEMO_MAX_LENGTH_BYTES * 2);
//  bytesArray.push(hexToBytes(paddedContent));
//  return concatArray(bytesArray);
//}
