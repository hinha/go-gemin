package controllers

import (
	"net/http"

	"go-gemin/controllers/users"
	"go-gemin/controllers/blog"
	"github.com/gorilla/mux"
	"github.com/pytimer/mux-logrus"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/articles", blog.ArticleGETALLHanddler).Methods("GET")
	r.HandleFunc("/articles", blog.CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/v1/auth/signup", users.RegistersHandler).Methods("POST")
	// r.HandleFunc("/api/v1/auth/signup", signup.RegistersHandler).Methods("POST")
	// r.HandleFunc("/api/v1/users/profile", signup.RegistersHandler).Methods("POST")


	// Doc
	// r.HandleFunc("/articles/{category}/{id:[0-9]}", ArticleHandler).Methods("GET")
	r.Use(muxlogrus.NewLogger().Middleware)
	return r
}
