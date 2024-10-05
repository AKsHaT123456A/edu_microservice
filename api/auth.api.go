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

type Role struct {
	Role string `json:"role"`
}

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
	var UserLoginRequest models.UserLoginRequest
	// Decode the user credentials from the request body
	if err := json.NewDecoder(r.Body).Decode(&UserLoginRequest); err != nil {
		fmt.Println("Error decoding credentials:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve user by username and date of birth (or other criteria)
	user, err := utils.GetUserByUsername(UserLoginRequest.Username, UserLoginRequest.DOB)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		http.Error(w, "Invalid username or password ", http.StatusUnauthorized)
		return
	}

	// Validate the password hash
	if !auth.ValidatePassword(UserLoginRequest.PasswordHash, user.PasswordHash) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	// Generate JWT token using username and role
	accessToken, refreshToken := auth.GenerateTokens(user.Username)
	if accessToken == "" || refreshToken == "" {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	// Build the response object
	response := struct {
		Token                string `json:"token"`
		RefreshToken         string `json:"refreshToken"`
		Username             string `json:"username"`
		StudentName          string `json:"studentName,omitempty"`
		DOB                  string `json:"dob,omitempty"`
		Email                string `json:"email,omitempty"`
		StudentNumber        string `json:"studentNumber,omitempty"`
		UniversityRollNumber int64  `json:"universityRollNumber,omitempty"`
		ProfileImage         string `json:"profileImage,omitempty"`
		Branch               string `json:"branch,omitempty"`
		Hostel               string `json:"hostel,omitempty"`
		Semester             string `json:"semester,omitempty"`
		AdmissionDate        string `json:"admissionDate,omitempty"`
		AdmissionMode        string `json:"admissionMode,omitempty"`
		Categories           string `json:"categories,omitempty"`
		JeeRank              string `json:"jeeRank,omitempty"`
		IsLateral            bool   `json:"isLateral,omitempty"`
		Section              string `json:"section,omitempty"`
		Course               string `json:"course,omitempty"`
		IsRemembered         bool   `json:"isRemembered,omitempty"`
		IsAdmin              bool   `json:"isAdmin,omitempty"`
		UserID               uint   `json:"userId,omitempty"`
	}{
		Token:                accessToken,
		RefreshToken:         refreshToken,
		Username:             user.Username,
		StudentName:          user.StudentName,
		DOB:                  user.DOB,
		Email:                user.Email,
		StudentNumber:        user.StudentNumber,
		UniversityRollNumber: user.UniversityRollNumber,
		ProfileImage:         user.ProfileImage,
		Branch:               user.Branch,
		Hostel:               user.Hostel,
		Semester:             user.Semester,
		AdmissionDate:        user.AdmissionDate,
		AdmissionMode:        user.AdmissionMode,
		Categories:           user.Categories,
		JeeRank:              user.JeeRank,
		IsLateral:            user.IsLateral,
		Section:              user.Section,
		Course:               user.Course,
		IsRemembered:         user.IsRemembered,
		IsAdmin:              user.IsAdmin,
		UserID:               user.ID,
	}

	// Set the response header and return the JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println("Failed to encode response:", err)
	}
}
