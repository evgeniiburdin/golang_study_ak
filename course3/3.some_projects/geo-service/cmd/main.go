package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"geo-service/internal/controller"
	"geo-service/internal/repository"
	"geo-service/internal/repository/monitoring"
	"geo-service/internal/service"
)

var tokenAuth *jwtauth.JWTAuth

func main() {

	prometheus.MustRegister(monitoring.SearchRequestsTotal)
	prometheus.MustRegister(monitoring.GeocodeRequestsTotal)
	prometheus.MustRegister(monitoring.LoginRequestsTotal)
	prometheus.MustRegister(monitoring.RegisterRequestsTotal)
	prometheus.MustRegister(monitoring.SearchRequestsDuration)
	prometheus.MustRegister(monitoring.GeocodeRequestsDuration)
	prometheus.MustRegister(monitoring.LoginRequestsDuration)
	prometheus.MustRegister(monitoring.RegisterRequestsDuration)

	// These are not implemented yet

	//prometheus.MustRegister(monitoring.CacheRequestsDuration)
	//prometheus.MustRegister(monitoring.DBRequestsDuration)
	//prometheus.MustRegister(monitoring.OpenCageAPIRequestsDuration)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpListenerPort := os.Getenv("HTTP_LISTENER_PORT")
	if httpListenerPort == "" {
		httpListenerPort = "8080"
	}
	openCageApiKey := os.Getenv("OPEN_CAGE_API_KEY")
	if openCageApiKey == "" {
		log.Fatal("missing OPEN_CAGE_API_KEY")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("missing JWT_SECRET")
	}
	dbHost, dbPort, dbUser, dbPass, dbName :=
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatal("missing database credentials")
	}
	dbCredentials := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	tokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)

	//userRepo := repository.NewInMemoryUserRepository()

	db, err := sql.Open("postgres", dbCredentials)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := repository.NewPostgresUserRepository(db)
	if err != nil {
		log.Fatal("error connecting to user database: ", err.Error())
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	redisCache := repository.NewRedisCache(redisClient)

	userRepoProxy := repository.NewRedisUserRepositoryProxy(userRepo, redisCache)

	userService := service.NewUserService(userRepoProxy, tokenAuth)
	addressService := service.NewGeoService(openCageApiKey)

	httpRouter := chi.NewRouter()

	httpRouter.Use(middleware.Logger)
	httpRouter.Use(middleware.Recoverer)

	// Public routes
	authController := controller.NewAuthController(userService)
	httpRouter.Post("/api/login", authController.Login)
	httpRouter.Post("/api/register", authController.Register)
	http.Handle("/metrics", promhttp.Handler())

	// Protected routes
	addressController := controller.NewAddressController(addressService)
	httpRouter.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", addressController.Search)
		r.Post("/api/address/geocode", addressController.Geocode)
	})

	/*
		// Регистрация pprof-обработчиков
		httpRouter.HandleFunc("/debug/pprof/", pprof.Index)
		httpRouter.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		httpRouter.HandleFunc("/debug/pprof/profile", pprof.Profile)
		httpRouter.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		httpRouter.HandleFunc("/debug/pprof/trace", pprof.Trace)
	*/

	log.Println("redis: ", redisClient.Ping())
	log.Println("database connection error: ", db.Ping())

	log.Println("Starting server at port " + httpListenerPort)
	log.Fatal(http.ListenAndServe(":"+httpListenerPort, httpRouter))
}
