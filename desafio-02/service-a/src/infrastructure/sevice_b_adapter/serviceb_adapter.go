package sevice_b_adapter

import (
	"context"
	"encoding/json"
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-a/src/commons/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-a/src/infrastructure/env"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"net/http"
	"net/url"
	"path"
)

type ServiceBDTO struct {
	City           string  `json:"city"`
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type ServiceBAdapter struct {
	config env.EnvConfig
}

func NewServiceBAdapter(config env.EnvConfig) *ServiceBAdapter {
	return &ServiceBAdapter{config: config}
}

func (s *ServiceBAdapter) GetWeather(ctx context.Context, cep string) (*ServiceBDTO, *commons.Problem) {
	baseUrl, err := url.Parse(s.config.ServiceBUri)
	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	baseUrl.Path = path.Join("weather", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl.String(), nil)
	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, commons.NewError422("invalid zipcode", "invalid zipcode")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, commons.NewError404("can not find zipcode", "can not find zipcode")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, commons.NewInternalServerError("unexpected status code")
	}

	serviceBDTO := &ServiceBDTO{}
	err = json.NewDecoder(resp.Body).Decode(serviceBDTO)
	if err != nil {
		log.Printf("Error decoding response body: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	return serviceBDTO, nil

}
