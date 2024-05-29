package entity

type Weather struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func (w *Weather) CalculateKelvinByCelsius(celsius float64) {
	w.Kelvin = celsius + 273
}
