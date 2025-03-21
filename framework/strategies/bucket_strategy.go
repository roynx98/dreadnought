package strategies

import (
	"adeptus-limitarius/entities"
	"sync/atomic"
	"time"
)

type BucketLimiter struct {
	currentTokens atomic.Int32
	maxTokens     int
}

func (limiter *BucketLimiter) Start() {
	ticker := time.NewTicker(5000 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				if limiter.currentTokens.Load() < int32(limiter.maxTokens) {
					limiter.currentTokens.Add(1)
				}
			}
		}
	}()
}

func (limiter *BucketLimiter) ShouldLimit(rule entities.LimitRule) bool {
	if limiter.currentTokens.Load() <= 0 {
		return true
	}

	limiter.currentTokens.Add(-1)
	return false
}
