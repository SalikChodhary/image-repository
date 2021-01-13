package main

import (
    "log"
    "net/http"

		"github.com/gorilla/mux"
		"github.com/SalikChodhary/shopify-challenge/routes"
		"github.com/SalikChodhary/shopify-challenge/services"
		"github.com/SalikChodhary/shopify-challenge/middleware"
)

func main() {
		router := mux.NewRouter()
		services.InitServices()
		router.Handle("/api/v1/add", middleware.IsJWTAuthorized(routes.AddImage)).Methods("POST")
		router.Handle("/api/v1/search", middleware.IsJWTAuthorized(routes.SearchImage)).Methods("GET")
		router.HandleFunc("/login", routes.Login).Methods("POST")
		router.HandleFunc("/signup", routes.Signup).Methods("POST")
	
		log.Fatal(http.ListenAndServe(":8000", router))
}