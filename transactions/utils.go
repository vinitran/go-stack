package transactions

import "strings"

func RightPadHexToLength(hexString string, length int) string {
	return strings.Repeat("0", length-len(hexString)) + hexString
}
