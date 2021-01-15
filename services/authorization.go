package services

import (
	"errors"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))


func MakeJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tknString, err := token.SignedString(secretKey)

	if (err != nil) {
		return "", err
	}
	return tknString, nil
}

func hasAuthHeader(r *http.Request) (string, bool) {
	ts, in := r.Header["Authorization"] 
	return ts[0], in
}

func getTokenFromString(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Internal error.")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, true
	}

	return token, false
}

func isValidToken(tokenString string) bool {
	token, hasErr := getTokenFromString(tokenString)

	if hasErr {
		return false
	}

	_, validClaims := token.Claims.(jwt.MapClaims)
	validToken := token.Valid
	
	return validClaims && validToken
}

func IsAuthorizedRequest(r *http.Request) bool {
	val, hasToken := hasAuthHeader(r)

	return hasToken && isValidToken(val)
}

func GetUserFromHeader(r *http.Request) string {
	ts, _ := hasAuthHeader(r)
	token, _ := getTokenFromString(ts)
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["sub"].(string)
}

