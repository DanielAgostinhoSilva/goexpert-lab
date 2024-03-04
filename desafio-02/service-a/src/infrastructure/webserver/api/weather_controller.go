package api

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-01/service-a/src/application"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type WeatherController struct {
	tracer             trace.Tracer
	findWeatherUseCase *application.FindWeatherUseCase
}

func NewWeatherController(tracer trace.Tracer, findWeatherUseCase *application.FindWeatherUseCase) *WeatherController {
	return &WeatherController{
		tracer:             tracer,
		findWeatherUseCase: findWeatherUseCase,
	}
}

func (wc *WeatherController) FindWeather(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	ctx, span := wc.tracer.Start(ctx, "service-a-request")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")
	var input application.FindWeatherInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		wc.writeError(w, span, err, http.StatusBadRequest)
		return
	}

	weatherOutput, err := wc.findWeatherUseCase.Execute(input)
	if err != nil {
		wc.writeError(w, span, err, http.StatusUnprocessableEntity)
		return
	}
	json.NewEncoder(w).Encode(weatherOutput)

}

func (wc *WeatherController) writeError(w http.ResponseWriter, span trace.Span, err error, status int) {
	span.RecordError(err) // Record error in span
	span.SetAttributes(attribute.Key("error").Bool(true))
	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(status))
	http.Error(w, err.Error(), status)
	return
}

func (wc *WeatherController) Router(router chi.Router) {
	router.Post("/", wc.FindWeather)
}
func (wc *WeatherController) Path() string {
	return "/weather"
}
