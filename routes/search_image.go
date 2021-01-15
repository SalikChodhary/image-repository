package routes

import (
	"encoding/json"
	"net/http"

	"github.com/SalikChodhary/shopify-challenge/models"
	"github.com/SalikChodhary/shopify-challenge/services"
)

var searchTypes map[string]struct{} = map[string]struct{}{
	"by_id": struct{}{}, 
	"by_tags": struct{}{},
	"by_user": struct{}{},
	"all": struct{}{},
}

func sendResponse(res interface{}, w http.ResponseWriter) {
	services.SendResponse(services.Success, res, http.StatusOK, w)
}

func handleByIDRequest(id string, user string, w http.ResponseWriter)  {
	res, err := services.SearchImageByID(id, user)

	if err != nil {
		services.SendResponse(services.Error, "no image found", http.StatusNoContent, w)
		return 
	}

	services.SendResponse(services.Success, res, http.StatusOK, w)
}

func handleByTagsRequest(tags []string, user string, w http.ResponseWriter) {
	res, err := services.SearchImageByTags(tags, user)

	if err != nil {
		services.SendResponse(services.Error, "no image found", http.StatusNoContent, w)
		return 
	}

	services.SendResponse(services.Success, res, http.StatusOK, w)
}

func handleByUserRequest(requestingUser string, requestedUser string, w http.ResponseWriter) {
	res, err := services.SearchImageByUser(requestingUser, requestedUser)

	if err != nil {
		services.SendResponse(services.Error, "no image found", http.StatusNoContent, w)
		return 
	}

	services.SendResponse(services.Success, res, http.StatusOK, w)
}

func handleGetAllRequest(user string, w http.ResponseWriter) {
	res, err := services.GetAllImages(user)

	if err != nil {
		services.SendResponse(services.Error, "no image found", http.StatusNoContent, w)
		return 
	}

	services.SendResponse(services.Success, res, http.StatusOK, w)
}

func SearchImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var searchRequest models.SearchRequest

	err := decoder.Decode(&searchRequest)

	if err != nil {
		services.SendResponse(services.Error, err.Error(), http.StatusBadRequest, w)
		return
	}

	_, isValidSearchType := searchTypes[searchRequest.Type]

	if !isValidSearchType {
		services.SendResponse(services.Error, "invalid search type", http.StatusBadRequest, w)
		return
	}

	searchType := searchRequest.Type
	user := services.GetUserFromHeader(r)

	if searchType == "by_id" {
		handleByIDRequest(searchRequest.QueryID, user, w)
	} else if searchType == "by_tags" {
		handleByTagsRequest(searchRequest.QueryTags, user, w)
	} else if searchType == "by_user" {
		handleByUserRequest(searchRequest.QueryUser, user, w)
	} else { 
		handleGetAllRequest(user, w)
	}
	
}