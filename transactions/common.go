package transactions

import "go-stack/c32check"

type MessageSignature struct {
	Type StacksMessageType `json:"type"`
	Data string            `json:"data"`
}

type Address struct {
	Type    StacksMessageType `json:"type"`
	Version AddressVersion    `json:"version"`
	Hash160 string            `json:"hash_160"`
}

func AddressFromVersionHash(version AddressVersion, hash string) Address {
	return Address{
		Type:    AddressMessageType,
		Version: version,
		Hash160: hash,
	}
}

func CreateAddress(c32AddressString string) Address {
	version, address := c32check.C32addressDecode(c32AddressString)
	return Address{
		Type:    AddressMessageType,
		Version: AddressVersion(version),
		Hash160: address,
	}
}
