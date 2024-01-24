package transactions

type MessageSignature struct {
	Type StacksMessageType
	Data string
}

type Address struct {
	Type    StacksMessageType
	Version AddressVersion
	Hash160 string
}
