package main

import (
	"log"
	"monorepo/user-service/docs"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title User Service API
// @version 1.0
// @description This is a sample user service API.
// @host localhost:8081
// @BasePath /

func main() {
	r := mux.NewRouter()

	// Swagger docs route
	docs.SwaggerInfo.BasePath = "/"
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// User routes
	r.HandleFunc("/profile", GetProfile).Methods("GET")

	log.Println("User service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
