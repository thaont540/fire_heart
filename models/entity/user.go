package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

var UserCollection = "users"

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserName string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	//VerifiedAt *time.Time
	//Created time.Time `bson:"_created" json:"_created"`
	//UpdatedAt time.Time `bson:"_modified" json:"_modified"`
}
