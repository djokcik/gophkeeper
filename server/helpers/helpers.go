package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func hexaNumberToInteger(hexaString string) string {
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func IntToHex(value uint32) string {
	return fmt.Sprintf("%x", value)
}

func HexToInt(numberStr string) (uint32, error) {
	output, err := strconv.ParseUint(hexaNumberToInteger(numberStr), 16, 32)
	if err != nil {
		return 0, err
	}

	return uint32(output), nil
}
