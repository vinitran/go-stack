package transactions

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"testing"
)

func TestEmptyMessageSignature(t *testing.T) {
	log.Println(hex.EncodeToString([]byte{byte(RECOVERABLE_ECDSA_SIG_LENGTH_BYTES)}))
}

func TestNextSignature(t *testing.T) {
	curSignHash := "f44ab308a3c29e08c37b93e2bbf20b91b3143bbb11f3ae4790b97fe9400c9379"
	senderKey := "1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56"
	privKey, err := CreateStacksPrivateKey(senderKey)
	if err != nil {
		t.Fatal(err)
	}

	sign, _, err := NextSignature(curSignHash, StandardAuthType, big.NewInt(155727), big.NewInt(10), privKey)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, sign.Data, "0002b9af1b6a86e5bb384a2b7a3fede402fd14e7c67f907bd550cc87c0930a99ac42bda63903391acd60b428494facbe31fd78e0635d3e600d43a8a2cfe83d2fd3", "invalid")
}
