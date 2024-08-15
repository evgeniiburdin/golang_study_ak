package v1

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "geo-service/docs"
	"geo-service/internal/entity"
	"geo-service/internal/usecase"
	"geo-service/pkg/logger"
)

type addressRoutes struct {
	uc usecase.Addresser
	lg logger.Interface
}

func newAddressRoutes(handler *chi.Mux, swaggerURL string, uc usecase.Addresser, lg logger.Interface) {
	routes := &addressRoutes{
		uc, lg,
	}

	handler.Route("/api/address", func(router chi.Router) {
		router.Post("/geocode", routes.geocodeToAddress)
	})
	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))
	handler.Get("/metrics", handlerToHandlerFunc(promhttp.Handler()))
}

type geocodeToAddressRequest struct {
	Lat string `json:"lat" binding:"required"`
	Lng string `json:"lng" binding:"required"`
}

type geocodeToAddressResponse struct {
	Country string `json:"country" binding:"required"`
	City    string `json:"city" binding:"required"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// @Summary     Convert geocode to address
// @Description Converts latitude and longitude to a human-readable address
// @ID          geocodeToAddress
// @Tags        address
// @Accept      json
// @Produce     json
// @Param       lat query string true "Latitude"
// @Param       lng query string true "Longitude"
// @Success     200 {object} geocodeToAddressResponse
// @Failure     400 {object} ErrorResponse "invalid request parameters"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/address/geocode [post]
func (routes *addressRoutes) geocodeToAddress(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lng := r.URL.Query().Get("lng")
	if lat == "" || lng == "" {
		routes.lg.Error("http - geocodeToAddress - lat, lng query string required")
		errorResponse(w, http.StatusBadRequest, "query string required")
	}
	_, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		routes.lg.Error(err, fmt.Errorf("http - geocodeToAddress: %w", err))
		errorResponse(w, http.StatusBadRequest, "invalid request params")
	}
	_, err = strconv.ParseFloat(lng, 64)
	if err != nil {
		routes.lg.Error(err, fmt.Errorf("http - geocodeToAddress: %w", err))
		errorResponse(w, http.StatusBadRequest, "invalid request params")
	}

	address, err := routes.uc.GeocodeToAddress(
		r.Context(),
		entity.Geocode{
			Lat: lat,
			Lng: lng,
		},
	)
	if err != nil {
		routes.lg.Error(err, fmt.Errorf("http - geocodeToAddress: %w", err))
		errorResponse(w, http.StatusInternalServerError, "internal server error")

		return
	}

	response := geocodeToAddressResponse{
		Country: address.Country,
		City:    address.City,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handlerToHandlerFunc(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
