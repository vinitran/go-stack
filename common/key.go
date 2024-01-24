package common

import (
	"encoding/hex"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"log"
)

func CreateStacksPrivateKey(key string) *secp256k1.PrivateKey {
	pkBytes, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}

	privateKey := secp256k1.PrivKeyFromBytes(pkBytes)
	return privateKey
}
