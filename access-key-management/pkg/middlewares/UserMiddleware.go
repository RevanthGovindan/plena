package middlewares

import "net/http"

func UserAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate jwt
		next.ServeHTTP(w, r)
	})
}
