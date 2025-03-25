package cases

import (
	"dreadnought/entities"
	"sync"
)

type LimiterInteractor struct {
	mediator       LimiterMediator
	activeLimiters *map[string]entities.Limiter
	mutex          *sync.RWMutex
}

func (interactor LimiterInteractor) ShouldLimit(rule entities.LimitRule) bool {
	interactor.mutex.RLock()
	limiter := (*interactor.activeLimiters)[rule.IP]
	interactor.mutex.RUnlock()

	if limiter == nil {
		limiter = interactor.mediator.Create(rule.Strategy)
		interactor.mutex.Lock()
		(*interactor.activeLimiters)[rule.IP] = limiter
		interactor.mutex.Unlock()

		limiter.Start(func() {
			interactor.mutex.Lock()
			delete(*interactor.activeLimiters, rule.IP)
			interactor.mutex.Unlock()
		})
	}

	return limiter.ShouldLimit(rule)
}

func ProvideLimiterInteractor(mediator LimiterMediator) LimiterInteractor {
	limiters := make(map[string]entities.Limiter)
	return LimiterInteractor{activeLimiters: &limiters, mediator: mediator, mutex: &sync.RWMutex{}}
}
