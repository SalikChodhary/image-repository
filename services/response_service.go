package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SalikChodhary/shopify-challenge/models"
)

type ResponseType int

const (
	Error ResponseType = iota
	Success
)

func SendResponse(rt ResponseType, message interface{}, code int, w http.ResponseWriter) {
	var res models.Response

	res.Success = rt == Success

	if rt == Success { 
		res.Data = message
	} else if rt == Error {
		res.Error = message
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	
	j, _ := json.Marshal(res)

	fmt.Fprintf(w, "%v", string(j))
}