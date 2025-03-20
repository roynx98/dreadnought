package strategies

import "fmt"

type BucketLimiter struct {
}

func (limiter BucketLimiter) Start() {
	fmt.Println("BucketLimiter started")
}
