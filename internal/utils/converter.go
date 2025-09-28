package utils

import (
	"encoding/binary"
	"math/big"
)

// Convert [6]byte -> uint64
func Bytes6ToUint64(b [6]byte) uint64 {
	// Pad to 8 bytes
	var buf [8]byte
	copy(buf[2:], b[:]) // left-pad with 0s, put b in the last 6 slots
	return binary.BigEndian.Uint64(buf[:])
}

// Convert [6]byte -> *big.Int (safer if nonce can be large)
func Bytes6ToBigInt(b [6]byte) *big.Int {
	return new(big.Int).SetBytes(b[:])
}
