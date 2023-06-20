package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Id          primitive.ObjectID
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func connectDB() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://alissonskt:UMFybDMljV2aLxqL@api-cluster.luhg2mi.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func getProducts() *mongo.Cursor {

	db := connectDB()

	collection := db.Database("alura-web").Collection("products")

	filter := bson.D{{}}

	result, err := collection.Find(context.TODO(), filter)

	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil
		}
		panic(err)
	} else {
		return result
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	p := Product{}
	products := []Product{}

	result := getProducts()

	for result.Next(context.TODO()) {
		err := result.Decode(&p)

		if err != nil {
			fmt.Println("cursor.Next() error:", err)
			panic(err)
			// If there are no cursor.Decode errors
		}
		products = append(products, p)

	}

	temp.ExecuteTemplate(w, "Index", products)
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
