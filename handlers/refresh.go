package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"edumarshal.com/api/auth"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type TokenResponse struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// RefreshTokenHandler handles the refresh token request.
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
		return
	}
	log.Printf("Received refresh token: %s", request)
	// Verify the refresh token
	userID := auth.VerifyRefreshToken(request.RefreshToken)
	if userID == "" {
		// Invalid refresh token
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid refresh token"})
		return
	}

	// Generate new tokens
	accessToken, refreshToken := auth.GenerateTokens(userID)
	if accessToken == "" || refreshToken == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to generate tokens"})
		return
	}

	// Respond with new tokens
	json.NewEncoder(w).Encode(TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
