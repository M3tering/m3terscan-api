package util

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

// Convert [6]byte -> uint64
func Bytes6ToInt64(b [6]byte) int64 {
	// Pad to 8 bytes
	var buf [8]byte
	copy(buf[2:], b[:]) // left-pad with 0s, put b in the last 6 slots
	return int64(binary.BigEndian.Uint64(buf[:]))
}

// Convert [6]byte -> *big.Int (safer if nonce can be large)
func Bytes6ToBigInt(b [6]byte) *big.Int {
	return new(big.Int).SetBytes(b[:])
}

func GetQuarterHoursSinceYear(year int) (int64, error) {
	start := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now().UTC()

	if start.After(now) {
		return 0, fmt.Errorf("year %d is in the future", year)
	}

	diff := now.Unix() - start.Unix() // difference in seconds
	return diff / 900, nil            // 900 seconds = 15 minutes
}

func DaysUntil(year int, month *int) int {
	val := 12
	now := time.Now()
	currentYear := now.Year()

	// startDate = Jan 1 of supplied year
	startDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	if month != nil {
		val = *month
	}
	var stopDate time.Time
	if year == currentYear {
		stopDate = now
	} else {
		// End of supplied year = Dec 31, 23:59:59
		stopDate = time.Date(year, time.Month(val), 31, 23, 59, 59, 0, time.UTC)
	}

	// inclusive days
	days := int(stopDate.Sub(startDate).Hours()/24) + 1
	return days
}

func NonceRange(startNonce, endNonce int) []int {
	length := endNonce - startNonce + 1
	if length <= 0 {
		return []int{}
	}

	out := make([]int, length)
	for i := range length {
		val := max(startNonce+i, 0)
		out[i] = val
	}
	return out
}
