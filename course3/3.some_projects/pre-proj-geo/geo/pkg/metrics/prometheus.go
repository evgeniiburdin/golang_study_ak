package metrics

import "github.com/prometheus/client_golang/prometheus"

type HTTP struct {
	Handler    string
	Method     string
	StatusCode string
	Duration   float64
}

type MetricsService interface {
	SaveHTTP(h *HTTP)
}

type RealMetricsService struct {
	httpRequestHistogram *prometheus.HistogramVec
}

func NewMetricsService() (*RealMetricsService, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &RealMetricsService{
		httpRequestHistogram: http,
	}

	err := prometheus.Register(s.httpRequestHistogram)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *RealMetricsService) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
