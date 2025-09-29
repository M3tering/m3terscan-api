package util

import (
	"encoding/binary"
	"fmt"
	"m3terscan-api/internal/models"
	"math/big"
	"time"
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

func GetQuarterHoursSinceYear(year int) (int64, error) {
	start := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now().UTC()

	if start.After(now) {
		return 0, fmt.Errorf("year %d is in the future", year)
	}

	diff := now.Unix() - start.Unix() // difference in seconds
	return diff / 900, nil            // 900 seconds = 15 minutes
}

func DaysInYearUntil(year int) int {
	now := time.Now()
	currentYear := now.Year()

	// startDate = Jan 1 of supplied year
	startDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var stopDate time.Time
	if year == currentYear {
		stopDate = now
	} else {
		// End of supplied year = Dec 31, 23:59:59
		stopDate = time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
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
		out[i] = startNonce + i
	}
	return out
}

func GroupByWeek(data models.ReqStruct) [][]models.WeeklyResponse {
	grouped := make([][]struct {
		Node struct {
			Timestamp int64 `json:"timestamp"`
			Payload   struct {
				Energy float64 `json:"energy"`
			} `json:"payload"`
		} `json:"node"`
	}, 52)

	// group items by week
	for _, item := range data.MeterDataPoints {
		ts := item.Node.Timestamp
		if ts == 0 {
			continue
		}

		// convert ms to time.Time
		d := time.UnixMilli(ts)
		_, week := d.ISOWeek()
		if week >= 1 && week <= 52 {
			grouped[week-1] = append(grouped[week-1], item)
		}
	}

	// build 52 weeks
	weeks := make([]models.WeeklyResponse, 52)
	for i, items := range grouped {
		if len(items) == 0 {
			weeks[i] = models.WeeklyResponse{TotalEnergy: 0, Week: i + 1}
			continue
		}

		var totalEnergy float64
		for _, it := range items {
			totalEnergy += it.Node.Payload.Energy
		}

		d := time.UnixMilli(items[0].Node.Timestamp)
		_, week := d.ISOWeek()

		weeks[i] = models.WeeklyResponse{TotalEnergy: totalEnergy, Week: week}
	}

	// split into 4 quarters of 13 weeks
	final := make([][]models.WeeklyResponse, 4)
	for q := range 4 {
		start := q * 13
		end := start + 13
		final[q] = weeks[start:end]
	}

	return final
}
