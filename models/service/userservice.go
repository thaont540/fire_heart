package service

import (
	"context"
	"fire_heart/models/db"
	"fire_heart/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {}

func (userService *UserService)Create(user entity.User) (inserted *mongo.InsertOneResult, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	insertResult, err := collection.InsertOne(context.TODO(), user)

	return insertResult, err
}

func (userService *UserService)FindByEmail(email string) (found entity.User, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	filter := bson.D{{"email", email}}
	err = collection.FindOne(context.TODO(), filter).Decode(&found)

	return found, err
}

func (userService *UserService) DeleteByEmail(email string) (deleteResult *mongo.DeleteResult, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	filter := bson.D{{"email", email}}
	deleteResult, err = collection.DeleteOne(context.TODO(), filter)

	return deleteResult, err
}
