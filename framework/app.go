package framework

import (
	"adeptus-limitarius/framework/networking"
	"log"
	"net/url"
)

type App struct {
	server             networking.LimiterServer
	strategiesRegistry StrategiesRegistry
}

func (app App) Start() {
	targetURL := "http://localhost:3000/"

	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}

	app.strategiesRegistry.register()

	app.server.Start(target)
}

func ProvideApp(limiterServer networking.LimiterServer, strategiesRegistry StrategiesRegistry) App {
	return App{server: limiterServer, strategiesRegistry: strategiesRegistry}
}
