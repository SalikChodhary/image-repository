package services

import (
	"log"

)

func InitServices() {
	var err error
	
	err = InitS3Instance()

	if err != nil {
		log.Fatal(err.Error())
	}

	InitiateMongoClient()
	
	if err != nil {
		log.Fatal(err.Error())
	}
}