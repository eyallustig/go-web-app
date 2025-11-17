package main

import (
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	start := time.Now()
	defer func() {
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}()
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	start := time.Now()
	defer func() {
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}()
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	start := time.Now()
	defer func() {
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}()
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	start := time.Now()
	defer func() {
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}()
	http.ServeFile(w, r, "static/contact.html")
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func main() {
	// Set up file server for static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes with logging middleware
	http.HandleFunc("/home", loggingMiddleware(homePage))
	http.HandleFunc("/courses", loggingMiddleware(coursePage))
	http.HandleFunc("/about", loggingMiddleware(aboutPage))
	http.HandleFunc("/contact", loggingMiddleware(contactPage))

	// Log server start
	serverAddr := "0.0.0.0:8080"
	log.Printf("Server starting on http://%s", serverAddr)
	log.Printf("Available endpoints:")
	log.Printf("  - http://%s/home", serverAddr)
	log.Printf("  - http://%s/courses", serverAddr)
	log.Printf("  - http://%s/about", serverAddr)
	log.Printf("  - http://%s/contact", serverAddr)

	// Start server
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}