package transactions

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"strings"
)

func RightPadHexToLength(hexString string, length int) string {
	return fmt.Sprintf("%s%s", strings.Repeat("0", length-len(hexString)), hexString)
}

func Hash160(input []byte) []byte {
	sha256Hash := sha256.Sum256(input)
	ripemd160Hash := ripemd160.New()
	_, _ = ripemd160Hash.Write(sha256Hash[:])
	return ripemd160Hash.Sum(nil)
}

func HashP2PKH(input []byte) string {
	return hex.EncodeToString(Hash160(input))
}

func ExceedsMaxLengthBytes(str string, maxLengthBytes int) bool {
	if str != "" {
		return len([]byte(str)) > maxLengthBytes
	}
	return false
}

func TxidFromData(data []byte) string {
	hash := sha512.New512_256()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

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

func LeftPadHex(hexString string) string {
	if len(hexString)%2 == 0 {
		return hexString
	}
	return fmt.Sprintf("0%s", hexString)
}
