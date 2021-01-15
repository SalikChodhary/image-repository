package middleware

import (
	"net/http"
	"fmt"
	"github.com/SalikChodhary/shopify-challenge/services"
)

func IsJWTAuthorized(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 
		if !services.IsAuthorizedRequest(r) {
			fmt.Fprintf(w, "Unauthorized access")
			return
		}
		handler(w, r)
	})
}