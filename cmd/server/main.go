package main

import (
	"log"
	"net/http"

	"OTP/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // OTP Endpoints
    r.HandleFunc("/otp/request", handlers.RequestOTP).Methods("POST")
    r.HandleFunc("/otp/verify", handlers.VerifyOTP).Methods("POST")

    log.Println("Server running at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
