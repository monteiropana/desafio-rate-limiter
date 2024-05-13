package middleware

import (
	"net/http"

	"github.com/rate-limiter-desafio/internal/repository"
	"github.com/rate-limiter-desafio/pkg/rateLimit"
)

func RateLimitMiddleware(dbClient repository.Database) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !rateLimit.RateLimit(r, "ACCESS_TOKEN", dbClient) {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
