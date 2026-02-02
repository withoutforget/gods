package clibutils

import (
	"bytes"
	"encoding/gob"
	"hash/maphash"
)

// hash function has been taken from claude cuz don't wanna think
// about hashing
func HashFunction[K comparable](seed maphash.Seed, key K) int {

	switch k := any(key).(type) {
	case string:
		return int(maphash.String(seed, k))
	case int:
		return int(uint64(k) * 11400714819323198485)
	case int64:
		return int(uint64(k) * 11400714819323198485)
	default:
		return int(hashViaSerialization(seed, key))
	}

}
func hashViaSerialization[K comparable](seed maphash.Seed, key K) uint64 {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(key)
	return maphash.Bytes(seed, buf.Bytes())
}
