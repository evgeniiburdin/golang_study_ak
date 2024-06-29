package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockIndicator struct {
	mock.Mock
}

func (m *MockIndicator) StochPrice() ([]float64, []float64) {
	args := m.Called()
	return args.Get(0).([]float64), args.Get(1).([]float64)
}

func (m *MockIndicator) RSI(period int) ([]float64, []float64) {
	args := m.Called(period)
	return args.Get(0).([]float64), args.Get(1).([]float64)
}

func (m *MockIndicator) StochRSI(rsiPeriod int) ([]float64, []float64) {
	args := m.Called(rsiPeriod)
	return args.Get(0).([]float64), args.Get(1).([]float64)
}

func (m *MockIndicator) SMA(period int) []float64 {
	args := m.Called(period)
	return args.Get(0).([]float64)
}

func (m *MockIndicator) MACD() ([]float64, []float64) {
	args := m.Called()
	return args.Get(0).([]float64), args.Get(1).([]float64)
}

func (m *MockIndicator) EMA() []float64 {
	args := m.Called()
	return args.Get(0).([]float64)
}

func TestLinesProxy_StochPrice(t *testing.T) {
	mockIndicator := new(MockIndicator)
	mockIndicator.On("StochPrice").Return([]float64{10, 20, 30}, []float64{40, 50, 60})

	proxy := NewLinesProxy(mockIndicator)
	k, d := proxy.StochPrice()

	assert.Equal(t, []float64{10, 20, 30}, k)
	assert.Equal(t, []float64{40, 50, 60}, d)
	assert.Equal(t, []float64{10, 20, 30}, proxy.cache["k_stochprice"])
	assert.Equal(t, []float64{40, 50, 60}, proxy.cache["d_stochprice"])

	k, d = proxy.StochPrice()
	assert.Equal(t, []float64{10, 20, 30}, k)
	assert.Equal(t, []float64{40, 50, 60}, d)

	mockIndicator.AssertNumberOfCalls(t, "StochPrice", 1)
}

func TestLinesProxy_RSI(t *testing.T) {
	mockIndicator := new(MockIndicator)
	mockIndicator.On("RSI", 3).Return([]float64{1, 2, 3}, []float64{4, 5, 6})

	proxy := NewLinesProxy(mockIndicator)
	rs, rsi := proxy.RSI(3)

	assert.Equal(t, []float64{1, 2, 3}, rs)
	assert.Equal(t, []float64{4, 5, 6}, rsi)
	assert.Equal(t, []float64{1, 2, 3}, proxy.cache["rs_3"])
	assert.Equal(t, []float64{4, 5, 6}, proxy.cache["rsi_3"])

	rs, rsi = proxy.RSI(3)
	assert.Equal(t, []float64{1, 2, 3}, rs)
	assert.Equal(t, []float64{4, 5, 6}, rsi)

	mockIndicator.AssertNumberOfCalls(t, "RSI", 1)
}

func TestLinesProxy_StochRSI(t *testing.T) {
	mockIndicator := new(MockIndicator)
	mockIndicator.On("RSI", 3).Return([]float64{1, 2, 3}, []float64{4, 5, 6})
	mockIndicator.On("StochRSI", 3).Return([]float64{7, 8, 9}, []float64{10, 11, 12})

	proxy := NewLinesProxy(mockIndicator)
	k, d := proxy.StochRSI(3)

	assert.Equal(t, []float64{7, 8, 9}, k)
	assert.Equal(t, []float64{10, 11, 12}, d)
	assert.Equal(t, []float64{7, 8, 9}, proxy.cache["k_stochrsi_3"])
	assert.Equal(t, []float64{10, 11, 12}, proxy.cache["d_stochrsi_3"])

	k, d = proxy.StochRSI(3)
	assert.Equal(t, []float64{7, 8, 9}, k)
	assert.Equal(t, []float64{10, 11, 12}, d)

	mockIndicator.AssertNumberOfCalls(t, "RSI", 1)
	mockIndicator.AssertNumberOfCalls(t, "StochRSI", 1)
}

func TestLinesProxy_SMA(t *testing.T) {
	mockIndicator := new(MockIndicator)
	mockIndicator.On("SMA", 3).Return([]float64{13, 14, 15})

	proxy := NewLinesProxy(mockIndicator)
	sma := proxy.SMA(3)

	assert.Equal(t, []float64{13, 14, 15}, sma)
	assert.Equal(t, []float64{13, 14, 15}, proxy.cache["sma_3"])

	sma = proxy.SMA(3)
	assert.Equal(t, []float64{13, 14, 15}, sma)

	mockIndicator.AssertNumberOfCalls(t, "SMA", 1)
}

func TestLinesProxy_MACD(t *testing.T) {
	mockIndicator := new(MockIndicator)
	mockIndicator.On("MACD").Return([]float64{16, 17, 18}, []float64{19, 20, 21})

	proxy := NewLinesProxy(mockIndicator)
	macd, signal := proxy.MACD()

	assert.Equal(t, []float64{16, 17, 18}, macd)
	assert.Equal(t, []float64{19, 20, 21}, signal)
	assert.Equal(t, []float64{16, 17, 18}, proxy.cache["macd"])
	assert.Equal(t, []float64{19, 20, 21}, proxy.cache["signal"])

	macd, signal = proxy.MACD()
	assert.Equal(t, []float64{16, 17, 18}, macd)
	assert.Equal(t, []float64{19, 20, 21}, signal)

	mockIndicator.AssertNumberOfCalls(t, "MACD", 1)
}

func TestLinesProxy_EMA(t *testing.T) {
	mockIndicator := new(MockIndicator)
	mockIndicator.On("EMA").Return([]float64{22, 23, 24})

	proxy := NewLinesProxy(mockIndicator)
	ema := proxy.EMA()

	assert.Equal(t, []float64{22, 23, 24}, ema)
	assert.Equal(t, []float64{22, 23, 24}, proxy.cache["ema"])

	ema = proxy.EMA()
	assert.Equal(t, []float64{22, 23, 24}, ema)

	mockIndicator.AssertNumberOfCalls(t, "EMA", 1)
}
