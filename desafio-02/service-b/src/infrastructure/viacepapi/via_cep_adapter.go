package viacepapi

import (
	"encoding/json"
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/commons/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/infrastructure/env"
	"log"
	"net/http"
	"net/url"
	"path"
)

type ViaCepDto struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v *ViaCepDto) isValidLocation() bool {
	return len(v.Localidade) > 0
}

type ViaCepAdapter struct {
	config env.EnvConfig
}

func NewViaCepAdapter(config env.EnvConfig) *ViaCepAdapter {
	return &ViaCepAdapter{config: config}
}

func (v *ViaCepAdapter) GetZipCode(cep string) (*ViaCepDto, *commons.Problem) {
	baseUrl, err := url.Parse(v.config.ViaCepApiUri)
	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	baseUrl.Path = path.Join("ws", cep, "/json/")
	uri := baseUrl.String()

	resp, err := http.Get(uri)

	if resp.StatusCode == 400 {
		return nil, commons.NewError422("invalid zipcode", "invalid zipcode")
	}

	if resp.StatusCode == 404 {
		return nil, commons.NewError404("can not find zipcode", "can not find zipcode")
	}

	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	defer resp.Body.Close()

	viaCep := &ViaCepDto{}
	err = json.NewDecoder(resp.Body).Decode(viaCep)
	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	if !viaCep.isValidLocation() {
		return nil, commons.NewError404("can not find zipcode", "can not find zipcode")
	}

	return viaCep, nil
}
