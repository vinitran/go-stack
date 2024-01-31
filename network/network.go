package network

type StacksNetworkName string

const (
	MainnetNetworkName StacksNetworkName = "mainnet"
	TestnetNetworkName StacksNetworkName = "testnet"
	DevnetNetworkName  StacksNetworkName = "devnet"
	MocknetNetworkName StacksNetworkName = "mocknet"
)

type StacksNetwork struct {
	//Version                        transactions.TransactionVersion `json:"version"`
	//ChainID                        transactions.ChainID            `json:"chain_id"`
	BnsLookupUrl                   string  `json:"bnsLookupUrl"`
	BroadcastEndpoint              string  `json:"broadcastEndpoint"`
	TransferFeeEstimateEndpoint    string  `json:"transferFeeEstimateEndpoint"`
	TransactionFeeEstimateEndpoint string  `json:"transactionFeeEstimateEndpoint"`
	AccountEndpoint                string  `json:"accountEndpoint"`
	ContractAbiEndpoint            string  `json:"contractAbiEndpoint"`
	ReadOnlyFunctionCallEndpoint   string  `json:"readOnlyFunctionCallEndpoint"`
	CoreApiUrl                     string  `json:"coreApUrl"`
	FetchFn                        FetchFn `json:"fetchFn"`
}

type NetworkConfig struct {
	Url     string
	FetchFn FetchFn
}

//func NewStacksNetwork(networkName StacksNetworkName) StacksNetwork {
//	switch networkName {
//	case MainnetNetworkName:
//		return NewStacksMainnet()
//	}
//}

func NewStacksMainnet(networkConfig NetworkConfig) StacksNetwork {
	return StacksNetwork{
		//Version:                        transactions.MainnetTransactionVersion,
		//ChainID:                        transactions.MainnetChainID,
		BnsLookupUrl:                   "https://stacks-node-api.mainnet.stacks.co",
		BroadcastEndpoint:              "/v2/transactions",
		TransferFeeEstimateEndpoint:    "/v2/fees/transfer",
		TransactionFeeEstimateEndpoint: "/v2/fees/transaction",
		AccountEndpoint:                "/v2/accounts",
		ContractAbiEndpoint:            "/v2/contracts/interface",
		ReadOnlyFunctionCallEndpoint:   "/v2/contracts/call-read",
		CoreApiUrl:                     networkConfig.Url,
		FetchFn:                        networkConfig.FetchFn,
	}
}

func NewStacksTestnet(networkConfig NetworkConfig) StacksNetwork {
	return StacksNetwork{
		//Version:                        transactions.TestnetTransactionVersion,
		//ChainID:                        transactions.TestnetChainID,
		BnsLookupUrl:                   "https://stacks-node-api.mainnet.stacks.co",
		BroadcastEndpoint:              "/v2/transactions",
		TransferFeeEstimateEndpoint:    "/v2/fees/transfer",
		TransactionFeeEstimateEndpoint: "/v2/fees/transaction",
		AccountEndpoint:                "/v2/accounts",
		ContractAbiEndpoint:            "/v2/contracts/interface",
		ReadOnlyFunctionCallEndpoint:   "/v2/contracts/call-read",
		CoreApiUrl:                     networkConfig.Url,
		FetchFn:                        networkConfig.FetchFn,
	}
}
