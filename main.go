package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/gorilla/mux"
	"go-gemin/controllers"
)

const (
	defaultPort       = "8008"
	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second

	HTTPMethodOverrideHeader  = "X-HTTP-Method-Override"
	HTTPMethodOverrideFormKey = "_method"
	//methodsOk := []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

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


func main() {

	// Router
	handler := controllers.New()

	srv := &http.Server{
		Handler:           handler,
		Addr:              "127.0.0.1:" + defaultPort,
		WriteTimeout:      writeTimeout,
		ReadTimeout:       readTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
	}
	// log.Fatal(srv.ListenAndServe())
	// http.ListenAndServe(":8008", handlers.CORS()(r))
	log.Printf("Listen on address %s", defaultPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("ERR ListenAndServe: %s")
	}
}
