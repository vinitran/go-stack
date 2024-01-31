package transactions

import (
	"math/big"
)

type StacksTransaction struct {
	Version           TransactionVersion    `json:"version"`
	ChainId           ChainID               `json:"chainId"`
	Auth              StandardAuthorization `json:"auth"`
	AnchorMode        AnchorMode            `json:"anchorMode"`
	Payload           TokenTransferPayload  `json:"payload"`
	PostConditionMode PostConditionMode     `json:"postConditionMode"`
	PostConditions    LengthPrefixedList    `json:"postConditions"`
}

func CreateNewStacksTransaction(
	version TransactionVersion,
	auth StandardAuthorization,
	//postConditions LengthPrefixedList,
	postConditionMode PostConditionMode,
	//anchorMode AnchorMode,
	chainID ChainID,
	payload TokenTransferPayload,
) StacksTransaction {
	return StacksTransaction{
		Version:           version,
		ChainId:           chainID,
		Auth:              auth,
		AnchorMode:        AnyAnchorMode,
		Payload:           payload,
		PostConditionMode: postConditionMode,
		PostConditions:    CreateLPList(),
	}
}

func (s *StacksTransaction) SetFee(fee *big.Int) {
	s.Auth.SetFee(fee)
}

func (s *StacksTransaction) SetNonce(nonce *big.Int) {
	s.Auth.SetNonce(nonce)
}

func (s StacksTransaction) Serialize() ([]byte, error) {
	var bytesArray [][]byte
	bytesArray = append(bytesArray, []byte{byte(s.Version)})

	chainIdBytes := make([]byte, 4)
	WriteUInt32BE(&chainIdBytes, uint32(s.ChainId), 0)
	bytesArray = append(bytesArray, chainIdBytes)

	authByte, err := SerializeAuthorization(s.Auth)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, authByte)

	bytesArray = append(bytesArray, []byte{byte(s.AnchorMode)})
	bytesArray = append(bytesArray, []byte{byte(s.PostConditionMode)})

	postConditionBytes, err := SerializeLPList(s.PostConditions)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, postConditionBytes)

	payloadBytes, err := SerializeTokenTransferPayload(s.Payload)
	if err != nil {
		return nil, err
	}
	bytesArray = append(bytesArray, payloadBytes)
	concatArray, err := ConcatArray(bytesArray)
	if err != nil {
		return nil, err
	}

	return concatArray, nil
}

func (s StacksTransaction) Txid() (string, error) {
	serialized, err := s.Serialize()
	if err != nil {
		return "", err
	}
	return TxidFromData(serialized), nil
}

func (s StacksTransaction) SignBegin() (string, error) {
	tx := s
	auth, err := IntoInitialSighashAuth(tx.Auth)
	if err != nil {
		return "", err
	}
	tx.Auth = auth
	return tx.Txid()
}

func (s *StacksTransaction) SignNextOrigin(sigHash string, privateKey StacksPrivateKey) (string, error) {
	nextSig, nextSigHash, err := NextSignature(sigHash, s.Auth.AuthType, s.Auth.SpendingCondition.Fee, s.Auth.SpendingCondition.Nonce, privateKey)
	if err != nil {
		return "", err
	}

	s.Auth.SpendingCondition.Signature = nextSig
	return nextSigHash, nil
}
