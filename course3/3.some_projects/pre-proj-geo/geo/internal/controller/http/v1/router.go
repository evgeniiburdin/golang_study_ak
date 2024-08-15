package v1

import (
	"geo-service/pkg/metrics"
	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"geo-service/internal/usecase"
	"geo-service/pkg/logger"
)

// NewRouter -.
func NewRouter(httpRouter *chi.Mux, swaggerURL string, lg logger.Interface, uc usecase.Addresser, mtr metrics.MetricsService) {
	httpRouter.Use(middleware.Logger)
	httpRouter.Use(middleware.Recoverer)
	httpRouter.Use(metrics.NewMiddleware(mtr))

	newAddressRoutes(httpRouter, swaggerURL, uc, lg)
}
