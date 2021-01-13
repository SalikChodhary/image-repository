package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"sync"
  "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"github.com/SalikChodhary/shopify-challenge/models"
)

var client *mongo.Client
var imgCollection *mongo.Collection
var userCollection *mongo.Collection

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
		imgCollection = client.Database("ImageRepo").Collection("images")
		userCollection = client.Database("ImageRepo").Collection("users")
	})
	return err
}

func InsertNewImageToMongo(img models.Image) {
	imgCollection.InsertOne(context.TODO(), img)
}

func InsertNewUserToMongo(user *models.User) {
	userCollection.InsertOne(context.TODO(), user)
}

func IsExistingUsername(username string) (*models.User, bool) {
	filter := bson.D{{"username", username}}
	var res models.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&res)

	return &res, err == nil
}

func SearchImageByID(id string) {

}

func SearchImageByTags(tags string) {
	
}