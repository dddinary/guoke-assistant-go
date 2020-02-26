package middleware

import "golang.org/x/time/rate"

var userLimiter		map[int]*rate.Limiter
var noLoginLimiter	*rate.Limiter

func init() {
	userLimiter = map[int]*rate.Limiter{}
	noLoginLimiter = rate.NewLimiter(20, 20)
}

func isAllowed(uid int) bool {
	limiter, ok := userLimiter[uid]
	if !ok {
		limiter = rate.NewLimiter(5, 3)
		userLimiter[uid] = limiter
	}
	return limiter.Allow()
}