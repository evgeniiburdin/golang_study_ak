package main

import (
	"context"
	"fmt"
	"log"
	"student.vkusvill.ru/evgeniiburdin/go-course/course2/3.patterns/3.patterns_strategy/task2.3.3.2/pkg"
	"time"
)

// Exchanger определяет интерфейс для работы с биржей.
type Exchanger interface {
	GetCandlesHistory(pair string, resolution string, start, end time.Time) (pkg.CandlesHistory, error)
}

// GeneralIndicatorer определяет общий интерфейс для индикаторов.
type GeneralIndicatorer interface {
	GetData(ctx context.Context, pair string, period time.Duration, from, to time.Time, indicator Indicatorer) ([]float64, error)
}

// Indicatorer определяет интерфейс для конкретных стратегий индикаторов.
type Indicatorer interface {
	GetData(ctx context.Context, pair string, period time.Duration, from, to time.Time) ([]float64, error)
}

// RealExchanger представляет реальный обменник через API Exmo.
type RealExchanger struct{}

// NewRealExchanger создает новый экземпляр RealExchanger.
func NewRealExchanger() *RealExchanger {
	return &RealExchanger{}
}

// GetCandlesHistory получает историю свечей с биржи Exmo.
func (r *RealExchanger) GetCandlesHistory(pair string, resolution string, start, end time.Time) (pkg.CandlesHistory, error) {
	// Реализация метода для получения свечей через API Exmo
	// (реализация опущена для краткости)
	return pkg.CandlesHistory{}, nil
}

// IndicatorSMA представляет стратегию расчета простого скользящего среднего (SMA).
type IndicatorSMA struct {
	exchanger Exchanger
}

// NewIndicatorSMA создает новый экземпляр IndicatorSMA.
func NewIndicatorSMA(exchanger Exchanger) *IndicatorSMA {
	return &IndicatorSMA{exchanger: exchanger}
}

// GetData рассчитывает простое скользящее среднее (SMA) на основе данных от биржи Exmo.
func (s *IndicatorSMA) GetData(ctx context.Context, pair string, period time.Duration, from, to time.Time) ([]float64, error) {
	candles, err := s.exchanger.GetCandlesHistory(pair, "1h", from, to)
	if err != nil {
		return nil, err
	}

	// Расчет SMA на основе свечей
	sma := calculateSMA(candles, int(period.Hours()))

	return sma, nil
}

// calculateSMA рассчитывает простое скользящее среднее (SMA) на основе данных.
func calculateSMA(candles pkg.CandlesHistory, period int) []float64 {
	if period <= 0 || len(candles.Candles) < period {
		return nil
	}

	data := make([]float64, len(candles.Candles)-period+1)
	for i := 0; i < len(data); i++ {
		sum := 0.0
		for _, candle := range candles.Candles[i : i+period] {
			sum += candle.C // Используем закрытие свечи для SMA
		}
		data[i] = sum / float64(period)
	}
	return data
}

// IndicatorEMA представляет стратегию расчета экспоненциального скользящего среднего (EMA).
type IndicatorEMA struct {
	exchanger Exchanger
}

// NewIndicatorEMA создает новый экземпляр IndicatorEMA.
func NewIndicatorEMA(exchanger Exchanger) *IndicatorEMA {
	return &IndicatorEMA{exchanger: exchanger}
}

// GetData рассчитывает экспоненциальное скользящее среднее (EMA) на основе данных от биржи Exmo.
func (e *IndicatorEMA) GetData(ctx context.Context, pair string, period time.Duration, from, to time.Time) ([]float64, error) {
	candles, err := e.exchanger.GetCandlesHistory(pair, "1h", from, to)
	if err != nil {
		return nil, err
	}

	// Расчет EMA на основе свечей
	ema := calculateEMA(candles, int(period.Hours()))

	return ema, nil
}

// calculateEMA рассчитывает экспоненциальное скользящее среднее (EMA) на основе данных.
func calculateEMA(candles pkg.CandlesHistory, period int) []float64 {
	if len(candles.Candles) == 0 || period <= 0 {
		return nil
	}

	alpha := 2.0 / float64(period+1)
	ema := make([]float64, len(candles.Candles))
	ema[0] = candles.Candles[0].C

	for i := 1; i < len(candles.Candles); i++ {
		ema[i] = alpha*candles.Candles[i].C + (1-alpha)*ema[i-1]
	}

	return ema
}

// GeneralIndicator представляет общий индикатор для использования выбранного индикатора.
type GeneralIndicator struct {
	exchanger Exchanger
}

// NewGeneralIndicator создает новый экземпляр GeneralIndicator.
func NewGeneralIndicator(exchanger Exchanger) *GeneralIndicator {
	return &GeneralIndicator{exchanger: exchanger}
}

// GetData вызывает метод конкретного индикатора для получения данных.
func (g *GeneralIndicator) GetData(ctx context.Context, pair string, period time.Duration, from, to time.Time, indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(ctx, pair, period, from, to)
}

func main() {
	exchange := NewRealExchanger()
	generalIndicator := NewGeneralIndicator(exchange)

	// Пример использования SMA
	smaIndicator := NewIndicatorSMA(exchange)
	smaData, err := generalIndicator.GetData(context.Background(), "BTC_USD", 24*time.Hour, time.Now().Add(-30*time.Hour), time.Now(), smaIndicator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SMA Data:", smaData)

	// Пример использования EMA
	emaIndicator := NewIndicatorEMA(exchange)
	emaData, err := generalIndicator.GetData(context.Background(), "BTC_USD", 24*time.Hour, time.Now().Add(-30*time.Hour), time.Now(), emaIndicator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("EMA Data:", emaData)
}
