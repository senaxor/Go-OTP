// @title OTP Verification API
// @version 1.0
// @description API for sending and verifying OTP codes
// @host localhost:8080
// @BasePath /
// @schemes http

package main

import (
	"log"
	"net/http"

	"OTP/internal/database"
	"OTP/internal/handlers"
	"OTP/internal/middlewares"
	_ "OTP/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger"
)

func main() {
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load it")
	}

    database.InitMongo()
    database.InitRedis()
    r := mux.NewRouter()

    // OTP Endpoints
    r.HandleFunc("/otp/request", handlers.RequestOTP).Methods("POST")
    r.HandleFunc("/otp/verify", handlers.VerifyOTP).Methods("POST")

    // Protect /users with admin auth middleware
	r.Handle("/users", middlewares.AdminAuth(http.HandlerFunc(handlers.GetAllUsers))).Methods("GET")


	// Add this route to serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    log.Println("Server running at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
