package application

import (
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/commons/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/infrastructure/weatherapi"
)

type WeatherInput struct {
	Location string
}

type WeatherOutput struct {
	City           string  `json:"city"`
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type FindWeatherUseCase struct {
	weatherAdapter *weatherapi.WeatherAdapter
}

func NewFindWeatherUseCase(weatherAdapter *weatherapi.WeatherAdapter) *FindWeatherUseCase {
	return &FindWeatherUseCase{weatherAdapter: weatherAdapter}
}

func (cw FindWeatherUseCase) Execute(input WeatherInput) (*WeatherOutput, *commons.Problem) {
	weather, problem := cw.weatherAdapter.GetWeather(input.Location)
	if problem != nil {
		return nil, problem
	}
	return &WeatherOutput{
		City:           input.Location,
		TempCelsius:    weather.Current.TempC,
		TempFahrenheit: weather.Current.TempF,
		TempKelvin:     parseCelsiusToKelvin(weather.Current.TempC),
	}, nil
}

func parseCelsiusToKelvin(tempCelsius float64) float64 {
	return tempCelsius + 273
}
