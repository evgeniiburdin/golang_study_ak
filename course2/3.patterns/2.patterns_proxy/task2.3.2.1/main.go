package main

import (
	"encoding/json"
	"fmt"
	"github.com/cinar/indicator"
	"io"
	"log"
	"net/http"
)

type Indicator interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
}

type KLines struct {
	Pair    string   `json:"pair"`
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Lines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (t *Lines) StochPrice() ([]float64, []float64) {
	k, d := indicator.StochasticOscillator(t.high, t.low, t.closing)
	return k, d
}

func (t *Lines) RSI(period int) ([]float64, []float64) {
	rs, rsi := indicator.RsiPeriod(period, t.closing)
	return rs, rsi
}

func (t *Lines) StochRSI(rsiPeriod int) ([]float64, []float64) {
	_, rsi := t.RSI(rsiPeriod)
	k, d := indicator.StochasticOscillator(rsi, rsi, rsi)
	return k, d
}

func (t *Lines) SMA(period int) []float64 {
	return indicator.Sma(period, t.closing)
}

func (t *Lines) MACD() ([]float64, []float64) {
	macd, signal := indicator.Macd(t.closing)
	return macd, signal
}

func (t *Lines) EMA() []float64 {
	return indicator.Ema(5, t.closing)
}

type LinesProxy struct {
	lines Indicator
	cache map[string]interface{}
}

func NewLinesProxy(lines Indicator) *LinesProxy {
	return &LinesProxy{
		lines: lines,
		cache: make(map[string]interface{}),
	}
}

func (p *LinesProxy) StochPrice() ([]float64, []float64) {
	if k, ok := p.cache["k_stochprice"]; ok {
		if d, ok := p.cache["d_stochprice"]; ok {
			return k.([]float64), d.([]float64)
		}
	}
	k, d := p.lines.StochPrice()
	p.cache["k_stochprice"] = k
	p.cache["d_stochprice"] = d
	return k, d
}

func (p *LinesProxy) RSI(period int) ([]float64, []float64) {
	rsKey := fmt.Sprintf("rs_%d", period)
	rsiKey := fmt.Sprintf("rsi_%d", period)
	if rs, ok := p.cache[rsKey]; ok {
		if rsi, ok := p.cache[rsiKey]; ok {
			return rs.([]float64), rsi.([]float64)
		}
	}
	rs, rsi := p.lines.RSI(period)
	p.cache[rsKey] = rs
	p.cache[rsiKey] = rsi
	return rs, rsi
}

func (p *LinesProxy) StochRSI(rsiPeriod int) ([]float64, []float64) {
	kKey := fmt.Sprintf("k_stochrsi_%d", rsiPeriod)
	dKey := fmt.Sprintf("d_stochrsi_%d", rsiPeriod)
	if k, ok := p.cache[kKey]; ok {
		if d, ok := p.cache[dKey]; ok {
			return k.([]float64), d.([]float64)
		}
	}
	k, d := p.lines.StochRSI(rsiPeriod)
	p.cache[kKey] = k
	p.cache[dKey] = d
	return k, d
}

func (p *LinesProxy) SMA(period int) []float64 {
	key := fmt.Sprintf("sma_%d", period)
	if sma, ok := p.cache[key]; ok {
		return sma.([]float64)
	}
	sma := p.lines.SMA(period)
	p.cache[key] = sma
	return sma
}

func (p *LinesProxy) MACD() ([]float64, []float64) {
	if macd, ok := p.cache["macd"]; ok {
		if signal, ok := p.cache["signal"]; ok {
			return macd.([]float64), signal.([]float64)
		}
	}
	macd, signal := p.lines.MACD()
	p.cache["macd"] = macd
	p.cache["signal"] = signal
	return macd, signal
}

func (p *LinesProxy) EMA() []float64 {
	if ema, ok := p.cache["ema"]; ok {
		return ema.([]float64)
	}
	ema := p.lines.EMA()
	p.cache["ema"] = ema
	return ema
}

func UnmarshalKLines(data []byte) (KLines, error) {
	var r KLines
	err := json.Unmarshal(data, &r)
	return r, err
}

func LoadKlinesProxy(data []byte) *LinesProxy {
	klines, err := UnmarshalKLines(data)
	if err != nil {
		log.Fatal(err)
	}
	t := &Lines{}
	for _, v := range klines.Candles {
		t.closing = append(t.closing, v.C)
		t.low = append(t.low, v.L)
		t.high = append(t.high, v.H)
	}
	return NewLinesProxy(t)
}

func LoadCandles(pair string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.exmo.com/v1.1/candles_history?symbol=%s&resolution=30&from=1703056979&to=1705476839", pair), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	pair := "BTC_USD"
	candles := LoadCandles(pair)
	lines := LoadKlinesProxy(candles)
	lines.RSI(3)
}
