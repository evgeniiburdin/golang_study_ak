package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"geo-service/internal/service"
	"geo-service/models"

	"geo-service/internal/repository/monitoring"
)

type AuthController struct {
	userService service.UserServicer
}

func NewAuthController(userService service.UserServicer) *AuthController {
	return &AuthController{userService: userService}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := c.userService.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})

	timeTaken := time.Since(startTime).Seconds()
	monitoring.LoginRequestsDuration.Observe(timeTaken)
	monitoring.LoginRequestsTotal.Inc()
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.userService.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	timeTaken := time.Since(startTime).Seconds()
	monitoring.RegisterRequestsDuration.Observe(timeTaken)
	monitoring.RegisterRequestsTotal.Inc()
}
