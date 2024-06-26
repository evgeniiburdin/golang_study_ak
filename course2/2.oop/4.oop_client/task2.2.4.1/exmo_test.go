package exmo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExmo_GetTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"BTC_USD": {"buy_price": "10000", "sell_price": "10100", "last_trade": "10050"}}`))
	}))
	defer server.Close()

	client := NewExmo(WithURL(server.URL))

	ticker, err := client.GetTicker()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if ticker["BTC_USD"].BuyPrice != "10000" {
		t.Errorf("expected 10000, got %v", ticker["BTC_USD"].BuyPrice)
	}
}

func TestExmo_GetTrades(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"BTC_USD": [[{"trade_id": 1, "date": 162314234, "type": "buy", "price": "10000", "quantity": "1", "amount": "10000"}]]}`))
	}))
	defer server.Close()

	client := NewExmo(WithURL(server.URL))

	trades, err := client.GetTrades("BTC_USD")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(trades["BTC_USD"][0]) != 1 {
		t.Errorf("expected 1 trade, got %v", len(trades["BTC_USD"][0]))
	}
}

func TestExmo_GetOrderBook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"BTC_USD": {"ask_quantity": "10", "ask_amount": "100000", "ask_top": "10000", "bid_quantity": "5", "bid_amount": "50000", "bid_top": "9900"}}`))
	}))
	defer server.Close()

	client := NewExmo(WithURL(server.URL))

	orderBook, err := client.GetOrderBook(5, "BTC_USD")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if orderBook["BTC_USD"].AskTop != "10000" {
		t.Errorf("expected 10000, got %v", orderBook["BTC_USD"].AskTop)
	}
}

func TestExmo_GetCurrencies(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`["BTC", "USD", "ETH"]`))
	}))
	defer server.Close()

	client := NewExmo(WithURL(server.URL))

	currencies, err := client.GetCurrencies()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(currencies) != 3 {
		t.Errorf("expected 3 currencies, got %v", len(currencies))
	}
}

func TestExmo_GetCandlesHistory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"candles":[{"timestamp": 162314234, "open": 10000, "close": 10100, "high": 10200, "low": 9900, "volume": 10}]}`))
	}))
	defer server.Close()

	client := NewExmo(WithURL(server.URL))

	candlesHistory, err := client.GetCandlesHistory("BTC_USD", 1, 162314234, 162314234+3600)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(candlesHistory.Candles) != 1 {
		t.Errorf("expected 1 candle, got %v", len(candlesHistory.Candles))
	}
}

func TestExmo_GetClosePrice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"candles":[{"timestamp": 162314234, "open": 10000, "close": 10100, "high": 10200, "low": 9900, "volume": 10}]}`))
	}))
	defer server.Close()

	client := NewExmo(WithURL(server.URL))

	closePrices, err := client.GetClosePrice("BTC_USD", 1, 162314234, 162314234+3600)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(closePrices) != 1 {
		t.Errorf("expected 1 close price, got %v", len(closePrices))
	}

	if closePrices[0] != 10100 {
		t.Errorf("expected 10100, got %v", closePrices[0])
	}
}
