package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	defaultPort       = "8008"
	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second

	HTTPMethodOverrideHeader  = "X-HTTP-Method-Override"
	HTTPMethodOverrideFormKey = "_method"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

type Articles struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Post     string    `json:"post"`
	Complete bool      `json:"complete"`
	CreateAt time.Time `json:"created_at"`
}

var Article []Articles

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("--> %s %s", req.Method, req.URL.Path)
		// fmt.Println(req.URL.Path)
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		statusCode := lrw.statusCode
		log.Printf("<-- %d %s", statusCode, http.StatusText(statusCode))
	})
}

func ArticleGETALLHanddler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Article)
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusAccepted)
	// Articles := Article{
	// 	Article{Title: "ASD", CreateAt: time.Now().Unix()},
	// 	Article{Title: "ASD", CreateAt: time.Now().Unix()},
	// }
	// json.E
	fmt.Fprintf(w, "Categorys is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var articles Articles
	_ = json.NewDecoder(r.Body).Decode(&articles)
	articles.ID = strconv.Itoa(rand.Intn(10000))
	Article = append(Article, articles)
	json.NewEncoder(w).Encode(&articles)
}

func main() {

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/articles", ArticleGETALLHanddler).Methods("GET")
	r.HandleFunc("/articles", CreatePostHandler).Methods("POST")
	r.HandleFunc("/articles/{category}/{id:[0-9]}", ArticleHandler).Methods("GET")
	srv := &http.Server{
		Handler:           r,
		Addr:              "127.0.0.1:" + defaultPort,
		WriteTimeout:      writeTimeout,
		ReadTimeout:       readTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
	}
	// log.Fatal(srv.ListenAndServe())
	// http.ListenAndServe(":8008", handlers.CORS()(r))
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("ERR ListenAndServe: %s")
	}
}
