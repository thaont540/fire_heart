package service

import (
	"context"
	"fire_heart/models/db"
	"fire_heart/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {}

func (userService *UserService)Create(user entity.User) (inserted *mongo.InsertOneResult, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	insertResult, err := collection.InsertOne(context.TODO(), user)

	return insertResult, err
}

func (userService *UserService)FindById(id string) (user entity.User, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return entity.User{}, err
	}

	filter := bson.D{{"_id", objectId}}
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	user.Password = "secret"

	return user, err
}

func (userService *UserService)DeleteById(id string) (deleteCount int64, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, err
	}

	filter := bson.D{{"_id", objectId}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if deleteResult != nil {
		return deleteResult.DeletedCount, err
	}

	return 0, err
}

func (userService *UserService)FindByEmail(email string) (found entity.User, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	filter := bson.D{{"email", email}}
	err = collection.FindOne(context.TODO(), filter).Decode(&found)
	found.Password = "secret"

	return found, err
}

func (userService *UserService)FindByUsername(userName string) (found entity.User, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	filter := bson.D{{"username", userName}}
	err = collection.FindOne(context.TODO(), filter).Decode(&found)

	return found, err
}

func (userService *UserService) DeleteByEmail(email string) (deleteResult *mongo.DeleteResult, err error) {
	collection := db.Database.Collection(entity.UserCollection)
	filter := bson.D{{"email", email}}
	deleteResult, err = collection.DeleteOne(context.TODO(), filter)

	return deleteResult, err
}
