package api

import (
	"encoding/json"
	"errors"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-a/src/application"
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-a/src/commons/infrastructure/api"
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
	ctx, span := wc.tracer.Start(ctx, "INITIAL_SPAN")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")
	var input application.FindWeatherInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		wc.writeError(w, span, commons.NewError422("invalid zipcode", "invalid zipcode"))
		return
	}

	weatherOutput, problem := wc.findWeatherUseCase.Execute(ctx, input)
	if problem != nil {
		wc.writeError(w, span, problem)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(weatherOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (wc *WeatherController) writeError(w http.ResponseWriter, span trace.Span, problem *commons.Problem) {
	span.RecordError(errors.New(problem.Detail)) // Record error in span
	span.SetAttributes(attribute.Key("error").Bool(true))
	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(problem.Status))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(problem.Status)
	json.NewEncoder(w).Encode(problem)
}

func (wc *WeatherController) Router(router chi.Router) {
	router.Post("/", wc.FindWeather)
}
func (wc *WeatherController) Path() string {
	return "/weather"
}
