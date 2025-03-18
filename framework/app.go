package framework

import (
	"log"
	"net/url"
)

type App struct {
	server LimiterServer
}

func (app App) Start() {
	targetURL := "http://localhost:3000/"

	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}

	app.server.Start(target)
}

func ProvideApp(limiterServer LimiterServer) App {
	return App{server: limiterServer}
}
