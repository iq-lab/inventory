package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collectionName string) (*mongo.Collection, error) {
	mongoUri := os.Getenv("DB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	dbName := os.Getenv("DB_NAME")
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return collection, nil
}
