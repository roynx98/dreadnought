package strategies

import (
	"adeptus-limitarius/cases"
	"adeptus-limitarius/entities"
)

type StrategiesRegistry struct {
	mediator cases.LimiterMediator
}

func (registry StrategiesRegistry) Register() {
	registry.mediator.RegisterCreationHandler("bucket", func() entities.Limiter {
		return BucketLimiter{}
	})
}

func ProvideStrategiesRegistry(mediator cases.LimiterMediator) StrategiesRegistry {
	return StrategiesRegistry{mediator: mediator}
}
