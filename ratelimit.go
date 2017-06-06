package ratelimit

// RateLimiter 限速器
type RateLimiter struct {
}

// New 创建RateLimiter实例
func New() *RateLimiter {
	return &RateLimiter{}
}

// Limit 限制
func (rl *RateLimiter) Limit() bool {
	return false
}
