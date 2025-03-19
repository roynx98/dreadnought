//go:build wireinject
// +build wireinject

package framework

import (
	"adeptus-limitarius/adapters"
	"adeptus-limitarius/cases"
	"adeptus-limitarius/framework/networking"

	"github.com/google/wire"
)

func InitializeApp() App {
	wire.Build(
		ProvideApp,
		adapters.ProvideLimiterController,
		networking.ProvideLimiterServer,
		cases.ProvideLimiterInteractor,
		wire.Bind(new(networking.LimiterServer), new(networking.HttpLimiterServer)),
	)
	return App{}
}
