package auth

import (
	"log"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Define signing key globally
var signingKey []byte

// InitializeJWTKey initializes the signing key from the config.
func InitializeJWTKey(jwtKey string) {
	signingKey = []byte(jwtKey)
	log.Printf("JWT signing key initialized")
}

// GenerateTokens generates an access token and a refresh token.
func GenerateTokens(userID string) (string, string) {
	// Access token valid for 15 minutes
	accessToken, err := createToken(userID, time.Now().Add(time.Hour*24))
	if err != nil {
		log.Printf("Error creating access token: %v", err)
		return "", ""
	}

	// Refresh token valid for 7 days
	refreshToken, err := createToken(userID, time.Now().Add(time.Hour*24*7))
	if err != nil {
		log.Printf("Error creating refresh token: %v", err)
		return "", ""
	}

	return accessToken, refreshToken
}

// createToken creates a JWT token with an expiration time.
func createToken(userID string, expirationTime time.Time) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// VerifyToken verifies the JWT access token.
func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Unexpected signing method")
			return nil, nil
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		log.Printf("Invalid token: %v", err)
		return false
	}
	return true
}

// VerifyRefreshToken verifies the JWT refresh token.
func VerifyRefreshToken(tokenString string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Unexpected signing method")
			return nil, nil
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		log.Printf("Invalid refresh token: %v", err)
		return ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Printf("Invalid refresh token claims")
		return ""
	}

	userID := claims["sub"].(string)
	return userID
}

// ValidatePassword checks if the provided password matches the hash
func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword hashes the provided password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
