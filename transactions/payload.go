package transactions

import (
	"go-stack/common"
	"math/big"
)

type Payload interface{}

type ContractCallPayload struct {
	Type            StacksMessageType  `json:"type"`
	PayloadType     PayloadType        `json:"payloadType"`
	ContractAddress Address            `json:"contractAddress"`
	ContractName    LengthPrefixedList `json:"contractName"`
	FunctionName    LengthPrefixedList `json:"functionName"`
}

type TokenTransferPayload struct {
	Type        StacksMessageType   `json:"type"`
	PayloadType PayloadType         `json:"payloadType"`
	Recipient   StandardPrincipalCV `json:"recipient"`
	Amount      *big.Int            `json:"amount"`
	Memo        MemoString          `json:"memo"`
}

func SerializeTokenTransferPayload(payload TokenTransferPayload) ([]byte, error) {
	var bytesArray [][]byte
	bytesArray = append(bytesArray, []byte{byte(payload.PayloadType)})

	recipientBytes, err := SerializeCV(payload.Recipient)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, recipientBytes)
	bytesArray = append(bytesArray, common.BigIntToBytes(payload.Amount, 8))

	messageBytes, err := SerializeStacksMessage(payload.Memo)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, messageBytes)

	concatArray, err := common.ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}
	return concatArray, nil
}
