package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

// Exmo client structure
type Exmo struct {
	client *http.Client
	url    string
}

// NewExmo creates a new Exmo client
func NewExmo(opts ...func(*Exmo)) Exchanger {
	exmo := &Exmo{
		client: &http.Client{},
		url:    "https://api.exmo.com/v1",
	}
	for _, opt := range opts {
		opt(exmo)
	}
	return exmo
}

// WithClient sets a custom http.Client
func WithClient(client *http.Client) func(*Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

// WithURL sets a custom API URL
func WithURL(url string) func(*Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, period int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, resolution int, start, end time.Time) ([]float64, error)
}

type Ticker map[string]TickerValue
type Trades map[string][]Pair
type OrderBook map[string]OrderBookPair
type Currencies []string
type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

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

type Pair struct {
	TradeID  int64  `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

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

type Candle struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
}

// Example function implementations for the Exchanger interface
func (e *Exmo) GetTicker() (Ticker, error) {
	resp, err := e.client.Get(e.url + ticker)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tickerData Ticker
	if err := json.NewDecoder(resp.Body).Decode(&tickerData); err != nil {
		return nil, err
	}
	return tickerData, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	pairsParam := url.QueryEscape(pairs[0])
	resp, err := e.client.Get(e.url + trades + "?pair=" + pairsParam)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tradesData Trades
	if err := json.NewDecoder(resp.Body).Decode(&tradesData); err != nil {
		return nil, err
	}
	return tradesData, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	pairsParam := url.QueryEscape(pairs[0])
	resp, err := e.client.Get(e.url + orderBook + "?pair=" + pairsParam + "&limit=" + strconv.Itoa(limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var orderBookData OrderBook
	if err := json.NewDecoder(resp.Body).Decode(&orderBookData); err != nil {
		return nil, err
	}
	return orderBookData, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	resp, err := e.client.Get(e.url + currency)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var currenciesData Currencies
	if err := json.NewDecoder(resp.Body).Decode(&currenciesData); err != nil {
		return nil, err
	}
	return currenciesData, nil
}

func (e *Exmo) GetCandlesHistory(pair string, period int, start, end time.Time) (CandlesHistory, error) {
	params := url.Values{}
	params.Set("pair", pair)
	params.Set("resolution", strconv.Itoa(period))
	params.Set("from", strconv.FormatInt(start.Unix(), 10))
	params.Set("to", strconv.FormatInt(end.Unix(), 10))

	resp, err := e.client.Get(e.url + candlesHistory + "?" + params.Encode())
	if err != nil {
		return CandlesHistory{}, err
	}
	defer resp.Body.Close()

	var candlesHistoryData CandlesHistory
	if err := json.NewDecoder(resp.Body).Decode(&candlesHistoryData); err != nil {
		return CandlesHistory{}, err
	}
	return candlesHistoryData, nil
}

func (e *Exmo) GetClosePrice(pair string, resolution int, start, end time.Time) ([]float64, error) {
	candlesHistory, err := e.GetCandlesHistory(pair, resolution, start, end)
	if err != nil {
		return nil, err
	}

	closePrices := make([]float64, len(candlesHistory.Candles))
	for i, candle := range candlesHistory.Candles {
		closePrices[i] = candle.Close
	}
	return closePrices, nil
}
