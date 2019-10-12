package controllers

import (
	"net/http"

	"github.com/go-gemin/martinusdawan/controllers/signup"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/auth/signup", signup.RegistersHandler).Methods("POST")
	// r.HandleFunc("/api/v1/auth/signup", signup.RegistersHandler).Methods("POST")
	// r.HandleFunc("/api/v1/users/profile", signup.RegistersHandler).Methods("POST")

	return r
}
