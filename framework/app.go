package framework

import (
	"dreadnought/framework/networking"
	"dreadnought/framework/strategies"
	"log"
	"net/url"
)

type App struct {
	server             networking.LimiterServer
	strategiesRegistry strategies.StrategiesRegistry
}

func (app App) Start() {
	targetURL := "http://localhost:3000/"

	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}

	app.strategiesRegistry.Register()

	app.server.Start(target)
}

func ProvideApp(limiterServer networking.LimiterServer, strategiesRegistry strategies.StrategiesRegistry) App {
	return App{server: limiterServer, strategiesRegistry: strategiesRegistry}
}
