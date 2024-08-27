package main

import (
	"encoding/json"
	"monorepo/user-service/models"
	"net/http"
)

// GetProfile godoc
// @Summary Get user profile
// @Description Get user profile by user ID
// @Tags user
// @Produce  json
// @Success 200 {object} models.User
// @Router /profile [get]
func GetProfile(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID:    "1",
		Name:  "John Doe",
		Email: "john@example.com",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
