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
	Version                        common.TransactionVersion `json:"version"`
	ChainID                        common.ChainID            `json:"chain_id"`
	BnsLookupUrl                   string                    `json:"bnsLookupUrl"`
	BroadcastEndpoint              string                    `json:"broadcastEndpoint"`
	TransferFeeEstimateEndpoint    string                    `json:"transferFeeEstimateEndpoint"`
	TransactionFeeEstimateEndpoint string                    `json:"transactionFeeEstimateEndpoint"`
	AccountEndpoint                string                    `json:"accountEndpoint"`
	ContractAbiEndpoint            string                    `json:"contractAbiEndpoint"`
	ReadOnlyFunctionCallEndpoint   string                    `json:"readOnlyFunctionCallEndpoint"`
	CoreApiUrl                     string                    `json:"coreApUrl"`
	FetchFn                        FetchFn                   `json:"fetchFn"`
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
