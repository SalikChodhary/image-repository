package models

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )
type Image struct {
	// ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key string `json:"key,omitempty" bson:"key,omitempty"`
	Tags string `json:"data.tags,omitempty" bson:"tags,omitempty"`
	// BucketLink string `json:"bucket_link,omitempty" bson:"bucket_link,omitempty"`
	Owner string `json:"data.owner,omitempty" bson:"owner,omitempty"`//change to user later
	Private bool `json:"data.private,omitempty" bson:"private,omitempty"`

}