package cases

import (
	"adeptus-limitarius/entities"
)

type LimiterInteractor struct {
	mediator       LimiterMediator
	activeLimiters *map[string]entities.Limiter
}

func (interactor LimiterInteractor) ShouldLimit(rule entities.LimitRule) bool {
	limiter := (*interactor.activeLimiters)[rule.IP]

	if limiter == nil {
		limiter = interactor.mediator.Create(rule.Strategy)
		(*interactor.activeLimiters)[rule.IP] = limiter

		limiter.Start(func() {
			delete(*interactor.activeLimiters, rule.IP)
		})
	}

	return limiter.ShouldLimit(rule)
}

func ProvideLimiterInteractor(mediator LimiterMediator) LimiterInteractor {
	limiters := make(map[string]entities.Limiter)
	return LimiterInteractor{activeLimiters: &limiters, mediator: mediator}
}
