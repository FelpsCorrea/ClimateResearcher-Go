package main

import (
	"fmt"

	"github.com/FelpsCorrea/ClimateResearcher-Go/configs"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/infra/external"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/infra/web"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/infra/web/webserver"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Weather API Client
	weatherAPIClient := external.NewWeatherAPIClient(configs.WeatherAPIBaseURL, configs.WeatherAPIKey)

	// ViaCEP API Client
	viaCEPAPIClient := external.NewViaCepAPIClient(configs.ViaCepAPIBaseURL)

	webserver := webserver.NewWebServer(configs.WebServerPort)

	webWeatherHandler := web.NewWebWeatherHandler(weatherAPIClient, viaCEPAPIClient)

	webserver.AddHandler("GET /weather/{zipcode}", webWeatherHandler.Get)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start()

}
