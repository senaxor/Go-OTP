package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"OTP/internal/database"
	"OTP/internal/models"
	"OTP/internal/utils"
	
	"github.com/redis/go-redis/v9"
)

var otpStore = make(map[string]string) // In-memory OTP store (use Redis in production)

type OTPRequest struct {
	Phone string `json:"phone"`
}

type OTPVerifyRequest struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

// RequestOTP godoc
// @Summary Request an OTP
// @Description Send an OTP to the user's phone
// @Tags OTP
// @Accept  json
// @Produce  json
// @Param request body OTPRequest true "Phone number"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /otp/request [post]
func RequestOTP(w http.ResponseWriter, r *http.Request) {
	var req OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	otp := utils.GenerateOTP()
	
	// Save OTP in Redis with 2-minute expiry
	err := database.RDB.Set(database.Ctx, req.Phone, otp, 2*time.Minute).Err()
	if err != nil {
		log.Printf("Redis SET error: %v\n", err)
		http.Error(w, "Failed to store OTP", http.StatusInternalServerError)
		return
	}

	// Log the OTP to simulate "sending"
	log.Printf("Generated OTP for %s: %s\n", req.Phone, otp)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP generated (check logs)", // sms in production
	})
}

// VerifyOTP godoc
// @Summary Verify OTP
// @Description Verify the OTP sent to the user's phone and return a JWT token if valid
// @Tags OTP
// @Accept  json
// @Produce  json
// @Param request body OTPVerifyRequest true "Phone number and OTP code"
// @Success 200 {object} map[string]string "Returns JWT token"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 401 {object} map[string]string "Invalid or expired OTP"
// @Failure 500 {object} map[string]string "Server or Redis/DB error"
// @Router /otp/verify [post]
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req OTPVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Fetch OTP from Redis
	expectedOTP, err := database.RDB.Get(database.Ctx, req.Phone).Result()
	if err == redis.Nil || expectedOTP != req.OTP {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Printf("Redis GET error: %v\n", err)
		http.Error(w, "Redis error", http.StatusInternalServerError)
		return
	}

	// Clean up used OTP
	_ = database.RDB.Del(database.Ctx, req.Phone).Err()
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

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
