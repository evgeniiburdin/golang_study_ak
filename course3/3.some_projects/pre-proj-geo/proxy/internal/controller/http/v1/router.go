package v1

import (
	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"geo-service-proxy/internal/usecase"
	"geo-service-proxy/pkg/logger"
	"geo-service-proxy/pkg/metrics"
)

// NewRouter -.
func NewRouter(httpRouter *chi.Mux, swaggerURL string, lg logger.Interface, uc usecase.Proxyer, mtr metrics.MetricsService) {
	httpRouter.Use(middleware.Logger)
	httpRouter.Use(middleware.Recoverer)
	httpRouter.Use(metrics.NewMiddleware(mtr))

	newProxyRoutes(httpRouter, swaggerURL, uc, lg)
}
