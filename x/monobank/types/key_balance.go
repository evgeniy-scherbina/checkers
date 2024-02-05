package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BalanceKeyPrefix is the prefix to retrieve all Balance
	BalanceKeyPrefix = "Balance/value/"
)

// BalanceKey returns the store key to retrieve a Balance from the index fields
func BalanceKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
