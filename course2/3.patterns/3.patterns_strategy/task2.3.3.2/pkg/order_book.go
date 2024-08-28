package pkg

import "encoding/json"

// OrderBook представляет структуру стакана.
type OrderBook struct {
	Pairs []OrderBookPair `json:"pairs"`
}

// OrderBookPair представляет отдельную пару в стакане.
type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

func UnmarshalOrderBook(data []byte) (OrderBook, error) {
	var r OrderBook
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OrderBook) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
