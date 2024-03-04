package main

import (
	"context"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/application"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/infrastructure/env"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/infrastructure/opentelemetry"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/infrastructure/webserver"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/infrastructure/webserver/api"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	cfg := env.LoadConfig("./.env")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := opentelemetry.InitProvider(cfg.OtelServiceName, cfg.OtelExporterOtlpEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TraceProvider: %w", err)
		}
	}()

	tracer := otel.Tracer("microservice-tracer")
	controller := api.NewWeatherController(tracer, application.NewFindWeatherUseCase())
	server := webserver.NewWebServer(controller)
	router := server.CreateServer()

	go func() {
		log.Println("Starting server on port ", cfg.WebServerPort)
		if err := http.ListenAndServe(cfg.WebServerPort, router); err != nil {
			log.Fatal(err)
		}
	}()

	select {
	case <-sigCh:
		log.Println("Shutting down graceful shutdown")
	case <-ctx.Done():
		log.Println("Shutting down due to other reason...")
	}

	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
}
