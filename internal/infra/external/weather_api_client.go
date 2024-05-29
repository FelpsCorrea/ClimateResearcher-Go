package external

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"

	"github.com/valyala/fastjson"
)

type WeatherClient interface {
	GetWeather(city string) (*WeatherClientResponseDTO, WeatherClientResponseErrorDTO)
}

type WeatherAPIClient struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

func NewWeatherAPIClient(baseURL, apiKey string) *WeatherAPIClient {
	return &WeatherAPIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

func (c *WeatherAPIClient) GetWeather(city string) (*WeatherClientResponseDTO, WeatherClientResponseErrorDTO) {
	url := c.BaseURL + "/current.json?q=" + city + "&key=" + c.APIKey

	print(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, WeatherClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, WeatherClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}
	defer resp.Body.Close()

	var p fastjson.Parser

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, WeatherClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	data, err := p.ParseBytes(body)
	if err != nil {
		return nil, WeatherClientResponseErrorDTO{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	if resp.StatusCode == 400 && data.Get("error").Get("code").GetInt() == 1006 {
		return nil, WeatherClientResponseErrorDTO{Message: "invalid zipcode", StatusCode: http.StatusNotFound}
	}

	if resp.StatusCode != 200 {
		return nil, WeatherClientResponseErrorDTO{Message: "error on request", StatusCode: http.StatusInternalServerError}
	}

	locationJson := data.GetObject("location")
	currentJson := data.GetObject("current")

	weatherResponse := WeatherClientResponseDTO{
		Celsius:    currentJson.Get("temp_c").GetFloat64(),
		Fahrenheit: currentJson.Get("temp_f").GetFloat64(),
		City:       locationJson.Get("name").String(),
		Country:    locationJson.Get("country").String(),
	}

	return &weatherResponse, WeatherClientResponseErrorDTO{
		Message:    "",
		StatusCode: http.StatusOK,
	}
}

type WeatherClientResponseErrorDTO struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type WeatherClientResponseDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
}
