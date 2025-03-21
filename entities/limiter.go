package entities

type Limiter interface {
	Start(onDelete func())
	ShouldLimit(rule LimitRule) bool
}

type LimitRule struct {
	IP       string
	Strategy string
}
