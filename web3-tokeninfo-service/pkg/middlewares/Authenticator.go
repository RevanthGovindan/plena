package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"
	"web3-tokeninfo/internal/database"
)

type contextKey string

func getAccessKey(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("missing or invalid Authorization header")
	}

	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func RequestAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getAccessKey(r)
		if err != nil {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		keyData, exists := database.GetDb().GetAccessData(token)
		if !exists {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !keyData.Enabled {
			http.Error(w, "Access disabled", http.StatusUnauthorized)
			return
		}

		if keyData.Expiry < time.Now().Unix() {
			http.Error(w, "access key expired", http.StatusUnauthorized)
			return
		}

		limiter := database.LimiterStore.GetLimiter(token, keyData.RateLimit)

		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		ctx := context.WithValue(r.Context(), contextKey("userId"), keyData.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
