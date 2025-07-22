package main

import (
	"log"
	"net/http"

	"OTP/internal/database"
	"OTP/internal/handlers"
	"OTP/internal/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load it")
	}

    database.InitDB()
    r := mux.NewRouter()

    // OTP Endpoints
    r.HandleFunc("/otp/request", handlers.RequestOTP).Methods("POST")
    r.HandleFunc("/otp/verify", handlers.VerifyOTP).Methods("POST")

    // Protect /users with admin auth middleware
	r.Handle("/users", middlewares.AdminAuth(http.HandlerFunc(handlers.GetAllUsers))).Methods("GET")

    log.Println("Server running at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
