package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1

	db_password := os.Getenv("DB_PASSWORD")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	stringConnection := fmt.Sprintf(
		"mongodb+srv://alissonskt:%s@api-cluster.luhg2mi.mongodb.net/?retryWrites=true&w=majority", db_password)
	opts := options.Client().ApplyURI(stringConnection).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
