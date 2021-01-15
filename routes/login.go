package routes

import (
	"net/http"

	"github.com/SalikChodhary/shopify-challenge/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	u, err := services.InitUser(r)

	if err != nil {
		services.SendResponse(services.Error, err.Error(), http.StatusBadRequest, w)
		return
	}

	existingUser, exists := services.IsExistingUsername(u.Username)

	if !exists {
		services.SendResponse(services.Error, err.Error(), http.StatusNotFound, w)
		return
	}

	isValidLogin := services.AuthenticatePassword(u.Password, existingUser.HashedPassword)

	if !isValidLogin {
		services.SendResponse(services.Error, "wrong password, please try again!", http.StatusUnauthorized, w)
		return
	}

	s, err := services.MakeJWT(u.Username)
	
	if err != nil {
		services.SendResponse(services.Error, err.Error(), http.StatusInternalServerError, w)
		return
	}

	services.SendResponse(services.Success, s, http.StatusOK, w)
}