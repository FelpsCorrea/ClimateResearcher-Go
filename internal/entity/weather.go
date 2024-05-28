package entity

type WeatherRepositoryInterface interface {
	GetWeather(zipcode string) (Weather, error)
}

type Weather struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func (*Weather) CalculateKelvinByCelsius(celsius float64) float64 {
	return celsius + 273
}
