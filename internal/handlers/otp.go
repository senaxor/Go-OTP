package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"OTP/internal/models"
	"OTP/internal/utils"
)

var otpStore = make(map[string]string) // In-memory OTP store (use Redis in production)

type OTPRequest struct {
	Phone string `json:"phone"`
}

type OTPVerifyRequest struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

// POST /otp/request
func RequestOTP(w http.ResponseWriter, r *http.Request) {
	var req OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	otp := utils.GenerateOTP()
	otpStore[req.Phone] = otp

	// Log the OTP to simulate "sending"
	log.Printf("Generated OTP for %s: %s\n", req.Phone, otp)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP generated (check logs)", // sms in production
	})
}

// POST /otp/verify
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req OTPVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	expectedOTP, exists := otpStore[req.Phone]
	if !exists || expectedOTP != req.OTP {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	log.Printf(req.Phone)
	user, err := models.FindOrCreateUser(req.Phone)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	delete(otpStore, req.Phone) // Clean up used OTP

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
