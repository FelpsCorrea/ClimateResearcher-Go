package usecase

import (
	"errors"

	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/entity"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/infra/external"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/usecase/dto"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/validation"
)

type GetWeatherUseCase struct {
	WeatherRepository entity.WeatherRepositoryInterface
	WeatherAPIClient  external.WeatherClient
}

func NewGetWeatherUseCase(
	WeatherRepository entity.WeatherRepositoryInterface,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		WeatherRepository: WeatherRepository,
	}
}

func (u *GetWeatherUseCase) Execute(input dto.GetWeatherInputDTO) (dto.GetWeatherOutputDTO, error) {

	if !validation.IsValidZipCode(input.Zipcode) {
		return dto.GetWeatherOutputDTO{}, errors.New("Invalid zipcode")
	}

	// Do the external API call
	weatherResponse, err := u.WeatherAPIClient.GetWeather(input.Zipcode)

	if err != nil {
		return dto.GetWeatherOutputDTO{}, err
	}

}
