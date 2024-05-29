package web

import (
	"encoding/json"
	"net/http"

	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/infra/external"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/usecase"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/usecase/dto"
)

type WebWeatherHandler struct {
	WeatherAPIClient external.WeatherClient
	CepApiClient     external.CepAPIClient
}

func NewWebWeatherHandler(
	WeatherAPIClient external.WeatherClient,
	CepApiClient external.CepAPIClient,
) *WebWeatherHandler {
	return &WebWeatherHandler{
		WeatherAPIClient: WeatherAPIClient,
		CepApiClient:     CepApiClient,
	}
}

func (h *WebWeatherHandler) Get(w http.ResponseWriter, r *http.Request) {

	zipcode := r.PathValue("zipcode")

	getWeather := usecase.NewGetWeatherUseCase(h.WeatherAPIClient, h.CepApiClient)

	output, responseErr := getWeather.Execute(dto.GetWeatherInputDTO{Zipcode: zipcode})

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.StatusCode)
		return
	}

	err := json.NewEncoder(w).Encode(output)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
