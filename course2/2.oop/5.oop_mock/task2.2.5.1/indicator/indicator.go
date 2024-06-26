package indicator

import (
	"github.com/cinar/indicator"
	"time"
)

type Exchanger interface {
	GetClosePrice(pair string, resolution, start, end int64) ([]float64, error)
}

type Indicatorer interface {
	SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
}

type Indicator struct {
	exchange     Exchanger
	calculateSMA func([]float64, int) []float64
	calculateEMA func([]float64, int) []float64
}

func NewIndicator(exchange Exchanger, opts ...func(*Indicator)) *Indicator {
	ind := &Indicator{
		exchange:     exchange,
		calculateSMA: calculateSMA,
		calculateEMA: calculateEMA,
	}
	for _, opt := range opts {
		opt(ind)
	}
	return ind
}

func WithCalculateSMA(calculateSMA func([]float64, int) []float64) func(*Indicator) {
	return func(ind *Indicator) {
		ind.calculateSMA = calculateSMA
	}
}

func WithCalculateEMA(calculateEMA func([]float64, int) []float64) func(*Indicator) {
	return func(ind *Indicator) {
		ind.calculateEMA = calculateEMA
	}
}

func calculateSMA(data []float64, period int) []float64 {
	return indicator.Sma(period, data)
}

func calculateEMA(data []float64, period int) []float64 {
	return indicator.Ema(period, data)
}

func (ind *Indicator) SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	start := from.Unix()
	end := to.Unix()
	data, err := ind.exchange.GetClosePrice(pair, resolution, start, end)
	if err != nil {
		return nil, err
	}
	return ind.calculateSMA(data, period), nil
}

func (ind *Indicator) EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	start := from.Unix()
	end := to.Unix()
	data, err := ind.exchange.GetClosePrice(pair, resolution, start, end)
	if err != nil {
		return nil, err
	}
	return ind.calculateEMA(data, period), nil
}
