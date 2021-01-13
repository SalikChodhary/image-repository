package routes

import (
	"fmt"
	"net/http"
	"log"
	"github.com/SalikChodhary/shopify-challenge/services"
)

func AddImage(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := services.ParseImageDataFromRequest(r)
	image := services.InitImageStruct(r, file, fileHeader)

  if err != nil{
    log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
    return
	}

	err = services.UploadFileToS3(file, fileHeader, image.Key)

	if err != nil {
		fmt.Fprintf(w, "Could not upload file")
		return
	}

	err = services.AppendAutoTagsToImage(&image)
	services.InsertNewImageToMongo(image)

	fmt.Fprintf(w, "Image uploaded successfully.")
}