package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TextPost   string             `json:"textPost" bson:"textPost"`
	ImgPostUrl string             `json:"imgPostUrl" bson:"imgPostUrl"`
}
