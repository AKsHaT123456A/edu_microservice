package router

import (
	"edumarshal.com/api/api"
	"github.com/gorilla/mux"
)

// InitRouter initializes the HTTP router with routes
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", api.CreateUserPost).Methods("POST")
	r.HandleFunc("/login", api.UserLoginPost).Methods("POST")

	return r
}
