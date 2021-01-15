package services

import (
	"context"
	"os"
	"sync"

	"github.com/SalikChodhary/shopify-challenge/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var imgCollection *mongo.Collection
var userCollection *mongo.Collection

func InitiateMongoClient() (error) {
	var once sync.Once
	var err error = nil
	once.Do(func() {
		var localErr error
		uri := "mongodb+srv://"+os.Getenv("MONGO_USER")+":"+os.Getenv("MONGO_PASSWORD")+"@cluster0.xplun.mongodb.net/"+os.Getenv("DB_NAME")+"?retryWrites=true&w=majority"
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

func InsertNewImageToMongo(img models.Image) ([]byte, error) {
	res, err := imgCollection.InsertOne(context.TODO(), img)

	if err != nil {
		return nil, err
	}

	oid := res.InsertedID.(primitive.ObjectID)
	idAsBytes := oid[:]

	return idAsBytes, nil
}

func InsertNewUserToMongo(user *models.User) ([]byte, error) {
	res, err := userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	oid := res.InsertedID.(primitive.ObjectID)
	idAsBytes := oid[:]

	return idAsBytes, nil
}

func IsExistingUsername(username string) (*models.User, bool) {
	filter := bson.D{{"username", username}}
	var res models.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&res)

	return &res, err == nil
}

func SearchImageByID(id string, user string) (*models.Image, error) {
	filter := bson.D{
		{"key", id},
		{"$or", bson.A{
			bson.M{"owner": user},
			bson.M{"private": false},
		}},
	}

	var res models.Image
	err := imgCollection.FindOne(context.TODO(),filter).Decode(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func SearchImageByTags(tags []string, user string) ([]models.Image, error) {
	filter := bson.D{
		{"tags", bson.M{
			"$all": tags,
		}},
		{"$or", bson.A{
			bson.M{"owner": user},
			bson.M{"private": false},
		}},
	}

	var res []models.Image
	c, err := imgCollection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	if err = c.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	
	return res, err
}

func SearchImageByUser(userQuery string, user string) ([]models.Image, error) {
	// var filter primitive.D
	filter := bson.M{"owner": userQuery}
	if user != userQuery {
		filter["private"] = false
	}
	//fmt.Printf("%v", filter)
	var res []models.Image
	c, err := imgCollection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	if err = c.All(context.TODO(), &res); err != nil {
		return nil, err
	}

	if res == nil {
		res = make([]models.Image, 0)
	}
	
	return res, err
}

func GetAllImages(user string) ([]models.Image, error) {
	filter := bson.M{
		"$or": bson.A{
			bson.M{"private": false},
			bson.M{"owner": user, "private": true},
		},
	}

	var res []models.Image
	c, err := imgCollection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	if err = c.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	
	return res, err
}