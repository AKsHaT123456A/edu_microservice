package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"edumarshal.com/api/auth"
)

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	fmt.Println(authHeader)
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("missing or malformed authorization header")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		tokenString, err := extractToken(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, err.Error())
			return
		}

		// Verify the access token
		if !auth.VerifyToken(tokenString) {
			log.Println("Invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		// Proceed with the request
		next.ServeHTTP(w, r)
	})
}
