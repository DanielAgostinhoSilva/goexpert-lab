package application

import "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/domain"

type FindWeatherInput struct {
	Cep string `json:"cep"`
}

type FindWeatherOutput struct {
	Cep string
}

type FindWeatherUseCase struct {
}

func NewFindWeatherUseCase() *FindWeatherUseCase {
	return &FindWeatherUseCase{}
}

func (f *FindWeatherUseCase) Execute(input FindWeatherInput) (*FindWeatherOutput, error) {
	cep, err := domain.NewCep(input.Cep)
	if err != nil {
		return nil, err
	}
	return &FindWeatherOutput{
		Cep: cep.Numero(),
	}, nil
}
