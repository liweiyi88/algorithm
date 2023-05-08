package ratelimiter

import (
	"math"
	"sync"
	"time"
)

type RefillRate struct {
	Interval time.Duration
	Number   int
}

func NewRefillRate(interval time.Duration, number int) RefillRate {
	return RefillRate{
		Interval: interval,
		Number:   number,
	}
}

type TokenBucket struct {
	rate          RefillRate
	mu            sync.Mutex
	currentTokens int
	maxToken      int
	lastRefillAt  time.Time
}

func NewTokenBucket(maxToken int, rate RefillRate) *TokenBucket {
	return &TokenBucket{
		rate:          rate,
		currentTokens: maxToken,
		maxToken:      maxToken,
		lastRefillAt:  time.Now(),
	}
}

func (tb *TokenBucket) Accept(tokens int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.Refill()

	if tb.currentTokens >= tokens {
		tb.currentTokens -= tokens
		return true
	}

	return false
}

func (tb *TokenBucket) Refill() {
	elapse := time.Since(tb.lastRefillAt)
	tokensToBeAdded := float64(elapse.Nanoseconds() * int64(tb.rate.Number) / tb.rate.Interval.Nanoseconds())
	tb.currentTokens = int(math.Min(float64(tb.currentTokens)+tokensToBeAdded, float64(tb.maxToken)))
	tb.lastRefillAt = time.Now()
}
