package dto

type GetWeatherInputDTO struct {
	Zipcode string `json:"zipcode"`
}

type GetWeatherOutputDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
	Country    string  `json:"country"`
	City       string  `json:"city"`
}

type GetWeatherErrorDTO struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
