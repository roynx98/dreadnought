package cases

import (
	"adeptus-limitarius/entities"
)

type LimiterMediator struct {
	creationHandlers *map[string]func() entities.Limiter
}

func (m *LimiterMediator) RegisterCreationHandler(id string, handler func() entities.Limiter) {
	(*m.creationHandlers)[id] = handler
}

func (m *LimiterMediator) Create(name string) entities.Limiter {
	return (*m.creationHandlers)[name]()
}

func ProvideLimiterMediator() LimiterMediator {
	return LimiterMediator{creationHandlers: &map[string]func() entities.Limiter{}}
}
