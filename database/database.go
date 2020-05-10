package database

import (
	"context"
	"github.com/santiagomb/udemy-twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

func GetUser(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataBase := MongoClient.Database("twittor")
	collection := dataBase.Collection("users")

	filter := bson.M{"email": email}

	var user *models.User
	err := collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func SaveUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataBase := MongoClient.Database("twittor")
	collection := dataBase.Collection("users")

	var err error

	err = encryptPassword(user)
	if err != nil {
		return err
	}

	var result *mongo.InsertOneResult
	result, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)
	log.Println(objectID)
	return nil
}

func encryptPassword(user *models.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	user.Password = string(encryptedPassword)
	return nil
}
