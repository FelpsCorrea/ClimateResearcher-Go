package main

import (
	"github.com/FelpsCorrea/ClimateResearcher-Go/configs"
	"github.com/FelpsCorrea/ClimateResearcher-Go/internal/infra/web/webserver"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	server := webserver.NewWebServer(configs.WebServerPort)

}
