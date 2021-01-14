package models
type Image struct {
	// ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key string `json:"key,omitempty" bson:"key,omitempty"`
	Tags []string `json:"tags,omitempty" bson:"tags,omitempty"`
	Owner string `json:"owner,omitempty" bson:"owner,omitempty"`//change to user later
	Private bool `json:"private" bson:"private"`

}