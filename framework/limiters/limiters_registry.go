package limiters

import (
	"dreadnought/cases"
	"dreadnought/entities"
	"dreadnought/framework/config"
	"sync/atomic"
)

type LimitersRegistry struct {
	mediator      cases.LimiterMediator
	configManager config.ConfigManager
}

func (registry LimitersRegistry) Register() {
	registry.mediator.RegisterCreationHandler("bucket", func() entities.Limiter {
		const maxTokens = 2
		var currentTokens atomic.Int32
		currentTokens.Store(maxTokens)
		return &BucketLimiter{currentTokens: &currentTokens, maxTokens: maxTokens, refillRate: 2000, overflowsToClose: 3}
	})
}

func ProvideStrategiesRegistry(mediator cases.LimiterMediator, configManager config.ConfigManager) LimitersRegistry {
	return LimitersRegistry{mediator: mediator, configManager: configManager}
}
