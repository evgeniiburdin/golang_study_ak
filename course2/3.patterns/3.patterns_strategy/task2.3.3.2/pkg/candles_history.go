package pkg

import (
	"encoding/json"
)

// CandlesHistory представляет структуру истории свечей.
type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

// Candle представляет отдельную свечу.
type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

func UnmarshalCandlesHistory(data []byte) (CandlesHistory, error) {
	var r CandlesHistory
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CandlesHistory) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
