package ratelimiter

import (
	"testing"
	"time"
)

func TestAllow(t *testing.T) {
	tokenBucket := NewTokenBucket(3, NewRefillRate(5*time.Second, 3))

	if !tokenBucket.Accept(2) {
		t.Error("expected accept, but got deny")
	}

	if !tokenBucket.Accept(1) {
		t.Error("expected accept, but got deny")
	}

	if tokenBucket.Accept(1) {
		t.Error("expected too many requests, but got accepted")
	}

	time.Sleep(5 * time.Second)

	if !tokenBucket.Accept(3) {
		t.Error("expected accept, but got deny")
	}
}
