package main

import (
	"context"
	"fmt"
	"inventory/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	loadEnv()

	fmt.Println("Connecting...")
	mongoUri := os.Getenv("DB_URI")
	fmt.Println(mongoUri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Mongo. Verifying connection...")
	if pingErr := client.Ping(context.TODO(), readpref.Primary()); pingErr != nil {
		panic(pingErr)
	}

	router := gin.Default()

	router.GET("/pstorages", routes.GetParentStorages)
	router.Run("localhost:3000")
}
