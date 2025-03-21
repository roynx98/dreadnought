//go:build wireinject
// +build wireinject

package framework

import (
	"dreadnought/adapters"
	"dreadnought/cases"
	"dreadnought/framework/networking"
	"dreadnought/framework/strategies"

	"github.com/google/wire"
)

func InitializeApp() App {
	wire.Build(
		ProvideApp,
		strategies.ProvideStrategiesRegistry,
		adapters.ProvideLimiterController,
		networking.ProvideLimiterServer,
		cases.ProvideLimiterInteractor,
		cases.ProvideLimiterMediator,
		wire.Bind(new(networking.LimiterServer), new(networking.HttpLimiterServer)),
	)
	return App{}
}
