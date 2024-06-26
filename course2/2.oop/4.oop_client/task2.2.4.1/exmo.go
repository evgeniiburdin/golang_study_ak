package exmo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	baseURL        = "https://api.exmo.com/v1"
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type Exmo struct {
	client *http.Client
	url    string
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end int64) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end int64) ([]float64, error)
}

func NewExmo(opts ...func(*Exmo)) *Exmo {
	e := &Exmo{
		client: &http.Client{Timeout: 10 * time.Second},
		url:    baseURL,
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func WithClient(client *http.Client) func(*Exmo) {
	return func(e *Exmo) {
		e.client = client
	}
}

func WithURL(url string) func(*Exmo) {
	return func(e *Exmo) {
		e.url = url
	}
}

type Ticker map[string]TickerValue

type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
}

type Trades map[string][][]Trade

type Trade struct {
	TradeID  int    `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Amount   string `json:"amount"`
}

type OrderBook map[string]OrderBookPair

type OrderBookPair struct {
	AskQuantity string `json:"ask_quantity"`
	AskAmount   string `json:"ask_amount"`
	AskTop      string `json:"ask_top"`
	BidQuantity string `json:"bid_quantity"`
	BidAmount   string `json:"bid_amount"`
	BidTop      string `json:"bid_top"`
}

type Currencies []string

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
}

func (e *Exmo) GetTicker() (Ticker, error) {
	resp, err := e.client.Get(e.url + ticker)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ticker Ticker
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		return nil, err
	}
	return ticker, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	pairsParam := strings.Join(pairs, ",")
	resp, err := e.client.Get(fmt.Sprintf("%s%s?pair=%s", e.url, trades, pairsParam))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var trades Trades
	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
		return nil, err
	}
	return trades, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	pairsParam := strings.Join(pairs, ",")
	resp, err := e.client.Get(fmt.Sprintf("%s%s?limit=%d&pair=%s", e.url, orderBook, limit, pairsParam))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var orderBook OrderBook
	if err := json.NewDecoder(resp.Body).Decode(&orderBook); err != nil {
		return nil, err
	}
	return orderBook, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	resp, err := e.client.Get(e.url + currency)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var currencies Currencies
	if err := json.NewDecoder(resp.Body).Decode(&currencies); err != nil {
		return nil, err
	}
	return currencies, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end int64) (CandlesHistory, error) {
	resp, err := e.client.Get(fmt.Sprintf("%s%s?pair=%s&limit=%d&from=%d&to=%d", e.url, candlesHistory, pair, limit, start, end))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var candlesHistory CandlesHistory
	if err := json.NewDecoder(resp.Body).Decode(&candlesHistory); err != nil {
		return nil, err
	}
	return candlesHistory, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end int64) ([]float64, error) {
	candlesHistory, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return nil, err
	}

	closePrices := make([]float64, len(candlesHistory.Candles))
	for i, candle := range candlesHistory.Candles {
		closePrices[i] = candle.Close
	}
	return closePrices, nil
}
