package main

import (
	"main/routes"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          primitive.ObjectID
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
