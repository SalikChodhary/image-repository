package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SalikChodhary/shopify-challenge/middleware"
	"github.com/SalikChodhary/shopify-challenge/routes"
	"github.com/SalikChodhary/shopify-challenge/services"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting...")
	services.InitServices()
	router.Handle("/api/v1/add", middleware.IsJWTAuthorized(routes.AddImage)).Methods("POST")
	router.Handle("/api/v1/search", middleware.IsJWTAuthorized(routes.SearchImage)).Methods("GET")
	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/signup", routes.Signup).Methods("POST")

	fmt.Println("Server listening on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
