package pkg

import "encoding/json"

// Ticker представляет структуру тикера.
type Ticker struct {
	Values []TickerValue `json:"values"`
}

// TickerValue представляет отдельное значение в тикере.
type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

func UnmarshalTicker(data []byte) (Ticker, error) {
	var r Ticker
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Ticker) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
