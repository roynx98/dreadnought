package framework

import (
	"adeptus-limitarius/cases"
	"adeptus-limitarius/framework/strategies"
)

type StrategiesRegistry struct {
	mediator cases.LimiterMediator
}

func (registry StrategiesRegistry) register() {
	registry.mediator.RegisterCreationHandler("bucket", func() cases.Limiter {
		return strategies.BucketLimiter{}
	})
}

func ProvideStrategiesRegistry(mediator cases.LimiterMediator) StrategiesRegistry {
	return StrategiesRegistry{mediator: mediator}
}
