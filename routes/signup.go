package routes

import (
	"net/http"

	"github.com/SalikChodhary/shopify-challenge/services"
)


func Signup(w http.ResponseWriter, r *http.Request) { 
	u, err := services.InitUser(r)

	if err != nil {
		services.SendResponse(services.Error, err.Error(), http.StatusBadRequest, w)
		return
	}
	_, exists := services.IsExistingUsername(u.Username)
	if exists {
		services.SendResponse(services.Error, "User already exists with that username.", http.StatusConflict, w)
		return 
	}

	_, err = services.InsertNewUserToMongo(u)

	if err != nil {
		services.SendResponse(services.Error, "could not upload user", http.StatusInternalServerError, w)
	}

	services.SendResponse(services.Success, string(u.Username), http.StatusOK, w)
}