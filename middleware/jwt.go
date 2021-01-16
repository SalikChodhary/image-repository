package middleware

import (
	"net/http"
	"github.com/SalikChodhary/shopify-challenge/services"
)

func IsJWTAuthorized(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 
		if !services.IsAuthorizedRequest(r) {
			services.SendResponse(services.Error, "Unauthorized access.", http.StatusUnauthorized, w)
			return
		}
		handler(w, r)
	})
}
