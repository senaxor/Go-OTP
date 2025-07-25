package handlers

import (
	"encoding/json"
	"net/http"

	"OTP/internal/models"
	"OTP/internal/utils"
)

// GET /users/me
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch userss", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
