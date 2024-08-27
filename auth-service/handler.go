package main

import (
	"encoding/json"
	"net/http"
)

// Login godoc
// @Summary Login user
// @Description Login user with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} string "ok"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Login successful")
}

// Register godoc
// @Summary Register new user
// @Description Register new user with username, email, and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 201 {string} string "created"
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully")
}
