package transactions

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSerializeTransaction(t *testing.T) {
	tx := CreateNewStacksTransaction(
		TestnetTransactionVersion,
		StandardAuthorization{
			AuthType: StandardAuthType,
			SpendingCondition: SingleSigSpendingCondition{
				HashMode:     SerializeP2PKHSingleSigHashMode,
				Signer:       "1cf09bd26ec797f7562563ac9ea437f6909342f7",
				Nonce:        big.NewInt(9),
				Fee:          big.NewInt(114041),
				KeyEncodeing: 1,
				Signature: MessageSignature{
					Type: MessageSignatureMessageType,
					Data: "00cb88bc0da84438ac53930d52a507b9534317bed02b2a812c1370a60626bfeb8c440d232ec6e8d04c2dbd45fffc8318ec6beafc63fa977c8a2970b4da7012e28c",
				},
			},
		},
		DenyPostConditionMode,
		TestnetChainID,
		TokenTransferPayload{
			Type:        PayloadMessageType,
			PayloadType: TokenTransferPayloadType,
			Recipient: StandardPrincipalCV{
				Type: 5,
				Address: Address{
					Type:    0,
					Version: 26,
					Hash160: "a2954b5eb8270cabeaef121d7fd05cd5e1b12ab7",
				},
			},
			Amount: big.NewInt(1),
			Memo: MemoString{
				Type:    3,
				Content: "",
			},
		},
	)

	serializeTxBytes, err := tx.serialize()
	if err != nil {
		t.Fatal(err)
	}

	serializeTxHex := hex.EncodeToString(serializeTxBytes)
	serializeTxHexTest := "808000000004001cf09bd26ec797f7562563ac9ea437f6909342f70000000000000009000000000001bd790100cb88bc0da84438ac53930d52a507b9534317bed02b2a812c1370a60626bfeb8c440d232ec6e8d04c2dbd45fffc8318ec6beafc63fa977c8a2970b4da7012e28c03020000000000051aa2954b5eb8270cabeaef121d7fd05cd5e1b12ab7000000000000000100000000000000000000000000000000000000000000000000000000000000000000"
	assert.Equal(t, serializeTxHex, serializeTxHexTest, "invalid")
}
