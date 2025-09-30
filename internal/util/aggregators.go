package util

import (
	"fmt"
	"m3terscan-api/internal/models"
	"time"
)

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

func GroupByMonth(data models.ReqStruct, month int, year int) []models.MonthlyResponse {
	// last day of the month
	t := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	daysInMonth := t.Day()

	// accumulator per day
	energyPerDay := make([]float64, daysInMonth)

	for _, item := range data.MeterDataPoints {
		ts := item.Node.Timestamp
		if ts == 0 {
			continue
		}
		d := time.UnixMilli(ts).UTC()

		if d.Year() == year && int(d.Month()) == month {
			day := d.Day() - 1
			energyPerDay[day] += item.Node.Payload.Energy
		}
	}

	// build result
	result := make([]models.MonthlyResponse, daysInMonth)
	for i := range daysInMonth {
		date := fmt.Sprintf("%d-%d-%d", year, month, i+1)
		result[i] = models.MonthlyResponse{
			Date:   date,
			Energy: int64(energyPerDay[i]),
		}
	}

	return result
}
