package application

import (
	"context"
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-a/src/commons/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-a/src/domain"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-a/src/infrastructure/sevice_b_adapter"
)

type FindWeatherInput struct {
	Cep string `json:"cep"`
}

type FindWeatherOutput struct {
	City           string  `json:"city"`
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type FindWeatherUseCase struct {
	serviceB *sevice_b_adapter.ServiceBAdapter
}

func NewFindWeatherUseCase(serviceB *sevice_b_adapter.ServiceBAdapter) *FindWeatherUseCase {
	return &FindWeatherUseCase{serviceB: serviceB}
}

func (f *FindWeatherUseCase) Execute(ctx context.Context, input FindWeatherInput) (*FindWeatherOutput, *commons.Problem) {
	cep, err := domain.NewCep(input.Cep)
	if err != nil {
		return nil, commons.NewError422("invalid zipcode", "invalid zipcode")
	}
	weatherDto, problem := f.serviceB.GetWeather(ctx, cep.Numero())
	if problem != nil {
		return nil, problem
	}

	return &FindWeatherOutput{
		weatherDto.City,
		weatherDto.TempCelsius,
		weatherDto.TempFahrenheit,
		weatherDto.TempKelvin,
	}, nil
}
