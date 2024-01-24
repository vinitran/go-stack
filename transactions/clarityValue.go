package transactions

type ClarityValue interface {
	GetType() ClarityType
}

type TrueBooleanCV struct{}

func (b TrueBooleanCV) GetType() ClarityType {
	return BoolTrueClarityType
}

type FalseBooleanCV struct{}

func (b FalseBooleanCV) GetType() ClarityType {
	return BoolFalseClarityType
}

type StandardPrincipalCV struct {
	Type    ClarityType `json:"type"`
	Address Address     `json:"address"`
}

func (b StandardPrincipalCV) GetType() ClarityType {
	return PrincipalStandardClarityType
}

//
//type BufferCV struct {
//	Value []byte
//}
//
//func (b *BufferCV) isClarityValue() {}
//
//type IntCV struct {
//	Value int64
//}
//
//func (i *IntCV) isClarityValue() {}
//
//type UIntCV struct {
//	Value uint64
//}
//
//func (u *UIntCV) isClarityValue() {}
//
//type StandardPrincipalCV struct {
//	Value string
//}
//
//func (s *StandardPrincipalCV) isClarityValue() {}
//
//type ContractPrincipalCV struct {
//	Value string
//}
//
//func (c *ContractPrincipalCV) isClarityValue() {}
//
//type ResponseErrorCV struct {
//	Value ClarityValue
//}
//
//func (r *ResponseErrorCV) isClarityValue() {}
//
//type ResponseOkCV struct {
//	Value ClarityValue
//}
//
//func (r *ResponseOkCV) isClarityValue() {}
//
//type NoneCV struct{}
//
//func (n *NoneCV) isClarityValue() {}
//
//type SomeCV struct {
//	Value ClarityValue
//}
//
//func (s *SomeCV) isClarityValue() {}
//
//type ListCV struct {
//	Values []ClarityValue
//}
//
//func (l *ListCV) isClarityValue() {}
//
//type TupleCV struct {
//	Values []ClarityValue
//}
//
//func (t *TupleCV) isClarityValue() {}
//
//type StringAsciiCV struct {
//	Value string
//}
//
//func (s *StringAsciiCV) isClarityValue() {}
//
//type StringUtf8CV struct {
//	Value string
//}
//
//func (s *StringUtf8CV) isClarityValue() {}
