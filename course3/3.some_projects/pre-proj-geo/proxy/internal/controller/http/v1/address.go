package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"geo-service-proxy/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "net/http/pprof"

	_ "geo-service-proxy/docs"
	"geo-service-proxy/internal/usecase"
	"geo-service-proxy/pkg/logger"
)

type proxyRoutes struct {
	uc usecase.Proxyer
	lg logger.Interface
}

func newProxyRoutes(handler *chi.Mux, swaggerURL string, uc usecase.Proxyer, lg logger.Interface) {
	routes := &proxyRoutes{
		uc, lg,
	}

	handler.Route("/api/address", func(router chi.Router) {
		router.Post("/geocode", routes.GeocodeToAddress)
	})
	handler.Route("/api/auth", func(router chi.Router) {
		router.Post("/register", routes.CreateUser)
		router.Get("/user", routes.GetUser)
		router.Get("/users", routes.GetUsers)
		router.Post("/user/update", routes.UpdateUser)
		router.Get("/user/delete", routes.DeleteUser)
		router.Post("/login", routes.Login)
		router.Get("/logout", routes.Logout)
		router.Get("/token/validate", routes.ValidateToken)
		router.Get("/token/refresh", routes.RefreshToken)
	})
	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))
	handler.Get("/metrics", handlerToHandlerFunc(promhttp.Handler()))
}

type (
	CreateUserRequest struct {
		User entity.User `json:"user"`
	}

	CreateUserResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	GetUserRequest struct {
		AccessToken string `json:"access_token"`
	}

	GetUserResponse struct {
		User entity.User `json:"user"`
	}

	GetUsersRequest struct {
		AccessToken string `json:"access_token"`
	}

	GetUsersResponse struct {
		Users []entity.User `json:"users"`
	}

	UpdateUserRequest struct {
		User entity.User `json:"user"`
	}

	DeleteUserRequest struct {
		AccessToken string `json:"access_token"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	ValidateTokenRequest struct {
		Token string `json:"token"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenResponse struct {
		NewAccessToken  string `json:"new_access_token"`
		NewRefreshToken string `json:"new_refresh_token"`
	}

	GeocodeToAddressRequest struct {
		Lat string `json:"lat"`
		Lng string `json:"lon"`
	}

	GeocodeToAddressResponse struct {
		Country string `json:"country"`
		City    string `json:"city"`
	}

	ErrorResponse struct {
		Message string `json:"message"`
	}
)

// @Summary     Register a new user
// @Description Creates a new user in the system
// @ID          CreateUser
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       user body CreateUserRequest true "User data"
// @Success     200 {string} object CreateUserResponse
// @Failure     400 {object} ErrorResponse "invalid request body"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/register [post]
func (routes *proxyRoutes) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid request body")
		routes.lg.Error(fmt.Errorf("http - CreateUser - invalid request body: %+v", r.Body))
		return
	}

	ctx := context.TODO()

	accessToken, refreshToken, err := routes.uc.CreateUser(ctx, req.User)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		routes.lg.Error(fmt.Errorf("http - CreateUser - usecase.CreateUser: %+v", err.Error()))
		return
	}

	resp := &CreateUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// @Summary     Get user information
// @Description Retrieves information about the currently authenticated user
// @ID          GetUser
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "token"
// @Success     200 {object} GetUserResponse
// @Failure     401 {object} ErrorResponse "invalid token"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/user [get]
func (routes *proxyRoutes) GetUser(w http.ResponseWriter, r *http.Request) {
	req := GetUserRequest{
		AccessToken: r.Header.Get("Authorization"),
	}

	if req.AccessToken == "" {
		errorResponse(w, http.StatusUnauthorized, "access token is required")
		routes.lg.Error(fmt.Errorf("http - GetUser - access token is required"))
		return
	}

	ctx := context.TODO()

	user, err := routes.uc.GetUser(ctx, req.AccessToken)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - GetUser - usecase.GetUser: %w", err.Error()))
		return
	}

	resp := &GetUserResponse{
		User: user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// @Summary     Get all users
// @Description Retrieves a list of all registered users
// @ID          GetUsers
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "token"
// @Success     200 {object} GetUsersResponse
// @Failure     401 {object} ErrorResponse "invalid token"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/users [get]
func (routes *proxyRoutes) GetUsers(w http.ResponseWriter, r *http.Request) {
	req := GetUsersRequest{
		AccessToken: r.Header.Get("Authorization"),
	}

	if req.AccessToken == "" {
		errorResponse(w, http.StatusUnauthorized, "access token is required")
		routes.lg.Error(fmt.Errorf("http - GetUsers - access token is required"))
		return
	}

	ctx := context.TODO()

	users, err := routes.uc.GetUsers(ctx, req.AccessToken)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - GetUsers - usecase.GetUsers: %w", err.Error()))
		return
	}

	resp := &GetUsersResponse{
		Users: users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// @Summary     Update user information
// @Description Updates information for the currently authenticated user
// @ID          UpdateUser
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "token"
// @Param       user body UpdateUserRequest true "Updated user data"
// @Success     200 {string} string "ok"
// @Failure     400 {object} ErrorResponse "invalid request body"
// @Failure     401 {object} ErrorResponse "invalid token"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/user/update [post]
func (routes *proxyRoutes) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid request body")
		routes.lg.Error(fmt.Errorf("http - UpdateUser - invalid request body: %+v", r.Body))
		return
	}

	req := &UpdateUserRequest{
		User: user,
	}

	accessToken := r.Header.Get("Authorization")

	if accessToken == "" {
		errorResponse(w, http.StatusBadRequest, "access token is required")
		routes.lg.Error(fmt.Errorf("access token is required"))
		return
	}

	ctx := context.TODO()

	routes.lg.Info("sending request to update user with user = %+v", req)
	err = routes.uc.UpdateUser(ctx, accessToken, user)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - UpdateUser - usecase.UpdateUser: %w", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary     Delete user
// @Description Deletes the currently authenticated user
// @ID          DeleteUser
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "token"
// @Success     200 {string} string "ok"
// @Failure     400 {object} ErrorResponse "access token required"
// @Failure     401 {object} ErrorResponse "invalid token"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/user/delete [get]
func (routes *proxyRoutes) DeleteUser(w http.ResponseWriter, r *http.Request) {
	req := DeleteUserRequest{
		AccessToken: r.Header.Get("Authorization"),
	}

	if req.AccessToken == "" {
		errorResponse(w, http.StatusBadRequest, "access token required")
		routes.lg.Error(fmt.Errorf("http - DeleteUser - access token required"))
		return
	}

	ctx := context.TODO()

	err := routes.uc.DeleteUser(ctx, req.AccessToken)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - DeleteUser - usecase.DeleteUser: %w", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary     Login user
// @Description Authenticates a user and returns access and refresh tokens
// @ID          Login
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       user body LoginRequest true "User credentials"
// @Success     200 {object} LoginResponse
// @Failure     400 {object} ErrorResponse "invalid request body"
// @Failure     401 {object} ErrorResponse "invalid credentials"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/login [post]
func (routes *proxyRoutes) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid request body")
		routes.lg.Error(fmt.Errorf("http - Login - invalid request body: %+v", r.Body))
		return
	}

	ctx := context.TODO()

	accessToken, refreshToken, err := routes.uc.Login(ctx, req.Email, req.Password)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - Login - usecase.Login: %w", err.Error()))
		return
	}

	resp := &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

type logoutRequest struct {
	AccessToken string `json:"access_token"`
}

// @Summary     Logout user
// @Description Logs out the currently authenticated user
// @ID          Logout
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "token"
// @Success     200 {string} string "ok"
// @Failure     400 {object} ErrorResponse "access token required"
// @Failure     401 {object} ErrorResponse "invalid token"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/logout [get]
func (routes *proxyRoutes) Logout(w http.ResponseWriter, r *http.Request) {
	req := logoutRequest{
		AccessToken: r.Header.Get("Authorization"),
	}

	if req.AccessToken == "" {
		errorResponse(w, http.StatusBadRequest, "access token is required")
		routes.lg.Error(fmt.Errorf("http - Logout - access token is required"))
		return
	}

	ctx := context.TODO()

	err := routes.uc.Logout(ctx, req.AccessToken)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - Logout - usecase.Logout: %w", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary     Validate token
// @Description Validates the provided access token
// @ID          ValidateToken
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "token"
// @Success     200 {string} string "valid"
// @Failure     400 {object} ErrorResponse "token is required"
// @Failure     401 {object} ErrorResponse "invalid token"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/token/validate [get]
func (routes *proxyRoutes) ValidateToken(w http.ResponseWriter, r *http.Request) {
	req := ValidateTokenRequest{
		Token: r.Header.Get("Authorization"),
	}

	if req.Token == "" {
		errorResponse(w, http.StatusBadRequest, "token is required")
		routes.lg.Error(fmt.Errorf("http - ValidateToken - token is required"))
		return
	}

	ctx := context.TODO()

	err := routes.uc.ValidateToken(ctx, req.Token)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - ValidateToken - usecase.ValidateToken: %w", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary     Refresh token
// @Description Refreshes the access and refresh tokens using the provided refresh token
// @ID          RefreshToken
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "token"
// @Success     200 {object} RefreshTokenResponse
// @Failure     400 {object} ErrorResponse "invalid request body"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/auth/token/refresh [get]
func (routes *proxyRoutes) RefreshToken(w http.ResponseWriter, r *http.Request) {
	req := RefreshTokenRequest{
		RefreshToken: r.Header.Get("Authorization"),
	}
	if req.RefreshToken == "" {
		errorResponse(w, http.StatusBadRequest, "invalid request body")
		routes.lg.Error(fmt.Errorf("http - RefreshToken - refresh token required"))
		return
	}

	ctx := context.TODO()

	accessToken, refreshToken, err := routes.uc.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		routes.lg.Error(fmt.Errorf("http - RefreshToken - failed to refresh token: %w", err.Error()))
		return
	}

	resp := &RefreshTokenResponse{
		NewAccessToken:  accessToken,
		NewRefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// @Summary     Convert geocode to address
// @Description Converts latitude and longitude to a human-readable address
// @ID          GeocodeToAddress
// @Tags        address
// @Accept      json
// @Produce     json
// @Param       lat query string true "Latitude"
// @Param       lng query string true "Longitude"
// @Success     200 {object} GeocodeToAddressResponse
// @Failure     400 {object} ErrorResponse "invalid request parameters"
// @Failure		401 {object} ErrorResponse "invalid token"
// @Failure     500 {object} ErrorResponse "internal server error"
// @Router      /api/address/geocode [post]
func (routes *proxyRoutes) GeocodeToAddress(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Authorization")

	req := &GeocodeToAddressRequest{
		Lat: r.URL.Query().Get("lat"),
		Lng: r.URL.Query().Get("lng"),
	}
	if req.Lat == "" || req.Lng == "" {
		errorResponse(w, http.StatusBadRequest, "invalid request parameters")
		routes.lg.Error(fmt.Errorf("http - GeocodeToAddress - invalid request parameters: lat: %s lng: %s", req.Lat, req.Lng))
		return
	}

	ctx := context.TODO()

	err := routes.uc.ValidateToken(ctx, accessToken)
	if err != nil {
		errorResponse(w, http.StatusUnauthorized, err.Error())
		routes.lg.Error(fmt.Errorf("http - GeocodeToAddress - invalid token: %s", err.Error()))
		return
	}

	var address entity.Address
	address, err = routes.uc.GeocodeToAddress(ctx, entity.Geocode{
		Lat: req.Lat,
		Lng: req.Lng,
	})
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		routes.lg.Error(fmt.Errorf("http - GeocodeToAddress - usecase.GeocodeToAddress: %s", err.Error()))
		return
	}

	resp := &GeocodeToAddressResponse{
		Country: address.Country,
		City:    address.City,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func handlerToHandlerFunc(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
