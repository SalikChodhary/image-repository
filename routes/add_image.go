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
	
	// isValid := services.IsValidImage(file)
	
  // if !isValid {
  //   log.Printf("could not decode body into an image")
  //   w.Header().Add("Access-Control-Allow-Origin", "*")
  //   w.WriteHeader(http.StatusBadRequest)
  //   w.Write([]byte("could not decode body image"))
  //   return
	// }

	// err = services.InitS3Instance()

	// if err != nil {
	// 	fmt.Fprintf(w, "Could not upload file")
	// }

	err = services.UploadFileToS3(file, fileHeader, image.Key)

	if err != nil {
		fmt.Fprintf(w, "Could not upload file")
	}
	services.InsertNewImageToMongo(image)


	fmt.Fprintf(w, "Image uploaded successfully.")
}