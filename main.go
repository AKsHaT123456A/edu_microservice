package main

import (
	"log"
	"net/http"

	api "edumarshal.com/api/api"
	config "edumarshal.com/api/config"
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

	// Handler for /api/v1/login
	server.HandleFunc("/api/v1/auth/login", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %v", r)
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.UserLoginPost(w, r)
	})

	// Database connection
	dbConn, err := db.DB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer func() {
		sqlDB, err := dbConn.DB()
		if err != nil {
			log.Printf("Error getting database connection: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Load configuration
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	port := conf.HostPort
	if port == "" {
		port = "4000"
	}

	// Log server start message before starting the server
	log.Printf("Server starting on port %s", port)

	// Start the server
	err = http.ListenAndServe(":"+port, server)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
