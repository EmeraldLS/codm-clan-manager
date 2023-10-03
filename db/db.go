package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var AttendanceCollection *mongo.Collection
var PlayersCollection *mongo.Collection

func dbInstance(uri string) (*mongo.Client, error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()
	if err != nil {
		return &mongo.Client{}, err
	}
	return client, nil
}

func openCollection() error {
	uri := os.Getenv("API_URI")
	db := os.Getenv("API_DB_NAME")
	col1 := os.Getenv("API_COLLECTION_ONE")
	col2 := os.Getenv("API_COLLECTION_TWO")
	client, err := dbInstance(uri)
	if err != nil {
		return err
	}
	AttendanceCollection = client.Database(db).Collection(col1)
	PlayersCollection = client.Database(db).Collection(col2)
	return nil
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	if err := openCollection(); err != nil {
		log.Fatalf("An Error Occured. %v\n", err)
	}
}
