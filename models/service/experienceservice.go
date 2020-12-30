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

type ExperienceService struct {}

func (experienceService *ExperienceService)Create(experience entity.Experience) (inserted *mongo.InsertOneResult, err error) {
	collection := db.Database.Collection(entity.ExperienceCollection)
	insertResult, err := collection.InsertOne(context.TODO(), experience)

	return insertResult, err
}

func (experienceService *ExperienceService)FindByUserId(userId string) (results []*entity.Experience, err error) {
	collection := db.Database.Collection(entity.ExperienceCollection)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"startdate", 1}})

	filter := bson.D{{"userid", userId}}
	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem entity.Experience
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
