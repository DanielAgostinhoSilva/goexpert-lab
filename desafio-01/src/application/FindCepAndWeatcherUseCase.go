package application

import (
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/commons/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/infrastructure/viacepapi"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/src/infrastructure/weatherapi"
)

type CepWeatherInput struct {
	Cep string
}

type CepWeatherOutput struct {
	Location       string  `json:"location"`
	Condition      string  `json:"condition"`
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type FindCepAndWeatherUseCase struct {
	cepAdapter     viacepapi.ViaCepAdapter
	weatherAdapter weatherapi.WeatherAdapter
}

func NewFindCepAndWeatherUseCase(cepAdapter viacepapi.ViaCepAdapter, weatherAdapter weatherapi.WeatherAdapter) *FindCepAndWeatherUseCase {
	return &FindCepAndWeatherUseCase{cepAdapter: cepAdapter, weatherAdapter: weatherAdapter}
}

func (cw FindCepAndWeatherUseCase) Execute(input CepWeatherInput) (*CepWeatherOutput, *commons.Problem) {
	viaCep, problem := cw.cepAdapter.GetZipCode(input.Cep)
	if problem != nil {
		return nil, problem
	}
	weather, problem := cw.weatherAdapter.GetWeather(viaCep.Localidade)
	if problem != nil {
		return nil, problem
	}
	return &CepWeatherOutput{
		Location:       viaCep.Localidade,
		Condition:      weather.Current.Condition.Text,
		TempCelsius:    weather.Current.TempC,
		TempFahrenheit: weather.Current.TempF,
		TempKelvin:     parseCelsiusToKelvin(weather.Current.TempC),
	}, nil
}

func parseCelsiusToKelvin(tempCelsius float64) float64 {
	return tempCelsius + 273
}
