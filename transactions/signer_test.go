package transactions

import (
	"github.com/stretchr/testify/assert"
	"go-stack/network"
	"math/big"
	"testing"
)

func TestCreateNewTransactionSigner(t *testing.T) {
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
			Fee:        big.NewInt(1586),
			Nonce:      big.NewInt(10),
			Network:    network.StacksNetwork{},
			AnchorMode: 3,
			Memo:       "",
			Sponsored:  false,
		},
		PublicKey: publicKey,
	})

	signer, err := CreateTransactionSigner(&transaction)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, signer.SigHash, "f44ab308a3c29e08c37b93e2bbf20b91b3143bbb11f3ae4790b97fe9400c9379", "invalid")
}
