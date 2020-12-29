package service

import (
	"context"
	"fire_heart/models/db"
	"fire_heart/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileService struct {}

func (profileService *ProfileService)Create(profile entity.Profile) (inserted *mongo.InsertOneResult, err error) {
	collection := db.Database.Collection(entity.ProfileCollection)
	insertResult, err := collection.InsertOne(context.TODO(), profile)

	return insertResult, err
}

func (profileService *ProfileService)FindByUserId(userId string) (profile entity.Profile, err error) {
	collection := db.Database.Collection(entity.ProfileCollection)

	filter := bson.D{{"userid", userId}}
	err = collection.FindOne(context.TODO(), filter).Decode(&profile)

	return profile, err
}

func (profileService *ProfileService)FindByEmail(email string) (found entity.User, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	filter := bson.D{{"email", email}}
	err = collection.FindOne(context.TODO(), filter).Decode(&found)

	return found, err
}

func (profileService *ProfileService) DeleteByEmail(email string) (deleteResult *mongo.DeleteResult, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	filter := bson.D{{"email", email}}
	deleteResult, err = collection.DeleteOne(context.TODO(), filter)

	return deleteResult, err
}
