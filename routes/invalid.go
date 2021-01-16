package routes

import (
	"net/http"

	"github.com/SalikChodhary/shopify-challenge/services"
)

func Invalid(w http.ResponseWriter, r *http.Request) {
	services.SendResponse(
		services.Error, 
		"Invalid route, please visit https://github.com/SalikChodhary/shopify-challenge for documentation.",
		http.StatusNotImplemented, 
		w,
	)
}