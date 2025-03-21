package strategies

import (
	"dreadnought/entities"
	"sync/atomic"
	"time"
)

type BucketLimiter struct {
	currentTokens    *atomic.Int32
	maxTokens        int32
	refillRate       int
	overflowCounter  int
	overflowsToClose int
}

func (limiter *BucketLimiter) Start(onDelete func()) {
	ticker := time.NewTicker(time.Duration(limiter.refillRate) * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				if limiter.currentTokens.Load() < limiter.maxTokens {
					limiter.currentTokens.Add(1)
					limiter.overflowCounter = 0
				} else {
					limiter.overflowCounter++
				}

				if limiter.overflowCounter >= limiter.overflowsToClose {
					onDelete()
					ticker.Stop()
					return
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
