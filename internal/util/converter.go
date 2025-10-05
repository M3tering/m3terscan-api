package util

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"m3terscan-api/internal/models"
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

func HexToChunks(hexStr string, chunkSize int) ([]*big.Int, error) {
	if chunkSize <= 0 {
		chunkSize = 6
	}

	// Remove "0x" prefix if present
	if len(hexStr) >= 2 && hexStr[:2] == "0x" {
		hexStr = hexStr[2:]
	}

	// Pad with leading "00"
	hexStr = "00" + hexStr

	chunks := []*big.Int{}
	bytesData, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(bytesData); i += chunkSize {
		end := i + chunkSize
		if end > len(bytesData) {
			end = len(bytesData)
		}
		part := bytesData[i:end]

		// Convert to big.Int
		num := new(big.Int).SetBytes(part)
		chunks = append(chunks, num)
	}

	return chunks, nil
}

func CombineAccountsNonces(accounts, nonces []*big.Int) ([]models.StateResponse, error) {
	if len(accounts) != len(nonces) {
		return nil, fmt.Errorf("accounts and nonces must have the same length")
	}

	divisor := big.NewInt(1000000)
	result := make([]models.StateResponse, len(accounts))

	for i := range accounts {
		accountValue := new(big.Int).Div(accounts[i], divisor)
		result[i] = models.StateResponse{
			M3terNo: i + 1,
			Account: accountValue,
			Nonce:   nonces[i],
		}
	}

	return result, nil
}
