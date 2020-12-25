package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Database *mongo.Database

func Connection() {
	fmt.Println("Starting to connect MongoDB!")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)

		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)

		return
	}
	//defer client.Disconnect(ctx)

	Database = client.Database("fire_heart")

	fmt.Println("Connected to MongoDB!")
}
