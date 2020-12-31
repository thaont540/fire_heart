package service

import (
	"context"
	"fire_heart/models/db"
	"fire_heart/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type EducationService struct {}

func (educationService *EducationService)Create(education entity.Education) (inserted *mongo.InsertOneResult, err error) {
	collection := db.Database.Collection(entity.EducationCollection)
	insertResult, err := collection.InsertOne(context.TODO(), education)

	return insertResult, err
}

func (educationService *EducationService)FindByUserId(userId string) (results []*entity.Education, err error) {
	collection := db.Database.Collection(entity.EducationCollection)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"startdate", 1}})

	filter := bson.D{{"userid", userId}}
	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem entity.Education
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	return results, err
}

