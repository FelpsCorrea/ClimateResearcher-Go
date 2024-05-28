package web

import (
	"encoding/json"
	"net/http"

	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/entity"
)

type WebWeatherHandler struct {
	WeatherRepository entity.WeatherRepositoryInterface
}

func NewWebOrderHandler(
	WeatherRepository entity.WeatherRepositoryInterface,
) *WebWeatherHandler {
	return &WebWeatherHandler{
		WeatherRepository: WeatherRepository,
	}
}

func (h *WebWeatherHandler) Get(w http.ResponseWriter, r *http.Request) {
	zipcode := r.PathValue("zipcode")
	getOrder := usecase.NewGetOrderUseCase(h.OrderRepository)
	output, err := getOrder.Execute(dto.GetOrderInputDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
