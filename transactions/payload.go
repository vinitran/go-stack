package transactions

import (
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
	bytesArray = append(bytesArray, BigIntToBytes(payload.Amount, 8))

	messageBytes, err := SerializeStacksMessage(payload.Memo)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, messageBytes)

	concatArray, err := ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}
	return concatArray, nil
}

func CreateTokenTransferPayload(recipient string, amount *big.Int, memo string) (TokenTransferPayload, error) {
	recipientAddr := CreateAddress(recipient)
	memoData, err := CreateMemotring(memo)
	if err != nil {
		return TokenTransferPayload{}, err
	}

	return TokenTransferPayload{
		Type:        PayloadMessageType,
		PayloadType: TokenTransferPayloadType,
		Recipient: StandardPrincipalCV{
			Type:    PrincipalStandardClarityType,
			Address: recipientAddr,
		},
		Amount: amount,
		Memo:   memoData,
	}, nil
}
