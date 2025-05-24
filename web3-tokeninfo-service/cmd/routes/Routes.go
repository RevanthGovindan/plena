package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"web3-tokeninfo/internal/database"

	"github.com/gorilla/mux"
)

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

		fmt.Println(keyData)

		ctx := context.WithValue(r.Context(), "userID", "user123")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRoutes(r *mux.Router) *mux.Router {
	r.Use(requestAuthenticator)
	r.HandleFunc("/tokens/{tokenId}", fetchTokenInfo).Methods(http.MethodGet)
	return r
}
