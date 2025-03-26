package framework

import (
	"dreadnought/framework/config"
	"dreadnought/framework/limiters"
	"dreadnought/framework/networking"
)

type App struct {
	server           networking.LimiterServer
	limitersRegistry limiters.LimitersRegistry
	configManager    config.ConfigManager
}

func (app App) Start() {
	app.configManager.Load()

	app.limitersRegistry.Register()

	app.server.Start()
}

func ProvideApp(
	limiterServer networking.LimiterServer,
	strategiesRegistry limiters.LimitersRegistry,
	configManager config.ConfigManager) App {
	return App{
		server:           limiterServer,
		limitersRegistry: strategiesRegistry,
		configManager:    configManager,
	}
}
