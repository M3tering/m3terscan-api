package models

type DataStruct struct {
	MeterDataPoints []struct {
		Node struct {
			Timestamp int64 `json:"timestamp"`
			Payload   struct {
				Energy    float64 `json:"energy"`
				Nonce     int     `json:"nonce"`
				Signature string  `json:"signature"`
			} `json:"payload"`
		} `json:"node"`
	} `json:"meterDataPoints"`
}

type ReqStruct struct {
	MeterDataPoints []struct {
		Node struct {
			Timestamp int64 `json:"timestamp"`
			Payload   struct {
				Energy float64 `json:"energy"`
			} `json:"payload"`
		} `json:"node"`
	} `json:"meterDataPoints"`
}
