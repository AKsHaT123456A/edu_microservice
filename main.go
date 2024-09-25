package main

import (
	"log"
	"net/http"

	api "edumarshal.com/api/api"
	db "edumarshal.com/api/db"
	"edumarshal.com/api/middleware"
)

func main() {
	server := http.NewServeMux()
	signingKey := []byte("aadfsfkdskmdkfmdkdfd")
	middleware.InitJWTMiddleware(signingKey)
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("Welcome to EduMarshal API")); err != nil {
				log.Println("Failed to write response:", err)
			}
		})).ServeHTTP(w, r)
	})

	// Handler for /api/v1/register
	server.HandleFunc("/api/v1/auth/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.CreateUserPost(w, r)
	})

	// Handler for /api/v1/auth/login
	server.HandleFunc("/api/v1/auth/login", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %v", r)
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.UserLoginPost(w, r)
	})

	// Database connection check
	_, err := db.DB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Start the server on port 3000
	err = http.ListenAndServe(":3000", server)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
