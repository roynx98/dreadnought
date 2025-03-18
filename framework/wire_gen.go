// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package framework

import (
	"adeptus-limitarius/adapters"
)

// Injectors from wire.go:

func InitializeApp() App {
	limiterController := adapters.ProvideLimiterController()
	httpLimiterServer := ProvideLimiterServer(limiterController)
	app := ProvideApp(httpLimiterServer)
	return app
}
