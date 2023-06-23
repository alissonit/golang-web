package models

import (
	"context"
	"fmt"

	"main/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	Id          primitive.ObjectID
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {

	db := db.ConnectDB()

	collection := db.Database("alura-web").Collection("products")

	filter := bson.D{{}}

	result, err := collection.Find(context.TODO(), filter)

	p := Product{}
	products := []Product{}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil
		}
		panic(err)
	} else {

		for result.Next(context.TODO()) {
			err := result.Decode(&p)

			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				panic(err)
				// If there are no cursor.Decode errors
			}
			products = append(products, p)

		}

		defer func() {
			if err = db.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
		return products
	}
}

func CreateNewProduct(name string, description string, price float64, quantity int) {
	db := db.ConnectDB()

	collection := db.Database("alura-web").Collection("products")

	insertValue := bson.D{
		{"Name", name},
		{"Description", description},
		{"Price", price},
		{"Quantity", quantity},
	}

	_, err := collection.InsertOne(context.TODO(), insertValue)

	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

}
