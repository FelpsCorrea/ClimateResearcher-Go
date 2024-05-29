package usecase

import (
	"net/http"

	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/entity"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/infra/external"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/usecase/dto"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/validation"
)

type GetWeatherUseCase struct {
	WeatherAPIClient external.WeatherClient
	CepApiClient     external.CepAPIClient
}

func NewGetWeatherUseCase(
	WeatherAPIClient external.WeatherClient,
	CepApiClient external.CepAPIClient,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		WeatherAPIClient: WeatherAPIClient,
		CepApiClient:     CepApiClient,
	}
}

func (u *GetWeatherUseCase) Execute(input dto.GetWeatherInputDTO) (*dto.GetWeatherOutputDTO, *dto.GetWeatherErrorDTO) {

	if !validation.IsValidZipCode(input.Zipcode) {
		return &dto.GetWeatherOutputDTO{}, &dto.GetWeatherErrorDTO{
			Message:    "invalid zipcode",
			StatusCode: 422,
		}
	}

	// Do the external CEP API call
	cepResponse, errCep := u.CepApiClient.Get(input.Zipcode)
	if errCep.StatusCode != 200 {
		return &dto.GetWeatherOutputDTO{}, &dto.GetWeatherErrorDTO{
			Message:    errCep.Message,
			StatusCode: errCep.StatusCode,
		}
	}

	// Do the external weather API call based the found city name
	weatherResponse, err := u.WeatherAPIClient.GetWeather(cepResponse.Localidade)

	if err.StatusCode != http.StatusOK {
		return &dto.GetWeatherOutputDTO{}, &dto.GetWeatherErrorDTO{
			Message:    err.Message,
			StatusCode: 422,
		}
	}

	// Calculate the Kelvin temperature
	weather := entity.Weather{
		Celsius:    weatherResponse.Celsius,
		Fahrenheit: weatherResponse.Fahrenheit,
	}

	weather.CalculateKelvinByCelsius(weather.Celsius)

	return &dto.GetWeatherOutputDTO{
		Celsius:    weather.Celsius,
		Fahrenheit: weather.Fahrenheit,
		Kelvin:     weather.Kelvin,
		Country:    weatherResponse.Country,
		City:       weatherResponse.City,
	}, nil

}
