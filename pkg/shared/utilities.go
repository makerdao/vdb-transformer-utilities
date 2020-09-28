package shared

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

func BigIntToString(value *big.Int) string {
	result := value.String()
	if result == "<nil>" {
		return ""
	} else {
		return result
	}
}

func ConvertIntStringToHex(n string) (string, error) {
	b := big.NewInt(0)
	b, ok := b.SetString(n, 10)
	if !ok {
		return "", errors.New("error converting int to hex")
	}
	leftPaddedBytes := common.LeftPadBytes(b.Bytes(), 32)
	hex := common.Bytes2Hex(leftPaddedBytes)
	return hex, nil
}

func ConvertInt256HexToBigInt(hex string) *big.Int {
	n := ConvertUint256HexToBigInt(hex)
	return math.S256(n)
}

func ConvertUint256HexToBigInt(hex string) *big.Int {
	hexBytes := common.FromHex(hex)
	return big.NewInt(0).SetBytes(hexBytes)
}

func DecodeHexToText(payload string) string {
	return string(bytes.Trim(common.FromHex(payload), "\x00"))
}

func FormatRollbackError(field string, err error) error {
	return fmt.Errorf("failed to rollback transaction after failing to insert %s: %w", field, err)
}

func GetFullTableName(schema, table string) string {
	return schema + "." + table
}

func GetChecksumAddressString(address string) string {
	return common.HexToAddress(address).Hex()
}
