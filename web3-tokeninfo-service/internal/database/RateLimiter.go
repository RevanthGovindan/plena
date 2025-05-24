package database

import (
	"sync"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	mu       sync.RWMutex
	limiters map[string]*rate.Limiter
}

var LimiterStore = newRateLimiter()

func newRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (rl *RateLimiter) GetLimiter(key string, rateLimit int) *rate.Limiter {
	rl.mu.RLock()
	defer rl.mu.RUnlock()

	if _, exists := rl.limiters[key]; !exists {
		rl.limiters[key] = rate.NewLimiter(rate.Limit(float64(rateLimit)/60), rateLimit)
	}
	return rl.limiters[key]
}

func (rl *RateLimiter) UpdateRateLimiter(key string, rateLimit int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.limiters[key] = rate.NewLimiter(rate.Limit(float64(rateLimit)/60), rateLimit)
}
