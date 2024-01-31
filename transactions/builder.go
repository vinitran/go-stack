package transactions

import (
	"go-stack/network"
	"math/big"
)

type TokenTransferOptions struct {
	Recipient  string                `json:"recipient"`
	Amount     *big.Int              `json:"amount"`
	Fee        *big.Int              `json:"fee"`
	Nonce      *big.Int              `json:"nonce"`
	Network    network.StacksNetwork `json:"network"`
	AnchorMode AnchorMode            `json:"anchorMode"`
	Memo       string                `json:"memo"`
	Sponsored  bool                  `json:"sponsored"`
}

type UnsignedTokenTransferOptions struct {
	TokenTransferOptions
	PublicKey string
}

type SignedTokenTransferOptions struct {
	TokenTransferOptions
	SenderKey string
	PublicKey StacksPublicKey
}

func MakeUnsignedSTXTokenTransfer(txOpts SignedTokenTransferOptions) (StacksTransaction, error) {
	payload, err := CreateTokenTransferPayload(txOpts.Recipient, txOpts.Amount, txOpts.Memo)
	if err != nil {
		return StacksTransaction{}, err
	}

	spendingCondition, err := CreateSingleSigSpendingCondition(
		SerializeP2PKHSingleSigHashMode,
		txOpts.PublicKey,
		txOpts.Nonce,
		txOpts.Fee)
	if err != nil {
		return StacksTransaction{}, err
	}

	authorization := CreateStandardAuth(spendingCondition)

	transaction := CreateNewStacksTransaction(
		TestnetTransactionVersion,
		authorization,
		DenyPostConditionMode,
		TestnetChainID,
		payload,
	)

	transaction.SetFee(txOpts.Fee)
	transaction.SetNonce(txOpts.Nonce)
	return transaction, nil
}

//func MakeSTXTokenTransfer(opts SignedTokenTransferOptions) (StacksTransaction, error) {
//	privateKey, err := CreateStacksPrivateKey(opts.SenderKey)
//	if err != nil {
//		return StacksTransaction{}, err
//	}
//
//	publicKey := GetPublicKey(privateKey)
//	opts.PublicKey = publicKey
//
//	transaction, err := MakeUnsignedSTXTokenTransfer(opts)
//	if err != nil {
//		return StacksTransaction{}, err
//	}
//
//	signer, err := CreateTransactionSigner(transaction)
//	if err != nil {
//		return StacksTransaction{}, err
//	}
//
//}
