package db

import (
	"context"
	"inventory/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertParentStorage(storage models.ParentStorage) (*mongo.InsertOneResult, error) {
	collection, err := GetCollection("parentStorage")
	if err != nil {
		return nil, err
	}

	result, err := collection.InsertOne(context.TODO(), storage)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetParentStorages() (any, error) {
	collection, err := GetCollection("parentStorage")
	if err != nil {
		return nil, err
	}

	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	// TODO: Figure out how to serialize everything into structs
	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	data := []primitive.M{}
	for _, result := range results {
		cursor.Decode(&result)
		data = append(data, result)
	}

	return data, nil
}
