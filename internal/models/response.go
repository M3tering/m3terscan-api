package models

type DailyResponseStruct struct {
	Hour   string  `json:"hour"`
	Energy float64 `json:"energy"`
}
