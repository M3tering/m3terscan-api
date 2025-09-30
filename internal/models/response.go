package models

type DailyResponseStruct struct {
	Hour   string  `json:"hour"`
	Energy float64 `json:"energy"`
}

type WeeklyResponse struct {
	Week        int     `json:"week"`
	TotalEnergy float64 `json:"totalEnergy"`
}

type MonthlyResponse struct {
	Date   string `json:"date"`
	Energy int64  `json:"energy"`
}
