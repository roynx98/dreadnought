package cases

import (
	"adeptus-limitarius/entities"
)

type LimitRule struct {
	IP       string
	Strategy string
}

type LimiterInteractor struct {
	mediator       LimiterMediator
	activeLimiters *map[string]entities.Limiter
}

func (interactor LimiterInteractor) ShouldLimit(rule LimitRule) bool {
	limiter := interactor.mediator.Create(rule.Strategy)
	limiter.Start()

	return false
}

func ProvideLimiterInteractor(mediator LimiterMediator) LimiterInteractor {
	limiters := make(map[string]entities.Limiter)
	return LimiterInteractor{activeLimiters: &limiters, mediator: mediator}
}
