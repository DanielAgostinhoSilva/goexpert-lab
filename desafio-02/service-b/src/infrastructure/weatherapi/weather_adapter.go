package weatherapi

import (
	"encoding/json"
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/commons/infrastructure/api"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/infrastructure/env"
	"log"
	"net/http"
	"net/url"
	"path"
)

type WeatherAdapter struct {
	config env.EnvConfig
}

func NewWeatherAdapter(config env.EnvConfig) *WeatherAdapter {
	return &WeatherAdapter{config: config}
}

func (w *WeatherAdapter) GetWeather(location string) (*WeatherData, *commons.Problem) {

	baseUrl, err := url.Parse(w.config.WeatherApiUri)
	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	baseUrl.Path = path.Join("v1", "current.json")

	urlParams := url.Values{}
	urlParams.Add("key", w.config.WeatherApiKey)
	urlParams.Add("q", location)
	urlParams.Add("aqi", "no")

	baseUrl.RawQuery = urlParams.Encode()
	uri := baseUrl.String()

	resp, err := http.Get(uri)
	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewError500("Failed to get weather data", "Failed to get weather data")
	}

	weather := &WeatherData{}
	err = json.NewDecoder(resp.Body).Decode(weather)

	if err != nil {
		log.Printf("Error making GET request: %v\n", err)
		return nil, commons.NewInternalServerError(err.Error())
	}

	return weather, nil

}
