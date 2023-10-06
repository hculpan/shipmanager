package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hculpan/shipmanager/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Serve static files from "assets" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/register-ship", handlers.RegisterShipHandler)
	http.HandleFunc("/load-ship", handlers.LoadShipHandler)

	fmt.Printf("Starting server on port %s\n", port)

	http.ListenAndServe(":"+port, LogRequest(http.DefaultServeMux))
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime))
	})
}
