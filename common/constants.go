package common

type AnchorModeNames string

const (
	OnChainOnlyAnchorMode  AnchorModeNames = "onChainOnly"
	OffChainOnlyAnchorMode AnchorModeNames = "offChainOnly"
	AnyChainAnchorMode     AnchorModeNames = "any"
)

type TransactionVersion int

const (
	MainnetTransactionVersion TransactionVersion = 0x00
	TestnetTransactionVersion TransactionVersion = 0x80
)

type ChainID uint32

const (
	TestnetChainID ChainID = 0x80000000
	MainnetChainID ChainID = 0x00000001
)

type PeerNetworkID int

const (
	MainnetPeerNetworkID PeerNetworkID = 0x17000000
	TestnetPeerNetworkID PeerNetworkID = 0xff000000
)

const PRIVATE_KEY_COMPRESSED_LENGTH = 33
const PRIVATE_KEY_UNCOMPRESSED_LENGTH = 32
const BLOCKSTACK_DEFAULT_GAIA_HUB_URL = "https://hub.blockstack.org"
