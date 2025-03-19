package cases

type LimitRule struct {
	IP string
}

type LimiterInteractor struct {
}

func (iteractor LimiterInteractor) ShouldLimit(rule LimitRule) bool {
	// TODO: Handle tickers
	return false
}

func ProvideLimiterInteractor() LimiterInteractor {
	return LimiterInteractor{}
}
