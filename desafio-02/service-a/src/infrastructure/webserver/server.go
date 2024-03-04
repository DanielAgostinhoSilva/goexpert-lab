package webserver

import (
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/infrastructure/webserver/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

type WebServer struct {
	controller *api.WeatherController
}

func NewWebServer(controller *api.WeatherController) *WebServer {
	return &WebServer{controller: controller}
}

func (w *WebServer) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	// promhttp
	router.Handle("/metrics", promhttp.Handler())
	router.Route(w.controller.Path(), w.controller.Router)
	return router
}
