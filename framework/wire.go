//go:build wireinject
// +build wireinject

package framework

import (
	"dreadnought/adapters"
	"dreadnought/cases"
	"dreadnought/framework/config"
	"dreadnought/framework/limiters"
	"dreadnought/framework/networking"

	"github.com/google/wire"
)

func InitializeApp() App {
	wire.Build(
		ProvideApp,
		config.ProvideConfigManager,
		limiters.ProvideStrategiesRegistry,
		adapters.ProvideLimiterController,
		networking.ProvideLimiterServer,
		cases.ProvideLimiterInteractor,
		cases.ProvideLimiterMediator,
		wire.Bind(new(networking.LimiterServer), new(networking.HttpLimiterServer)),
	)
	return App{}
}
