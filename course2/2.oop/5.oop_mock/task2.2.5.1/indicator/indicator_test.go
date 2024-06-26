package indicator

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

type MockExchanger struct {
	ctrl     *gomock.Controller
	recorder *MockExchangerMockRecorder
}

type MockExchangerMockRecorder struct {
	mock *MockExchanger
}

func NewMockExchanger(ctrl *gomock.Controller) *MockExchanger {
	mock := &MockExchanger{ctrl: ctrl}
	mock.recorder = &MockExchangerMockRecorder{mock}
	return mock
}

func (m *MockExchanger) EXPECT() *MockExchangerMockRecorder {
	return m.recorder
}

func (m *MockExchanger) GetClosePrice(pair string, resolution, start, end int64) ([]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClosePrice", pair, resolution, start, end)
	ret0, _ := ret[0].([]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockExchangerMockRecorder) GetClosePrice(pair, resolution interface{}, start, end int64) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClosePrice", reflect.TypeOf((*MockExchanger)(nil).GetClosePrice), pair, resolution, start, end)
}

func TestIndicator_SMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExchanger := NewMockExchanger(ctrl)

	mockExchanger.EXPECT().GetClosePrice("BTC_USD", 30, gomock.Any(), gomock.Any()).Return([]float64{1, 2, 3, 4, 5}, nil)

	ind := NewIndicator(mockExchanger)

	sma, err := ind.SMA("BTC_USD", 30, 3, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []float64{2, 3, 4}
	for i, v := range expected {
		if sma[i] != v {
			t.Errorf("expected %v, got %v", v, sma[i])
		}
	}
}

func TestIndicator_EMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExchanger := NewMockExchanger(ctrl)

	mockExchanger.EXPECT().GetClosePrice("BTC_USD", 30, gomock.Any(), gomock.Any()).Return([]float64{1, 2, 3, 4, 5}, nil)

	ind := NewIndicator(mockExchanger)

	ema, err := ind.EMA("BTC_USD", 30, 3, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []float64{2, 3, 4} // Replace with the actual expected EMA values
	for i, v := range expected {
		if ema[i] != v {
			t.Errorf("expected %v, got %v", v, ema[i])
		}
	}
}
