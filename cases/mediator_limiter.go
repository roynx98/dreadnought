package cases

type LimiterMediator struct {
	creationHandlers *map[string]func() Limiter
}

func (m *LimiterMediator) RegisterCreationHandler(id string, handler func() Limiter) {
	(*m.creationHandlers)[id] = handler
}

func (m *LimiterMediator) Create(name string) Limiter {
	return (*m.creationHandlers)[name]()
}

func ProvideLimiterMediator() LimiterMediator {
	return LimiterMediator{creationHandlers: &map[string]func() Limiter{}}
}
