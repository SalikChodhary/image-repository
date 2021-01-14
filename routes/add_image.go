package routes

import (
	"net/http"

	"github.com/SalikChodhary/shopify-challenge/services"
)

func AddImage(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := services.ParseImageDataFromRequest(r)
	image := services.InitImageStruct(r, file, fileHeader)

  if err != nil{
		services.SendResponse(services.Error, err.Error(), http.StatusBadRequest, w)
    return
	}

	err = services.UploadFileToS3(file, fileHeader, image.Key)

	if err != nil {
		services.SendResponse(services.Error, "could not upload file", http.StatusInternalServerError, w)
		return
	}

	err = services.AppendAutoTagsToImage(&image)
	res, err := services.InsertNewImageToMongo(image)

	if err != nil {
		services.SendResponse(services.Error, "could not upload file", http.StatusInternalServerError, w)
	}

	services.SendResponse(services.Success, string(res), http.StatusOK, w)
}