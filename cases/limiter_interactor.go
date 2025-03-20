package cases

type LimitRule struct {
	IP       string
	Strategy string
}

type LimiterInteractor struct {
	mediator       LimiterMediator
	activeLimiters *map[string]Limiter
}

func (interactor LimiterInteractor) ShouldLimit(rule LimitRule) bool {
	limiter := interactor.mediator.Create(rule.Strategy)
	limiter.Start()

	return false
}

func ProvideLimiterInteractor(mediaror LimiterMediator) LimiterInteractor {
	limiters := make(map[string]Limiter)
	return LimiterInteractor{activeLimiters: &limiters, mediator: mediaror}
}
