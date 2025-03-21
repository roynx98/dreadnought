package entities

type Limiter interface {
	Start(onDelete func())
	ShouldLimit(rule LimitRule) bool
}
