package application

import (
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/commons/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/infrastructure/viacepapi"
)

type FindCepInput struct {
	Cep string `json:"cep"`
}

type FindCepOutput struct {
	ViaCepDto *viacepapi.ViaCepDto
}

type FindCepUseCase struct {
	cepAdapter *viacepapi.ViaCepAdapter
}

func NewFindCepUseCase(cepAdapter *viacepapi.ViaCepAdapter) *FindCepUseCase {
	return &FindCepUseCase{cepAdapter: cepAdapter}
}

func (f *FindCepUseCase) Execute(input FindCepInput) (*FindCepOutput, *commons.Problem) {
	viaCep, problem := f.cepAdapter.GetZipCode(input.Cep)
	if problem != nil {
		return nil, problem
	}
	return &FindCepOutput{
		ViaCepDto: viaCep,
	}, nil
}
