//go:build wireinject
// +build wireinject

package framework

import (
	"adeptus-limitarius/adapters"

	"github.com/google/wire"
)

func InitializeApp() App {
	wire.Build(ProvideApp, adapters.ProvideLimiterController, ProvideLimiterServer, wire.Bind(new(LimiterServer), new(HttpLimiterServer)))
	return App{}
}
