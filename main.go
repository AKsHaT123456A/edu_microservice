package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	api "edumarshal.com/api/api"
	db "edumarshal.com/api/db"
	handler "edumarshal.com/api/handlers"
	"edumarshal.com/api/middleware"
)

func main() {
	router := mux.NewRouter()
	// config, err := config.LoadConfig()
	// if err != nil {
	// log.Fatalf("Error loading config: %v", err)
	// }
	// jwtKey := config.JWT_SIGN_KEY
	// signingKey := []byte(jwtKey)
	// middleware.InitJWTMiddleware(signingKey)

	router.Handle("/", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("Welcome to EduMarshal API")); err != nil {
			log.Println("Failed to write response:", err)
		}
	}))).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/auth/register", api.CreateUserPost).Methods(http.MethodPost)
	router.Handle("/api/v1/add_attendance", middleware.JWTMiddleware(http.HandlerFunc(api.AttendanceHandler))).Methods(http.MethodPost)
	router.Handle("/api/v1/get_attendance", middleware.JWTMiddleware(http.HandlerFunc(api.GetAttendanceHandler))).Methods(http.MethodGet)
	router.Handle("/api/v1/subject", middleware.JWTMiddleware(http.HandlerFunc(api.SubjectHandler))).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/login", api.UserLoginPost).Methods(http.MethodPost)
	router.HandleFunc("/refresh-token", handler.RefreshTokenHandler).Methods(http.MethodPost)

	// router.HandleFunc("/api/v1/get_attendance", api.GetAttendanceHandler).Methods(http.MethodPost)

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

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	port := os.Getenv("HOSTPORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Server starting on port %s", port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
