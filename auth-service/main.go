package main

import (
	"log"
	"monorepo/auth-service/docs"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Auth Service API
// @version 1.0
// @description This is a sample auth service API.
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	// Swagger docs route
	docs.SwaggerInfo.BasePath = "/"
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Auth routes
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/register", Register).Methods("POST")

	log.Println("Auth service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
