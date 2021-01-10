package services

import (
	"sync"
  "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"github.com/SalikChodhary/shopify-challenge/models"
)

var client *mongo.Client
var collection *mongo.Collection

func InitiateMongoClient() (error) {
	var once sync.Once
	var err error = nil
	once.Do(func() {
		var localErr error
		uri := "mongodb+srv://admin:dbtest@cluster0.xplun.mongodb.net/ShopifyChallenge?retryWrites=true&w=majority"
		opts := options.Client()
		opts.ApplyURI(uri)
		if client, localErr = mongo.Connect(context.Background(), opts); localErr != nil {
				err = localErr
				return
		}
		collection = client.Database("ImageRepo").Collection("images")
	})
	return err
}

func InsertNewImageToMongo(img models.Image) {
	collection.InsertOne(context.TODO(), img)
}

func SearchImageByID(id string) {

}

func SearchImageByTags(tags string) {
	
}