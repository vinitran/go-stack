package transactions

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-stack/network"
	"log"
	"math/big"
	"testing"
)

func TestMakeSTXTokenTransfer(t *testing.T) {
	senderKey := "1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56"
	privKey, err := CreateStacksPrivateKey(senderKey)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := GetPublicKey(privKey)

	transaction, err := MakeUnsignedSTXTokenTransfer(SignedTokenTransferOptions{
		TokenTransferOptions: TokenTransferOptions{
			Recipient:  "ST2H9AJTYQ0KGSAZAXW91TZYGBKAY3C9APZXSSGXW",
			Amount:     big.NewInt(1),
			Fee:        big.NewInt(75757),
			Nonce:      big.NewInt(0),
			Network:    network.StacksNetwork{},
			AnchorMode: 0,
			Memo:       "",
			Sponsored:  false,
		},
		PublicKey: publicKey,
	})
	if err != nil {
		t.Fatal(err)
	}

	tx, _ := json.Marshal(transaction)
	fmt.Println(string(tx))

	serializeTxBytes, err := transaction.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	serializeTxHex := hex.EncodeToString(serializeTxBytes)

	log.Println("asdasd", serializeTxHex)
}
