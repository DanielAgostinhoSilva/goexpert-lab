package api

import (
	"encoding/json"
	"errors"
	"github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/application"
	commons "github.com/DanielAgostinhoSilva/goexpert-lab/desafio-02/service-b/src/commons/infrastructure/api"
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
	findCepUseCase     *application.FindCepUseCase
	findWeatherUseCase *application.FindWeatherUseCase
}

func NewWeatherController(tracer trace.Tracer, findCepUseCase *application.FindCepUseCase, findWeatherUseCase *application.FindWeatherUseCase) *WeatherController {
	return &WeatherController{tracer: tracer, findCepUseCase: findCepUseCase, findWeatherUseCase: findWeatherUseCase}
}

func (wc *WeatherController) FindWeather(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, mainSpan := wc.tracer.Start(ctx, "INITIAL_SPAN")
	defer mainSpan.End()

	ctx, viaCepSpan := wc.tracer.Start(ctx, "external call viacep")
	viaCepOutput, problem := wc.findCepUseCase.Execute(application.FindCepInput{Cep: chi.URLParam(r, "cep")})
	if problem != nil {
		wc.writeError(w, viaCepSpan, problem)
		return
	}
	viaCepSpan.End()

	ctx, weatherSpan := wc.tracer.Start(ctx, "external call weather")
	weatherOutput, problem := wc.findWeatherUseCase.Execute(application.WeatherInput{
		Location: viaCepOutput.ViaCepDto.Localidade,
	})
	if problem != nil {
		wc.writeError(w, weatherSpan, problem)
		return
	}
	weatherSpan.End()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(weatherOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wc *WeatherController) writeError(w http.ResponseWriter, span trace.Span, problem *commons.Problem) {
	span.RecordError(errors.New(problem.Detail)) // Record error in span
	span.SetAttributes(attribute.Key("error").Bool(true))
	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(problem.Status))
	span.End()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(problem.Status)
	json.NewEncoder(w).Encode(problem)
}

func (wc *WeatherController) Router(router chi.Router) {
	router.Get("/", wc.FindWeather)
}
func (wc *WeatherController) Path() string {
	return "/weather/{cep}"
}
