package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FeedistKeyPrefix is the prefix to retrieve all Feedist
	FeedistKeyPrefix = "Feedist/value/"
)

// FeedistKey returns the store key to retrieve a Feedist from the index fields
func FeedistKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
