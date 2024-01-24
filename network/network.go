package network

import "go-stack/common"

type StackNetwork string

const (
	MainnetNetwork StackNetwork = "mainnet"
	TestnetNetwork StackNetwork = "testnet"
	DevnetNetwork  StackNetwork = "devnet"
	MocknetNetwork StackNetwork = "mocknet"
)

type StacksNetwork struct {
	Version                        common.TransactionVersion
	ChainID                        common.ChainID
	BnsLookupUrl                   string
	BroadcastEndpoint              string
	TransferFeeEstimateEndpoint    string
	TransactionFeeEstimateEndpoint string
	AccountEndpoint                string
	ContractAbiEndpoint            string
	ReadOnlyFunctionCallEndpoint   string
	CoreApiUrl                     string
	FetchFn                        FetchFn
}

type NetworkConfig struct {
	Url     string
	FetchFn *FetchFn
}

func (s StacksNetwork) Init(networkConfig NetworkConfig) StacksNetwork {
	return StacksNetwork{
		Version:                        common.TestnetTransactionVersion,
		ChainID:                        common.MainnetChainID,
		BnsLookupUrl:                   "https://stacks-node-api.mainnet.stacks.co",
		BroadcastEndpoint:              "/v2/transactions",
		TransferFeeEstimateEndpoint:    "/v2/fees/transfer",
		TransactionFeeEstimateEndpoint: "/v2/fees/transaction",
		AccountEndpoint:                "/v2/accounts",
		ContractAbiEndpoint:            "/v2/contracts/interface",
		ReadOnlyFunctionCallEndpoint:   "/v2/contracts/call-read",
		CoreApiUrl:                     networkConfig.Url,
	}
}
