package routes

import (
	"net/http"
	"encoding/json"
	"github.com/SalikChodhary/shopify-challenge/models"
	"github.com/SalikChodhary/shopify-challenge/services"
	"log"
)

var searchTypes map[string]struct{} = map[string]struct{}{
	"by_id": struct{}{}, 
	"by_tags": struct{}{},
}

func SearchImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var searchRequest models.SearchRequest

	err := decoder.Decode(&searchRequest)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	_, isValidSearchType := searchTypes[searchRequest.Type]

	if !isValidSearchType {
		log.Print("Invalid search type")
		w.WriteHeader(http.StatusBadRequest)
	}

	searchType := searchRequest.Type

	if searchType == "by_id" {
		services.SearchImageByID(searchRequest.QueryID)
	} else {
		services.SearchImageByTags(searchRequest.QueryTags)
	}
	
}