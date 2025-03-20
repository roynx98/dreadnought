//go:build wireinject
// +build wireinject

package framework

import (
	"adeptus-limitarius/adapters"
	"adeptus-limitarius/cases"
	"adeptus-limitarius/framework/networking"
	"adeptus-limitarius/framework/strategies"

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
