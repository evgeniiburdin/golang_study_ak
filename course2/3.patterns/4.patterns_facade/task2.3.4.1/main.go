package main

import (
	"fmt"
)

// Interface definitions
type Indicatorer interface {
	SMA(period int) ([]float64, error)
	EMA(period int) ([]float64, error)
}

type Dashboarder interface {
	GetDashboard(pair string, opts ...*IndicatorOpt) (*DashboardData, error)
}

type Exchanger interface {
	LoadCandles(candles CandlesHistory) error
	// Add other methods as needed
}

// Data structures
type CandlesHistory struct {
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

type DashboardData struct {
	Name           string
	CandlesHistory CandlesHistory
	Indicators     []IndicatorData
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator Indicatorer
}

// Dashboard implementation
type Dashboard struct {
	exchange Exchanger
}

func NewDashboard(exchange Exchanger) *Dashboard {
	return &Dashboard{exchange: exchange}
}

func (d *Dashboard) GetDashboard(pair string, opts ...*IndicatorOpt) (*DashboardData, error) {
	// Mock implementation for demonstration
	candles := CandlesHistory{
		Candles: []Candle{
			{T: 1624978800, O: 35000.0, C: 36000.0, H: 36500.0, L: 34000.0, V: 100.0},
			// Add more candles as needed
		},
	}

	indicators := make([]IndicatorData, 0)

	for _, opt := range opts {
		for _, period := range opt.Periods {
			// Mock computation of indicators
			indicatorValues := make([]float64, 0)
			for i := 0; i < 10; i++ {
				indicatorValues = append(indicatorValues, float64(i)*10.0)
			}
			indicator := IndicatorData{
				Name:     opt.Name,
				Period:   period,
				Indicate: indicatorValues,
			}
			indicators = append(indicators, indicator)
		}
	}

	dashboardData := &DashboardData{
		Name:           pair,
		CandlesHistory: candles,
		Indicators:     indicators,
	}

	return dashboardData, nil
}

// Mock Exchange implementation
type Exchange struct{}

func (e *Exchange) LoadCandles(candles CandlesHistory) error {
	// Mock implementation to load candles
	fmt.Println("Loading candles...")
	return nil
}

// Mock Indicator implementation
type Indicator struct{}

func (i *Indicator) SMA(period int) ([]float64, error) {
	// Mock implementation to compute SMA
	fmt.Printf("Computing SMA for period %d...\n", period)
	return []float64{100.0, 110.0, 120.0}, nil
}

func (i *Indicator) EMA(period int) ([]float64, error) {
	// Mock implementation to compute EMA
	fmt.Printf("Computing EMA for period %d...\n", period)
	return []float64{95.0, 105.0, 115.0}, nil
}

func main() {
	// Create a new dashboard instance with a mock exchange
	exchange := &Exchange{}
	dashboard := NewDashboard(exchange)

	// Define indicator options
	indicatorOpts := []*IndicatorOpt{
		{
			Name:      "SMA",
			Periods:   []int{5, 10, 20},
			Indicator: &Indicator{}, // Using mock indicator implementation
		},
		{
			Name:      "EMA",
			Periods:   []int{5, 10, 20},
			Indicator: &Indicator{}, // Using mock indicator implementation
		},
	}

	// Call GetDashboard method
	data, err := dashboard.GetDashboard("BTC_USD", indicatorOpts...)
	if err != nil {
		panic(err)
	}

	// Print the dashboard data
	fmt.Println("Dashboard Data:")
	fmt.Println("Pair:", data.Name)
	fmt.Println("Candles:")
	for _, candle := range data.CandlesHistory.Candles {
		fmt.Printf("  Timestamp: %d, Open: %.2f, Close: %.2f, High: %.2f, Low: %.2f, Volume: %.2f\n",
			candle.T, candle.O, candle.C, candle.H, candle.L, candle.V)
	}
	fmt.Println("Indicators:")
	for _, indicator := range data.Indicators {
		fmt.Println("  Indicator:", indicator.Name, "Period:", indicator.Period, "Values:", indicator.Indicate)
	}
}
