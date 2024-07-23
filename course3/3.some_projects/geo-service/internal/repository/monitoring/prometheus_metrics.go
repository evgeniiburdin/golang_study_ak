package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	SearchRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "search_requests_total",
		Help: "The total number of search requests",
	})

	GeocodeRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "geocode_requests_total",
		Help: "The total number of geocode requests",
	})
	LoginRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "login_requests_total",
		Help: "The total number of login requests",
	})
	RegisterRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "register_requests_total",
		Help: "The total number of register requests",
	})

	SearchRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "search_requests_duration",
		Help:    "The time taken for search requests",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	GeocodeRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "geocode_requests_duration",
		Help:    "The time taken for geocode requests",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	LoginRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "login_requests_duration",
		Help:    "The time taken for login requests",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	RegisterRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "register_requests_duration",
		Help:    "The time taken for register requests",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	CacheRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "cache_request_duration",
		Help:    "The time taken for cache requests",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	DBRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "db_request_duration",
		Help:    "The time taken for db requests",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	OpenCageAPIRequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "opencage_api_request_duration",
		Help:    "The time taken for opencage api requests",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
)
