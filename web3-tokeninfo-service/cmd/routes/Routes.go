package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"web3-tokeninfo/internal/database"

	"github.com/gorilla/mux"
)

type contextKey string

func requestAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
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
		fmt.Println(keyData)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRoutes(r *mux.Router) *mux.Router {
	r.Use(requestAuthenticator)
	r.HandleFunc("/tokens/{tokenId}", fetchTokenInfo).Methods(http.MethodGet)
	return r
}
