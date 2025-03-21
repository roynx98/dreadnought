package entities

type Limiter interface {
	Start()
	ShouldLimit(rule LimitRule) bool
}

type LimitRule struct {
	IP       string
	Strategy string
}
