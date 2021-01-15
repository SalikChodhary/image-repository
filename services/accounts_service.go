package services

import (
	"net/http"
	"github.com/SalikChodhary/shopify-challenge/models"

	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

func isValidUsername(username string) bool {
	l := len(username)

	return !(l < 3 || l > 25) 
}

func isValidPassword(pwd string) bool {
	l := len(pwd)

	return !(l < 8)
}

func getHashedPassword(pwd string) (string, error) {
	a := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(a, bcrypt.MinCost)

	return string(hash), err
}

func AuthenticatePassword(plainPwd string, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))

	return err == nil
}

func InitUser(r *http.Request) (*models.User, error) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if !isValidUsername(u.Username) {
		return nil, errors.New("Invalid username, must be between 3 and 25 characters.")
	}

	if !isValidPassword(u.Password) {
		return nil, errors.New("Invalid password, must be longer than 8 characters.")
	}

	hp, err := getHashedPassword(u.Password)

	if err != nil {
		return nil, err
	}

	u.HashedPassword = hp

	return &u, nil
}