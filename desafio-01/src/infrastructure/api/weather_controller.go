package api

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/application"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type WeatherController struct {
	findCepAndWeather application.FindCepAndWeatherUseCase
}

func NewWeatherController(findCepAndWeather application.FindCepAndWeatherUseCase) *WeatherController {
	return &WeatherController{findCepAndWeather: findCepAndWeather}
}

func (wc *WeatherController) FindWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cep := chi.URLParam(r, "cep")
	weatherOutput, problem := wc.findCepAndWeather.Execute(application.CepWeatherInput{
		Cep: cep,
	})
	if problem != nil {
		w.WriteHeader(problem.Status)
		json.NewEncoder(w).Encode(problem)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(weatherOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wc *WeatherController) Router(router chi.Router) {
	router.Get("/", wc.FindWeather)
}
func (wc *WeatherController) Path() string {
	return "/weather/{cep}"
}
