package transactions

import (
	"github.com/stretchr/testify/assert"
	"go-stack/network"
	"math/big"
	"testing"
)

func TestCreateStacksPrivateKey(t *testing.T) {
	privateKeyHex := "1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56"
	privateKey, err := CreateStacksPrivateKey(privateKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	privateKeyTest := []byte{
		26, 31, 242, 17, 100, 45, 105, 53,
		83, 158, 172, 212, 4, 232, 63, 125,
		194, 186, 8, 134, 139, 152, 160, 12,
		243, 254, 155, 194, 107, 63, 125, 86,
	}

	assert.Equal(t, privateKey.Data.Serialize(), privateKeyTest, "invalid")
}

func TestGetStacksPublicKey(t *testing.T) {
	privateKeyHex := "1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56"
	privateKey, err := CreateStacksPrivateKey(privateKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	publicKey := GetPublicKey(privateKey)
	publicKeyTest := []byte{
		4, 149, 105, 90, 113, 57, 223, 198, 140, 95, 12,
		23, 151, 94, 247, 0, 135, 192, 202, 162, 212, 110,
		179, 98, 159, 77, 215, 243, 62, 90, 149, 226, 210,
		45, 215, 215, 51, 248, 1, 1, 255, 194, 226, 168,
		25, 125, 50, 183, 141, 176, 2, 104, 54, 117, 161,
		71, 66, 174, 13, 173, 160, 242, 80, 124, 159,
	}

	assert.Equal(t, publicKey.Data.SerializeUncompressed(), publicKeyTest, "invalid")
}

func TestPublicKeyToString(t *testing.T) {
	privateKeyHex := "1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56"
	privateKey, err := CreateStacksPrivateKey(privateKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	publicKey := GetPublicKey(privateKey)
	publicKeyHex := "0495695a7139dfc68c5f0c17975ef70087c0caa2d46eb3629f4dd7f33e5a95e2d22dd7d733f80101ffc2e2a8197d32b78db002683675a14742ae0dada0f2507c9f"
	assert.Equal(t, publicKey.String(), publicKeyHex, "invalid")
}

func TestCreateStacksPublicKey(t *testing.T) {
	publicKeyHex := "0495695a7139dfc68c5f0c17975ef70087c0caa2d46eb3629f4dd7f33e5a95e2d22dd7d733f80101ffc2e2a8197d32b78db002683675a14742ae0dada0f2507c9f"
	publicKey, err := CreateStacksPublickey(publicKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	publicKeyTest := []byte{
		4, 149, 105, 90, 113, 57, 223, 198, 140, 95, 12,
		23, 151, 94, 247, 0, 135, 192, 202, 162, 212, 110,
		179, 98, 159, 77, 215, 243, 62, 90, 149, 226, 210,
		45, 215, 215, 51, 248, 1, 1, 255, 194, 226, 168,
		25, 125, 50, 183, 141, 176, 2, 104, 54, 117, 161,
		71, 66, 174, 13, 173, 160, 242, 80, 124, 159,
	}

	assert.Equal(t, publicKey.Data.SerializeUncompressed(), publicKeyTest, "invalid")
}

func TestSignWithKey(t *testing.T) {
	privateKey, err := CreateStacksPrivateKey("1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56")
	if err != nil {
		t.Fatal(err)
	}

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

	signer, err := CreateTransactionSigner(&transaction)
	if err != nil {
		t.Fatal(err)
	}

	err = signer.SignOrigin(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, transaction.Auth.SpendingCondition.Signature.Data, "008c5ce936e99e20518ec1dfbb76b6caf5c8084d452134890df55c068fc9d510004c04293c28248c0cd5d732625aa884661989c459e48c175c73c30cf15df8c926", "invalid")
}

func TestSignWithKey1(t *testing.T) {
	privateKey, err := CreateStacksPrivateKey("1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56")
	if err != nil {
		t.Fatal(err)
	}

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
			Fee:        big.NewInt(78043),
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

	err = signer.SignOrigin(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, transaction.Auth.SpendingCondition.Signature.Data, "0175a792d3d13d128c8d72f06f8fcc716b1724b92d9fc23df52b8b995375925ec74170dcc5ae1564d9154f28e7dbd980f02541eb3f91628a1304ceec6c3d6ee9ef", "invalid")
}
