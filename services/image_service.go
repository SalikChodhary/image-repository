package services

import (
	"image"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/SalikChodhary/shopify-challenge/models"
	"github.com/pkg/errors"

	// "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"strings"

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

func toLower(a []string) ([]string) {
	for _, s := range a { 
		s = strings.ToLower(s)
	}
	return a
}

func InitImageStruct(r *http.Request, file multipart.File, header *multipart.FileHeader) models.Image {
	m := make(map[string]bool)
	m["true"] = true
	m["false"] = false
	bytes, _ := uuid.NewRandom()
	uuid := bytes.String()
	isPrivateStr := r.FormValue("private")
	tags := toLower(strings.Split(r.FormValue("tags"), ","))

	username := GetUserFromHeader(r)

	var img models.Image
	img.Key = uuid
	img.Owner = username
	img.Tags = tags
	img.URI = "https://" + os.Getenv("S3_BUCKET_NAME") + ".s3." + os.Getenv("REGION_NAME") + ".amazonaws.com/" + img.Key

	if isPrivateStr == "" {
		img.Private = false
		return img
	}
	
	isPrivate, valid := m[isPrivateStr]

	if !valid {
		img.Private = false
	} else {
		img.Private = isPrivate
	}

	return img
}

func AppendAutoTagsToImage(img *models.Image) error {
	s3URI := img.URI
	tags, err := getImageTags(s3URI)

	if err != nil {
		return errors.New(err.Error())
	}

	for _, tag := range tags {
		img.Tags = append(img.Tags, tag)
	}
	return nil
}