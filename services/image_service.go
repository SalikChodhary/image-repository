package services

import (
	"github.com/pkg/errors"
	"log"
	"net/http"
	"mime/multipart"
	"image"
	"github.com/SalikChodhary/shopify-challenge/models"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"github.com/google/uuid"
)

const (
	maxSize = 2048000
	formDataImageKey = "img"
)

func ParseImageDataFromRequest(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	r.ParseMultipartForm(int64(maxSize))



	return r.FormFile(formDataImageKey)
}

func IsValidImage(f multipart.File) bool {
	_, _, err = image.Decode(f)
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}

func InitImageStruct(r *http.Request, file multipart.File, header *multipart.FileHeader) models.Image {
	bytes, _ := uuid.NewRandom()
	uuid := bytes.String()
	var img models.Image
	img.Key = uuid
	img.Owner = "admin"
	img.Private = false
	img.Tags = []string{"no", "tags"}

	return img
}

func AppendAutoTagsToImage(img *models.Image) error {
	s3URI := "https://salik-test-bucket.s3.us-east-2.amazonaws.com/" + img.Key
	tags, err := getImageTags(s3URI)

	if err != nil {
		return errors.New(err.Error())
	}

	for _, tag := range tags {
		img.Tags = append(img.Tags, tag)
	}
	return nil
}