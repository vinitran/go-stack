package transactions

type MessageSignature struct {
	Type StacksMessageType `json:"type"`
	Data string            `json:"data"`
}

type Address struct {
	Type    StacksMessageType `json:"type"`
	Version AddressVersion    `json:"version"`
	Hash160 string            `json:"hash_160"`
}
