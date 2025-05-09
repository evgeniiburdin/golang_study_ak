package metrics

import (
	"net/http"
	"strconv"
)

func NewMiddleware(mtr MetricsService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := NewResponseWriter(w)
			next.ServeHTTP(rw, r)
			mtr.SaveHTTP(&HTTP{
				Handler:    r.URL.Path,
				Method:     r.Method,
				StatusCode: strconv.Itoa(rw.StatusCode()),
				Duration:   rw.Duration(),
			})
		})
	}
}
