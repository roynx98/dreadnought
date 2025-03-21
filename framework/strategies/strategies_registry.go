package strategies

import (
	"dreadnought/cases"
	"dreadnought/entities"
	"sync/atomic"
)

type StrategiesRegistry struct {
	mediator cases.LimiterMediator
}

func (registry StrategiesRegistry) Register() {
	registry.mediator.RegisterCreationHandler("bucket", func() entities.Limiter {
		const maxTokens = 2
		var currentTokens atomic.Int32
		currentTokens.Store(maxTokens)
		return &BucketLimiter{currentTokens: &currentTokens, maxTokens: maxTokens, refillRate: 2000, overflowsToClose: 3}
	})
}

func ProvideStrategiesRegistry(mediator cases.LimiterMediator) StrategiesRegistry {
	return StrategiesRegistry{mediator: mediator}
}
