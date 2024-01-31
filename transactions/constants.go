package transactions

type AuthType byte

const (
	// StandardAuthType. The transaction is not sponsored. The sender will need to spend fees
	StandardAuthType AuthType = 0x04
	// SponsoredAuthType. The transaction is sponsored. The sponsor will spend fees on behalf of the sender
	SponsoredAuthType AuthType = 0x05
)

type PostConditionMode byte

const (
	AllowPostConditionMode PostConditionMode = 0x01
	DenyPostConditionMode  PostConditionMode = 0x02
)

type PostConditionType byte

const (
	STXPostConditionType         PostConditionType = 0x00
	FungiblePostConditionType    PostConditionType = 0x01
	NonFungiblePostConditionType PostConditionType = 0x02
)

type AddressHashMode byte

const (
	SerializeP2PKH  AddressHashMode = 0x00
	SerializeP2SH   AddressHashMode = 0x01
	SerializeP2WPKH AddressHashMode = 0x02
	SerializeP2WSH  AddressHashMode = 0x03
)

type SingleSigHashMode AddressHashMode

const (
	SerializeP2PKHSingleSigHashMode  SingleSigHashMode = SingleSigHashMode(SerializeP2PKH)
	SerializeP2WPKHSingleSigHashMode SingleSigHashMode = SingleSigHashMode(SerializeP2WPKH)
)

type PubKeyEncoding byte

const (
	Compressed   PubKeyEncoding = 0x00
	Uncompressed PubKeyEncoding = 0x01
)

type StacksMessageType int

const (
	AddressMessageType            StacksMessageType = 0
	MemoStringMessageType         StacksMessageType = 3
	PublicKeyMessageType          StacksMessageType = 6
	LengthPrefixedListMessageType StacksMessageType = 7
	PayloadMessageType            StacksMessageType = 8
	MessageSignatureMessageType   StacksMessageType = 9
)

type AnchorMode byte

const (
	OnChainOnlyAnchorMode  AnchorMode = 0x01
	OffChainOnlyAnchorMode AnchorMode = 0x02
	AnyAnchorMode          AnchorMode = 0x03
)

type TransactionVersion int

const (
	MainnetTransactionVersion TransactionVersion = 0x00
	TestnetTransactionVersion TransactionVersion = 0x80
)

type PayloadType byte

const (
	TokenTransferPayloadType          PayloadType = 0x00
	SmartContractPayloadType          PayloadType = 0x01
	VersionedSmartContractPayloadType PayloadType = 0x06
	ContractCallPayloadType           PayloadType = 0x02
	PoisonMicroblockPayloadType       PayloadType = 0x03
	CoinbasePayloadType               PayloadType = 0x04
	CoinbaseToAltRecipientPayloadType PayloadType = 0x05
)

type AddressVersion int

const (
	MainnetSingleSig AddressVersion = 22
	MainnetMultiSig  AddressVersion = 20
	TestnetSingleSig AddressVersion = 26
	TestnetMultiSig  AddressVersion = 21
)

type ChainID uint32

const (
	TestnetChainID ChainID = 0x80000000
	MainnetChainID ChainID = 0x00000001
)

type ClarityType byte

const (
	IntClarityType               ClarityType = 0x00
	UIntClarityType              ClarityType = 0x01
	BufferClarityType            ClarityType = 0x02
	BoolTrueClarityType          ClarityType = 0x03
	BoolFalseClarityType         ClarityType = 0x04
	PrincipalStandardClarityType ClarityType = 0x05
	PrincipalContractClarityType ClarityType = 0x06
	ResponseOkClarityType        ClarityType = 0x07
	ResponseErrClarityType       ClarityType = 0x08
	OptionalNoneClarityType      ClarityType = 0x09
	OptionalSomeClarityType      ClarityType = 0x0a
	ListClarityType              ClarityType = 0x0b
	TupleClarityType             ClarityType = 0x0c
	StringASCIIClarityType       ClarityType = 0x0d
	StringUTF8ClarityType        ClarityType = 0x0e
)

const (
	RECOVERABLE_ECDSA_SIG_LENGTH_BYTES int = 65
	MEMO_MAX_LENGTH_BYTES              int = 34
)

type AnchorModeNames string

const (
	//OnChainOnlyAnchorMode  AnchorModeNames = "onChainOnly"
	//OffChainOnlyAnchorMode AnchorModeNames = "offChainOnly"
	AnyChainAnchorMode AnchorModeNames = "any"
)

type PeerNetworkID int

const (
	MainnetPeerNetworkID PeerNetworkID = 0x17000000
	TestnetPeerNetworkID PeerNetworkID = 0xff000000
)

const PRIVATE_KEY_COMPRESSED_LENGTH = 33
const PRIVATE_KEY_UNCOMPRESSED_LENGTH = 32
const BLOCKSTACK_DEFAULT_GAIA_HUB_URL = "https://hub.blockstack.org"
