package transactions

import (
	"fmt"
)

type TransactionSigner struct {
	Transaction   *StacksTransaction `json:"transaction"`
	SigHash       string             `json:"sigHash"`
	OriginDone    bool               `json:"originDone"`
	CheckOversign bool               `json:"checkOversign"`
	CheckOverlap  bool               `json:"checkOverlap"`
}

func CreateTransactionSigner(transaction *StacksTransaction) (TransactionSigner, error) {
	signHash, err := transaction.SignBegin()
	if err != nil {
		return TransactionSigner{}, err
	}
	return TransactionSigner{
		Transaction:   transaction,
		SigHash:       signHash,
		OriginDone:    false,
		CheckOversign: true,
		CheckOverlap:  true,
	}, nil
}

func (s *TransactionSigner) SignOrigin(privateKey StacksPrivateKey) error {
	if s.CheckOverlap && s.OriginDone {
		return fmt.Errorf("cannot sign origin after sponsor key")
	}

	nextSignHash, err := s.Transaction.SignNextOrigin(s.SigHash, privateKey)
	if err != nil {
		return err
	}
	s.SigHash = nextSignHash
	return nil
}
