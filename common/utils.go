package common

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

func ConcatBytes(arrays ...[]byte) []byte {
	if len(arrays) == 0 {
		return nil
	}

	length := 0
	for _, arr := range arrays {
		length += len(arr)
	}

	result := make([]byte, length)
	pad := 0

	for _, arr := range arrays {
		copy(result[pad:], arr)
		pad += len(arr)
	}

	return result
}

func ConcatArray(elements [][]byte) ([]byte, error) {
	var arrays [][]byte
	for _, element := range elements {
		arrays = append(arrays, element)
	}
	return ConcatBytes(arrays...), nil
}

func IntToHex(integer int64, lengthBytes int) string {
	value := big.NewInt(integer)
	hexString := fmt.Sprintf("%x", value)
	paddedHexString := fmt.Sprintf("%0*s", lengthBytes*2, hexString)
	return paddedHexString
}

func BigIntToBytes(value *big.Int, length int) []byte {
	hexString := fmt.Sprintf("%x", value)
	paddedHexString := fmt.Sprintf("%0*s", length*2, hexString)
	bytes, _ := hex.DecodeString(paddedHexString)
	return bytes
}
