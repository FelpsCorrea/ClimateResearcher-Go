package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnInvalidZipCode(t *testing.T) {

	weatherClient := NewWeatherAPIClient("https://api.weatherapi.com/v1", "360ddfd38d0d4cd3b72102808240403")

	_, errResponse := weatherClient.GetWeather("New York")
	assert.NotNil(t, errResponse)
	assert.Equal(t, "invalid zipcode", errResponse.Message)
}

func TestGivenAValidZipCode(t *testing.T) {

	weatherClient := NewWeatherAPIClient("https://api.weatherapi.com/v1", "360ddfd38d0d4cd3b72102808240403")

	response, errResponse := weatherClient.GetWeather("Sao Paulo")
	assert.Equal(t, 200, errResponse.StatusCode)
	assert.Contains(t, response.City, "Sao Paulo")
}
