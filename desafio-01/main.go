package main

import (
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/application"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/infrastructure/env"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/infrastructure/viacepapi"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/infrastructure/weatherapi"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	cfg := env.LoadConfig("./.env")
	viacepApi := viacepapi.NewViaCepAdapter(*cfg)
	weatherApi := weatherapi.NewWeatherAdapter(*cfg)
	useCase := application.NewFindCepAndWeatherUseCase(*viacepApi, *weatherApi)
	controller := api.NewWeatherController(*useCase)

	r := chi.NewRouter()
	r.Route(controller.Path(), controller.Router)
	http.ListenAndServe(":8080", r)
}
