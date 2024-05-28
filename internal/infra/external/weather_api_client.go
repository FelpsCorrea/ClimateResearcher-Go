package external

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/valyala/fastjson"
)

type WeatherClient interface {
	GetWeather(zipCode string) (*WeatherClientResponseDTO, error)
}

type WeatherAPIClient struct {
	BaseURL string
	Client  *http.Client
}

func NewWeatherAPIClient(baseURL string) *WeatherAPIClient {
	return &WeatherAPIClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *WeatherAPIClient) GetWeather(zipCode string) (*WeatherClientResponseDTO, error) {
	url := c.BaseURL + "/weather?zip=" + zipCode
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser

	data, err := p.ParseBytes(body)
	if err != nil {
		return nil, err
	}

	locationJson := data.GetObject("location")
	currentJson := data.GetObject("current")

	weatherResponse := WeatherClientResponseDTO{
		Celsius:    currentJson.Get("temp_c").GetFloat64(),
		Fahrenheit: currentJson.Get("temp_f").GetFloat64(),
		City:       locationJson.Get("name").String(),
		Country:    locationJson.Get("country").String(),
	}

	return &weatherResponse, nil
}

type WeatherClientResponseDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
}
