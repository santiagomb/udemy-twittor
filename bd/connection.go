package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var MongoClient = Connect()

func Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://santiagomb:fr5l7c5K@cluster0-cidzf.mongodb.net/test?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	log.Println("Mongo DB connected succesfully!")
	return client
}

func Ping() bool {
	log.Println("Ping .......")
	err := MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	log.Println("Pong!")
	return true
}
