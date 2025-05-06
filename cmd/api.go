package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"github.com/weather-app/internal/env"
	// "github.com/weather-app/internal/repository"
	"github.com/weather-app/service"
)

// type Application struct {
// 	Env    *Env
// 	Client *redis.Client
// }

type application struct {
	config config
	// store         store.Storage
	// cacheStorage  cache.Storage
	logger *logrus.Logger
	// redisStore repository.WeatherRepository
	// authenticator auth.Authenticator
	weatherService service.WeatherService
}

type config struct {
	addr string
	// env      string
	apiURL   string
	apiKey   string
	redisCfg redisConfig
	contextTimeout int
}

type redisConfig struct {
	addr string
	pw   string
	db   int
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("CORS_ALLOWED_ORIGIN", "http://localhost:5174")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/health", func(r chi.Router) {
		r.Get("/", app.healthCheckHandler)
	})

	r.Route("/v1", func(r chi.Router) {
		r.Get("/weather", app.weatherHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	fmt.Println("Served has started", "addr", app.config.addr)	

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
