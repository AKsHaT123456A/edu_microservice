package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"edumarshal.com/api/auth"
	"edumarshal.com/api/models"
	"edumarshal.com/api/utils"
)

func CreateUserPost(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Print(err)

		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(user.PasswordHash)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = hashedPassword
	user.Username = user.Username + "194"
	if err := utils.InsertUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("User created successfully")); err != nil {
		log.Println("Failed to write response:", err)
	}
}

func UserLoginPost(w http.ResponseWriter, r *http.Request) {
	var credentials models.User

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		fmt.Print(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Println("Credentials", credentials)
	user, err := utils.GetUserByUsername(credentials.Username, credentials.DOB)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !auth.ValidatePassword(credentials.PasswordHash, user.PasswordHash) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.Username)
	log.Println("Token generated", token)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	response := struct {
		Token                string `json:"token"`
		Username             string `json:"username"`
		StudentName          string `json:"studentName,omitempty"`
		DOB                  string `json:"dob,omitempty"`
		Email                string `json:"email,omitempty"`
		StudentNumber        string `json:"studentNumber,omitempty"`
		UniversityRollNumber int64  `json:"universityRollNumber,omitempty"`
	}{
		Token:                token,
		Username:             user.Username,
		StudentName:          user.StudentName,
		DOB:                  user.DOB,
		Email:                user.Email,
		StudentNumber:        user.StudentNumber,
		UniversityRollNumber: user.UniversityRollNumber,
	}
	fmt.Print(user)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Failed to encode response:", err)
	}
}
