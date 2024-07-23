package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"geo-service/internal/service"
	"geo-service/models"

	"geo-service/internal/repository/monitoring"
)

type AddressController struct {
	geoService service.GeoServicer
}

func NewAddressController(geoService service.GeoServicer) *AddressController {
	return &AddressController{geoService: geoService}
}

func (c *AddressController) Search(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var req models.RequestAddressSearch
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	info, err := c.geoService.SearchAddress(req.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)

	timeTaken := time.Since(startTime).Seconds()
	monitoring.SearchRequestsTotal.Inc()
	monitoring.SearchRequestsDuration.Observe(timeTaken)
}

func (c *AddressController) Geocode(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var req models.RequestAddressGeocode
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	info, err := c.geoService.GeocodeAddress(req.Lat, req.Lng)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)

	timeTaken := time.Since(startTime).Seconds()
	monitoring.GeocodeRequestsTotal.Inc()
	monitoring.GeocodeRequestsDuration.Observe(timeTaken)
}
