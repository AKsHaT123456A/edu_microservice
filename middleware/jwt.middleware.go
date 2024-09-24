package middleware

import (
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/context"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware

// InitJWTMiddleware initializes the JWT middleware with the signing key.
func InitJWTMiddleware(signingKey []byte) {
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

// JWTMiddleware is the middleware for validating JWT tokens.
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if JWT token is valid

		err := jwtMiddleware.CheckJWT(w, r)
		log.Println("JWT Middleware", err)
		if err != nil {
			// Write an unauthorized response only if token validation fails
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Set decoded JWT token in context
		context.Set(r, "decoded", context.Get(r, "user"))

		// Call the next handler if the JWT check passed
		next.ServeHTTP(w, r)
	})
}
