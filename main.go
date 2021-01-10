package main

import (
    "log"
    "net/http"

		"github.com/gorilla/mux"
		"github.com/SalikChodhary/shopify-challenge/routes"
		"github.com/SalikChodhary/shopify-challenge/services"
)


func main() {
		router := mux.NewRouter()
		services.InitServices()
		router.HandleFunc("/api/v1/add", routes.AddImage).Methods("POST")
		log.Fatal(http.ListenAndServe(":8000", router))
}