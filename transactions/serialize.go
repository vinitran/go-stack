package transactions

import (
	"fmt"
	"go-stack/common"
)

func BytesWithTypeId(typeId ClarityType, bytes []byte) ([]byte, error) {
	var bytesArray [][]byte
	bytesArray = append(bytesArray, []byte{byte(typeId)})
	bytesArray = append(bytesArray, bytes)

	concatArray, err := common.ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}
	return concatArray, nil
}

func SerializeStandardPrincipalCV(cv StandardPrincipalCV) ([]byte, error) {
	serializeAddr, err := SerializeAddress(cv.Address)
	if err != nil {
		return nil, err
	}
	return BytesWithTypeId(cv.Type, serializeAddr)
}

func SerializeCV(value ClarityValue) ([]byte, error) {
	switch value.GetType() {
	case PrincipalStandardClarityType:
		return SerializeStandardPrincipalCV(value.(StandardPrincipalCV))
	default:
		return nil, fmt.Errorf("unable to serialize. Invalid Clarity Value")
	}
}
